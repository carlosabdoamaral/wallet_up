package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/grpc"
)

func main() {
	common.PrintStartMethod("Init grpc server")
	
	cmd.DefaultInit()
	grpc.InitServer()
}
