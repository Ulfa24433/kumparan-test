package v1

import (
	"fmt"
	"os"

	"github.com/Ulfa24433/kumparan-test/backend/util/v1/envvar"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgreClient() (db *gorm.DB, err error) {
	rdsHost := os.Getenv(envvar.RDSPostgresHost)
	rdsUsername := os.Getenv(envvar.RDSPostgresUsername)
	rdsPassword := os.Getenv(envvar.RDSPostgresPassword)
	rdsDatabase := os.Getenv(envvar.RDSPostgresDatabase)
	rdsSslCert := os.Getenv(envvar.RDSPostgresSslCert)
	rdsPort := os.Getenv(envvar.RDSPostgresPort)
	timezone := os.Getenv(envvar.Timezone)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		rdsHost, rdsUsername, rdsPassword, rdsDatabase, rdsPort, rdsSslCert, timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
