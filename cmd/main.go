package main

import (
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/handlers"
	"github.com/carlosabdoamaral/wallet_up/internal/middlewares"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	common.PrintInfo("Starting...")

	common.PrintInfo("Reading .env")
	ReadEnvVariables()

	common.PrintInfo("Starting API")
	InitAPI()
}

func InitAPI() {
	router := gin.Default()
	router.Use(middlewares.CORS())

	router.GET("/hello", handlers.HelloHandler)

	router.Run(fmt.Sprintf(":%d", common.API_PORT))
}

func ReadEnvVariables() {
	envErr := utils.GetEnvVariables()
	if envErr != nil {
		common.PrintError(envErr.Error())
		common.PrintFatal("Failed to load env file. Make sure .env file exists!")
	}
}
