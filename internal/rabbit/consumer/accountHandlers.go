package consumer

import (
	"encoding/json"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewAccountHandler(m *amqp.Delivery) {
	account := &pb.NewAccountRequest{}
	err := protojson.Unmarshal(m.Body, account)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	db.NewAccount(account)
}

func EditAccountHandler(m *amqp.Delivery) {
	account := &pb.EditAccountRequest{}

	err := protojson.Unmarshal(m.Body, account)
	if err != nil {
		common.PrintError("Error while unmarshaling JSON")
		return
	}

	db.EditAccount(account)
}

func SoftDeleteAccountHandler(m *amqp.Delivery) {
	common.PrintInfo("[RABBIT] SoftDeleteAccountHandler")

	accountID := &models.AccountId{}
	err := json.Unmarshal(m.Body, &accountID)
	if err != nil {
		common.PrintError("Error while unmarshaling JSON")
		return
	}

	db.SoftDeleteAccount(accountID)
}
