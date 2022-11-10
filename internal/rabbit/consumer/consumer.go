package consumer

import (
	"fmt"
	"log"

	"github.com/carlosabdoamaral/wallet_up/common"
)

func Start() {
	msgs, err := common.RabbitChannel.Consume(
		common.RabbitQueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Println("Error")
	}

	forever := make(chan bool)
	go func() {
		for m := range msgs {
			fmt.Println("")
			common.PrintSuccess("Received new message!")

			switch m.Type {
			case "NEWACCOUNT":
				NewAccountHandler(&m)
			case "EDITACCOUNT":
				EditAccountHandler(&m)
			}
		}
	}()

	common.PrintInfo("Waiting for messages")
	<-forever
}
