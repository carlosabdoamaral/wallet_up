package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/api"
)

func main() {
	cmd.DefaultInit()

	common.PrintInfo("Starting API")
	api.Init()
}
