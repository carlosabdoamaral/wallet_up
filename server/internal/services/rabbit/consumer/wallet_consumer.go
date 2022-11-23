package consumer

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func CreateWallet(m *amqp.Delivery) {
	model := &pb.CreateWalletRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.CreateWallet(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func EditWallet(m *amqp.Delivery) {
	model := &pb.EditWalletRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.EditWallet(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func DeleteWallet(m *amqp.Delivery) {
	model := &pb.Id{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.DeleteWallet(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func ShareWallet(m *amqp.Delivery) {
	model := &pb.ShareWalletRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.ShareWallet(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func UnShareWallet(m *amqp.Delivery) {
	model := &pb.UnShareWalletRequest{}

	err := protojson.Unmarshal(m.Body, model)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.UnShareWallet(model)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}
