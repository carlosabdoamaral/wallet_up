package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
)

func CreateWallet(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	model := &pb.CreateWalletRequest{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.WalletServiceClient.Create(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func WalletDetails(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	model := &pb.Id{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.WalletServiceClient.Details(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func EditWallet(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	model := &pb.EditWalletRequest{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.WalletServiceClient.Edit(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func DeleteWallet(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	model := &pb.Id{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.WalletServiceClient.Delete(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, res)
}
