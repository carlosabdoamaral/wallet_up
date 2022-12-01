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

			// ACCOUNT
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

			// CONFIG
			switch m.Type {
			case "CREATEAPPCONFIG":
				CreateAppConfig(&m)
			case "UPDATEAPPCONFIG":
				UpdateAppConfig(&m)
			case "DELETEAPPCONFIG":
				DeleteAppConfig(&m)
			}

			//TODO: WALLET
			switch m.Type {
			case "CREATEWALLET":
				CreateWallet(&m)
			case "EDITWALLET":
				EditWallet(&m)
			case "DELETEWALLET":
				DeleteWallet(&m)
			case "SHAREWALLET":
				ShareWallet(&m)
			case "UNSHAREWALLET":
				UnShareWallet(&m)
			}

			//TODO: OPERATION
			switch m.Type {
			case "DEPOSITWALLET":
				Deposit(&m)
			case "WITHDRAWWALLET":
				Withdraw(&m)
			case "DELETETRANSACTION":
				DeleteTransaction(&m)
			case "EDITTRANSACTION":
				EditTransaction(&m)
			}
		}
	}()

	<-forever
}
