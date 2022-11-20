package db

import "github.com/carlosabdoamaral/wallet_up/common"

func Disconnect() {
	common.Database.Close()
}
