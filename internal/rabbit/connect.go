package rabbit

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/streadway/amqp"
)

func Connect() {
	conn, err := amqp.Dial(common.RABBIT_URL)
	if err != nil {
		common.PrintError("Error connecting to RabbitMQ")
		common.PrintFatal(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		common.PrintError("Error connecting to RabbitMQ Channel")
		common.PrintFatal(err.Error())
	}

	common.RabbitConn = conn
	common.RabbitChannel = ch
}
