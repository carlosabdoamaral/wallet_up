package main

import (
	"github.com/carlosabdoamaral/wallet_up/cmd"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/consumer"
)

func main() {
	cmd.DefaultInit()
	consumer.DeclareQueue()
	consumer.Start()
}
