package consumer

import (
	"encoding/json"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/streadway/amqp"
)

func NewAccountHandler(m *amqp.Delivery) {
	account := &models.NewAccountRequest{}

	err := json.Unmarshal(m.Body, &account)
	if err != nil {
		common.PrintError("Error while unmarshaling JSON")
		return
	}

	db.NewAccount(account)
}
