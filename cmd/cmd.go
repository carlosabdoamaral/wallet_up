package cmd

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

func DefaultInit() {
	common.PrintStartMethod("Starting...")

	common.PrintInfo("[ENV] Reading...")
	ReadEnvVariables()
	common.PrintSuccess("[ENV] Success...")

	common.PrintInfo("[AMQP] Connecting...")
	rabbit.Connect()
	common.PrintSuccess("[AMQP] Success...")

	common.PrintInfo("[DB] Connecting...")
	db.Init()
	common.PrintSuccess("[DB] Success...")
}

func ReadEnvVariables() {
	envErr := utils.GetEnvVariables()
	if envErr != nil {
		common.PrintError(envErr.Error())
		common.PrintFatal("Failed to load env file. Make sure .env file exists!")
	}
}
