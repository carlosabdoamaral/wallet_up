package db

import "github.com/carlosabdoamaral/wallet_up/common"

func Init() {
	_, err := Connect()
	if err != nil {
		common.PrintError("Error connecting to database")
		common.PrintFatal(err.Error())
	}
}
