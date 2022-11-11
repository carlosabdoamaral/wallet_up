package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/api"
	"github.com/carlosabdoamaral/wallet_up/internal/handlers"
)

func main() {
	cmd.DefaultInit()

	handlers.ConnectToGRPCServer()

	common.PrintInfo("Starting API")
	api.Init()
}
