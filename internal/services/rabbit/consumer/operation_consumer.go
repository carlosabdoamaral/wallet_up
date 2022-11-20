package consumer

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func Deposit(m *amqp.Delivery) {
	model := &pb.TransactionRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.Deposit(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func Withdraw(m *amqp.Delivery) {
	model := &pb.TransactionRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.Withdraw(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func DeleteTransaction(m *amqp.Delivery) {
	model := &pb.Id{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.DeleteTransaction(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func EditTransaction(m *amqp.Delivery) {
	model := &pb.EditTransactionRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.EditTransaction(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}
