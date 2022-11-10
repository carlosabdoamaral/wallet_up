package cmd

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

func DefaultInit() {
	common.PrintInfo("Starting...")

	common.PrintInfo("Reading .env")
	ReadEnvVariables()

	common.PrintInfo("Connecting to RabbitMQ")
	rabbit.Connect()

	common.PrintInfo("Connecting to database")
	db.Init()
}

func ReadEnvVariables() {
	envErr := utils.GetEnvVariables()
	if envErr != nil {
		common.PrintError(envErr.Error())
		common.PrintFatal("Failed to load env file. Make sure .env file exists!")
	}
}
