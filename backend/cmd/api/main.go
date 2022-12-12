package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/services/api"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
)

func main() {
	common.PrintSuccess("[API] Starting")
	cmd.DefaultInit()
	utils.ConnectToGRPCServer()
	api.Init()
}
