package main

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/api"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

func main() {
	common.PrintInfo("Starting...")

	common.PrintInfo("Reading .env")
	ReadEnvVariables()

	common.PrintInfo("Connecting to database")
	db.Init()

	common.PrintInfo("Starting API")
	api.InitAPI()
}

func ReadEnvVariables() {
	envErr := utils.GetEnvVariables()
	if envErr != nil {
		common.PrintError(envErr.Error())
		common.PrintFatal("Failed to load env file. Make sure .env file exists!")
	}
}
