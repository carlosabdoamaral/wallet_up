package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/consumer"
)

func main() {
	cmd.DefaultInit()

	common.PrintInfo("Declaring queue")
	consumer.DeclareQueue()

	common.PrintInfo("Starting consumer")
	consumer.Start()
}
