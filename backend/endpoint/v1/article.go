package v1

import (
	"encoding/json"
	"fmt"
	"time"

	modelV1 "github.com/Ulfa24433/kumparan-test/backend/db/v1/model"
	utilV1 "github.com/Ulfa24433/kumparan-test/backend/util/v1"

	//	"github.com/blevesearch/bleve"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type AddArticleRequest struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func AddArticle(c *gin.Context) {
	payload := &AddArticleRequest{}
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "invalid request", err)
		return
	}
	// init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	//add article
	article := &modelV1.Article{
		Author:  payload.Author,
		Title:   payload.Title,
		Body:    payload.Body,
		Created: time.Now(),
	}
	err = tx.Create(&article).Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}
	tx.Commit()

	utilV1.CallSuccessOK(c, "success", nil)
}

func GetArticle(c *gin.Context) {
	//var ctx = context.Background()
	query := c.Query("query")
	author := c.Query("author")
	articles := []*modelV1.Article{}

	// init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}
	client := utilV1.RedisClient()
	log.Info(client)

	//get data from cache
	listCached, err := client.Get(fmt.Sprintf("%s_%s", query, author)).Result()
	if err != nil && err.Error() != "redis: nil" {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}
	if listCached != "" {
		err = json.Unmarshal([]byte(listCached), &articles)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("response from cache")
		utilV1.CallSuccessOK(c, "success", articles)
		return
	}

	//get data from database if not found in cache
	baseModel := db.
		Model(articles)
	if query != "" {
		qry := fmt.Sprintf("%%%s%%", query)
		baseModel = baseModel.Where("title like ? OR body like ?", qry, qry)
	}
	if author != "" {
		baseModel = baseModel.Where("author = ?", author)
	}
	err = baseModel.
		Order("created DESC").
		Find(&articles).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}

	//convert articles array to byte
	payloadBytes, err := json.Marshal(articles)
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}
	//set data to cache
	err = client.Set(fmt.Sprintf("%s_%s", query, author), payloadBytes, 0).Err()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occured", err)
		return
	}

	log.Info("response from db")
	utilV1.CallSuccessOK(c, "success", articles)
}
