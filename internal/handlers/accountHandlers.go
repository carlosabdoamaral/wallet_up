package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	"github.com/carlosabdoamaral/wallet_up/internal/rabbit/producer"
	"github.com/gin-gonic/gin"
)

func NewAccountHandler(c *gin.Context) {
	fmt.Println("")
	common.PrintInfo("Received NewAccount request!")

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

func AccountDetailsHandler(c *gin.Context) {}

func EditAccountHandler(c *gin.Context) {}

func DeleteAccountHandler(c *gin.Context) {}

func RestoreAccountHandler(c *gin.Context) {}

func AccountSettingsDetailsHandler(c *gin.Context) {}

func EditAccountSettingsHandler(c *gin.Context) {}
