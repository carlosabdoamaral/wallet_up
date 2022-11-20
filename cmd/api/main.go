package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/api"
	"github.com/carlosabdoamaral/wallet_up/internal/handlers"
)

func main() {
	common.PrintSuccess("[API] Starting")
	cmd.DefaultInit()
	handlers.ConnectToGRPCServer()
	api.Init()
}
