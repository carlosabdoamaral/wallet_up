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
	r := gin.Default()
	r.Use(middlewares.CORS())
	r.GET("/hello", handlers.HelloHandler)

	port := fmt.Sprintf(":%d", common.API_PORT)
	err := r.Run(port)
	if err != nil {
		common.PrintError("Error starting API!")
		return
	}
}

func ReadEnvVariables() {
	envErr := utils.GetEnvVariables()
	if envErr != nil {
		common.PrintError(envErr.Error())
		common.PrintFatal("Failed to load env file. Make sure .env file exists!")
	}
}
