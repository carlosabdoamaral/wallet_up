package consumer

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func NewAccountHandler(m *amqp.Delivery) {
	account := &pb.NewAccountRequest{}
	err := protojson.Unmarshal(m.Body, account)
	if err != nil {
		return
	}

	db.NewAccount(account)
}

func EditAccountHandler(m *amqp.Delivery) {
	account := &pb.EditAccountRequest{}
	err := protojson.Unmarshal(m.Body, account)
	if err != nil {
		return
	}

	db.EditAccount(account)
}

func SoftDeleteAccountHandler(m *amqp.Delivery) {
	id := &pb.Id{}
	err := protojson.Unmarshal(m.Body, id)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	err = db.SoftDeleteAccount(id)
	if err != nil {
		common.PrintError(err.Error())
		return
	}
}

func RestoreAccountHandler(m *amqp.Delivery) {
	id := &pb.Id{}
	err := protojson.Unmarshal(m.Body, id)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	err = db.RestoreAccount(id)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	common.PrintSuccess("Account restored successfully")
}
