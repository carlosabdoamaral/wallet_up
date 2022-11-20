package consumer

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/encoding/protojson"
)

func CreateAppConfig(m *amqp.Delivery) {
	appConfig := &pb.NewAppConfigRequest{}

	err := protojson.Unmarshal(m.Body, appConfig)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.CreateAppConfig(appConfig)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func UpdateAppConfig(m *amqp.Delivery) {
	appConfig := &pb.AppConfigDetails{}

	err := protojson.Unmarshal(m.Body, appConfig)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.UpdateAppConfig(appConfig)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}

func DeleteAppConfig(m *amqp.Delivery) {
	id := &pb.Id{}

	err := protojson.Unmarshal(m.Body, id)
	if err != nil {
		common.PrintFatal(err.Error())
	}

	err = db.DeleteAppConfig(id)
	if err != nil {
		common.PrintFatal(err.Error())
	}
}
