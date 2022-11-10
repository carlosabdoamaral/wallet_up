package utils

import (
	"github.com/carlosabdoamaral/wallet_up/common"
)

func CheckErr(err error, isFatal bool, msg string) error {
	if err != nil {
		common.PrintError(msg)

		if isFatal {
			common.PrintFatal(err.Error())
		}

		return err
	}

	return nil
}
