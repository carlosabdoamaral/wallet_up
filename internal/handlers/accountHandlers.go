package handlers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/db"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/producer"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
	"github.com/gin-gonic/gin"
)

func NewAccountHandler(c *gin.Context) {
	common.PrintStartMethod("NewAccountHandler")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.PrintError("Error reading body!")
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	common.PrintSuccess("Success Reading body")

	// TODO: Validar se o body esta vazio ou não de fato
	account := &models.NewAccountRequest{}
	err = json.Unmarshal(body, &account)
	if err != nil {
		common.PrintError("Error unmarshalling body")
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}
	common.PrintSuccess("Success unmarshalling body")

	err = producer.SendMessage(body, "NewAccount")
	if err != nil {
		common.PrintError("Error sending message to rabbitmq")
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	common.PrintSuccess("Success sending message to RabbitMQ")

}

func AccountDetailsHandler(c *gin.Context) {
	common.PrintStartMethod("AccountDetails")

	body, err := ioutil.ReadAll(c.Request.Body)
	utils.CheckErr(err, false, "Error trying to read body")
	common.PrintSuccess("Success reading body")

	accountId := &models.AccountId{}
	err = json.Unmarshal(body, &accountId)
	err = utils.CheckErr(err, false, "Error trying to unmarshal JSON")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	common.PrintSuccess("Success tring to unmarshal body")

	res, err := db.AccountDetails(accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	} else {
		c.IndentedJSON(http.StatusOK, res)
		return
	}
}

func EditAccountHandler(c *gin.Context) {
	common.PrintStartMethod("EditAccount")

	body, err := utils.ReadBody(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("Success reading body")

	newAccount := &models.EditAccountRequest{}
	err = json.Unmarshal(body, &newAccount)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("Success tring to unmarshal body")

	producer.SendMessage(body, "EDITACCOUNT")
}

func SoftDeleteAccountHandler(c *gin.Context) {
	common.PrintStartMethod("SoftDeleteAccountHandler")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("Success reading body")

	accountID := &models.AccountId{}
	err = json.Unmarshal(body, &accountID)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("Success unmarshaling body")

	if accountID.AccountId <= 0 {
		err = errors.New("account ID must be greater than 0")
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("AccountID value is ok!")

	producer.SendMessage(body, "SOFTDELETEACCOUNT")
}

func RestoreAccountHandler(c *gin.Context) {}

func AccountSettingsDetailsHandler(c *gin.Context) {}

func EditAccountSettingsHandler(c *gin.Context) {}
