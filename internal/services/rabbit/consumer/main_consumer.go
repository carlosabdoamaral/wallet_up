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
			common.PrintSuccess(fmt.Sprintf("New request -> %s", m.Type))

			switch m.Type {
			case "NEWACCOUNT":
				NewAccountHandler(&m)
			case "EDITACCOUNT":
				EditAccountHandler(&m)
			case "SOFTDELETEACCOUNT":
				SoftDeleteAccountHandler(&m)
			case "DELETEACCOUNT":
				DeleteAccount(&m)
			case "RESTOREACCOUNT":
				RestoreAccountHandler(&m)
			}
		}
	}()

	<-forever
}
