package rabbit

import (
	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/streadway/amqp"
)

func Connect() {
	conn, err := amqp.Dial(common.RABBIT_URL)
	if err != nil {
	}

	ch, err := conn.Channel()
	if err != nil {
	}

	common.RabbitConn = conn
	common.RabbitChannel = ch
}
