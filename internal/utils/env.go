package utils

import (
	"os"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/joho/godotenv"
)

func GetEnvVariables() error {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		return errEnv
	}

	common.DB_USER = os.Getenv("DB_USER")
	common.DB_PASS = os.Getenv("DB_PASS")
	common.DB_HOST = os.Getenv("DB_HOST")
	common.DB_NAME = os.Getenv("DB_NAME")
	common.DB_PORT = os.Getenv("DB_PORT")
	common.DB_SSL = os.Getenv("DB_SSL")
	common.DB_DRIVER = os.Getenv("DB_DRIVER")

	common.RABBIT_URL = os.Getenv("RABBIT_URL")
	common.RabbitQueueName = os.Getenv("RABBIT_QUEUENAME")

	return nil
}
