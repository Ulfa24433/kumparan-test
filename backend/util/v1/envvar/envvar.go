package envvar

const (
	// general
	Port = "PORT"

	// postgres
	RDSPostgresHost     = "RDS_POSTGRES_HOST"
	RDSPostgresPort     = "RDS_POSTGRES_PORT"
	RDSPostgresDatabase = "RDS_POSTGRES_DATABASE"
	RDSPostgresUsername = "RDS_POSTGRES_USERNAME"
	RDSPostgresPassword = "RDS_POSTGRES_PASSWORD"
	RDSPostgresSslCert  = "RDS_POSTGRES_SSL_CERT"
	Timezone            = "TIMEZONE"

	// redis
	RedisHost  = "REDIS_HOST"
	RedisPort  = "REDIS_PORT"
	RedisIndex = "REDIS_INDEX"
)
