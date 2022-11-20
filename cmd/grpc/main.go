package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/services/grpc"
)

func main() {
	common.PrintSuccess("[GRPC] Starting")
	cmd.DefaultInit()
	grpc.InitServer()
}
