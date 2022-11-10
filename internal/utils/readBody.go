package utils

import (
	"io/ioutil"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/gin-gonic/gin"
)

func ReadBody(c *gin.Context) ([]byte, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}

	err = CheckErr(err, false, "Error trying to read body")
	if err != nil {
		return nil, err
	}

	common.PrintSuccess("Success reading body")

	return body, nil
}
