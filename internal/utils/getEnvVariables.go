package utils

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"os"
	"strconv"
)

func GetEnvVariables() {
	_apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))
	common.API_PORT = _apiPort

	common.DB_USER = os.Getenv("DB_USER")
	common.DB_PASS = os.Getenv("DB_PASS")
	common.DB_HOST = os.Getenv("DB_HOST")
	common.DB_NAME = os.Getenv("DB_NAME")
	common.DB_PORT = os.Getenv("DB_PORT")
	common.DB_SSL = os.Getenv("DB_SSL")
	common.DB_DRIVER = os.Getenv("DB_DRIVER")

	common.RABBIT_URL = os.Getenv("RABBIT_URL")
	common.RABBIT_QUEUENAME = os.Getenv("RABBIT_QUEUENAME")
}
