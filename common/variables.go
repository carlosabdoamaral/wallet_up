package common

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

var (
	MOCK_MODE          = false
	IS_RUNNING_LOCALLY = false
)

var (
	API_PORT = 8080
	Router   = &gin.Engine{}
)

var (
	Database  = &sql.DB{}
	DB_USER   = ""
	DB_PASS   = ""
	DB_HOST   = ""
	DB_NAME   = ""
	DB_PORT   = ""
	DB_SSL    = ""
	DB_URL    = ""
	DB_DRIVER = "postgres"
)

var (
	RABBIT_URL       = ""
	RABBIT_QUEUENAME = ""
	RABBIT_CONN      = &amqp.Connection{}
	RABBIT_CHANNEL   = &amqp.Channel{}
	RABBIT_QUEUE     = &amqp.Queue{}
	RABBIT_PORT      = ""
)
