package api

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/utils"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
)

func NewAccountHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
	}

	model := &pb.NewAccountRequest{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		return
	}

	res, err := common.AccountServiceClient.Create(context.Background(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res.GetStatus())
}

func AccountDetailsHandler(c *gin.Context) {

	body, err := ioutil.ReadAll(c.Request.Body)
	utils.CheckErr(err, false, "Error trying to read body")

	accountId := &pb.Id{}
	err = json.Unmarshal(body, &accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
	}

	res, err := common.AccountServiceClient.Details(c.Request.Context(), accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func EditAccountHandler(c *gin.Context) {

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

	common.AccountServiceClient.Edit(c, account)

	c.IndentedJSON(http.StatusAccepted, "Success to append into queue!")
}

func SoftDeleteAccountHandler(c *gin.Context) {

	body, err := utils.ReadBody(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	id := &pb.Id{}
	if json.Unmarshal(body, id) != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	statusMessage, err := common.AccountServiceClient.SoftDelete(context.Background(), id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, statusMessage)
}

func RestoreAccountHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, err.Error())
		return
	}

	accountId := &pb.Id{}
	err = json.Unmarshal(body, &accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.AccountServiceClient.Restore(c.Request.Context(), accountId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}
