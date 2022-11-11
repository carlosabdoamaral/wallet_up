package common

import (
	"database/sql"

	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
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
	RABBIT_URL  = ""
	RABBIT_PORT = ""

	RabbitConn      = &amqp.Connection{}
	RabbitChannel   = &amqp.Channel{}
	RabbitQueue     = &amqp.Queue{}
	RabbitQueueName = ""
)

var (
	GrpcServer             = &grpc.Server{}
	GrpcConn               = &grpc.ClientConn{}
	AccountServiceClient   = pb.NewAccountServiceClient(GrpcConn)
	AppConfigServiceClient = pb.NewAppConfigServiceClient(GrpcConn)
	WalletServiceClient    = pb.NewWalletServiceClient(GrpcConn)
)
