package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/models"
	pb "github.com/carlosabdoamaral/wallet_up/protodefs/gen/proto"
	"github.com/gin-gonic/gin"
)

func CreateConfig(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	model := &pb.NewAppConfigRequest{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res, err := common.AppConfigServiceClient.Create(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func UpdateConfig(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	model := &pb.AppConfigDetails{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := common.AppConfigServiceClient.Update(c.Request.Context(), model)
	if err != nil {
		common.PrintError(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}

func ConfigDetails(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	model := &pb.Id{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res, err := common.AppConfigServiceClient.Details(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, &models.AppConfigDetails{
		IdConfig:          res.GetIdConfig(),
		AppTheme:          res.GetTheme(),
		BiometryActivated: res.GetBiometryActivated(),
		AlertOnEmail:      res.GetAlertOnEmail(),
		AlertOnMobile:     res.GetAlertOnMobile(),
		AppLanguageId:     res.GetIdLanguage(),
		AppLanguage:       res.GetAppLanguage(),
		AppLanguageKey:    res.GetAppLanguageKey(),
	})
}

func DeleteConfig(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	model := &pb.Id{}
	err = json.Unmarshal(body, &model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	res, err := common.AppConfigServiceClient.Delete(c.Request.Context(), model)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, res)
}
