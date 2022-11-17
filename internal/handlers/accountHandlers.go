package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/producer"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
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

	model := &pb.NewAccountRequest{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		common.PrintError(err.Error())
		return
	}

	res, err := common.AccountServiceClient.Create(context.Background(), model)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res.GetStatus())
}

func AccountDetailsHandler(c *gin.Context) {
	common.PrintStartMethod("AccountDetails")

	body, err := ioutil.ReadAll(c.Request.Body)
	utils.CheckErr(err, false, "Error trying to read body")
	common.PrintSuccess("Success reading body")

	accountId := &pb.Id{}
	err = json.Unmarshal(body, &accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}
	common.PrintSuccess("Success tring to unmarshal body")

	res, err := common.AccountServiceClient.Details(c.Request.Context(), accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func EditAccountHandler(c *gin.Context) {
	common.PrintStartMethod("EditAccount")

	body, err := utils.ReadBody(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	account := &pb.EditAccountRequest{}
	err = json.Unmarshal(body, account)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	common.PrintSuccess("Success trying to unmarshal body")

	common.AccountServiceClient.Edit(c, account)

	c.IndentedJSON(http.StatusAccepted, "Success to append into queue!")
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
