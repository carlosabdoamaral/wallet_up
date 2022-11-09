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
	common.LogInfo("Starting backend")

	common.LogInfo("Getting .env variables")
	utils.GetEnvVariables()

	common.LogInfo("Starting API")
	InitAPI()

}

func InitAPI() {
	r := gin.Default()
	r.Use(middlewares.CORS())
	r.GET("/hello", handlers.HelloHandler)

	port := fmt.Sprintf(":%d", common.API_PORT)
	err := r.Run(port)
	if err != nil {
		common.LogFatal("Error starting API")
		return
	}
}
