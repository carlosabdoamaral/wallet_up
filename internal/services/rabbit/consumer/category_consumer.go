package consumer

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func CreateCategory(m *amqp.Delivery) {
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

func EditCategory(m *amqp.Delivery) {
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

func DeleteCategory(m *amqp.Delivery) {
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
