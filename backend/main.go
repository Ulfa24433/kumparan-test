package main

import (
	"fmt"
	"os"

	endpointV1 "github.com/Ulfa24433/kumparan-test/backend/endpoint/v1"
	"github.com/Ulfa24433/kumparan-test/backend/util/v1/envvar"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {

	r := gin.Default()
	v1 := r.Group("/v1")

	//Endpoint to create an article
	article := v1.Group("/article")
	article.POST("/add", endpointV1.AddArticle)
	article.GET("/list", endpointV1.GetArticle)

	port := os.Getenv(envvar.Port)

	err := r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		logrus.Fatal(err)
	}
}

//docker run -p 5432:5432 --name postgres1 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -e POSTGRES_DB=kumparan -v $(pwd)/postgres/init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
