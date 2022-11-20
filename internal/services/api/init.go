package api

import (
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/services/api/middlewares"
	"github.com/gin-gonic/gin"
)

func Init() {
	common.Router = gin.Default()
	common.Router.Use(middlewares.CORS())

	MakeRoutes()

	common.Router.Run(fmt.Sprintf(":%d", common.API_PORT))
}

func MakeRoutes() {
	// ACCOUNT
	account := common.Router.Group("/account")
	account.POST("/new", NewAccountHandler)
	account.POST("/details", AccountDetailsHandler)
	account.PUT("/edit", EditAccountHandler)
	account.DELETE("/soft-delete", SoftDeleteAccountHandler)
	account.POST("/restore", RestoreAccountHandler)

	// SETTINGS
	accountSettings := account.Group("/settings")
	accountSettings.GET("/details", AccountSettingsDetailsHandler)
	accountSettings.POST("/new", EditAccountSettingsHandler)

	// WALLET
	wallet := common.Router.Group("/wallet")
	wallet.POST("/new", NewWalletHandler)
	wallet.GET("/details", WalletDetailsHandler)
	wallet.PUT("/edit", EditWalletHandler)
	wallet.DELETE("/delete", DeleteWalletHandler)

	// OPERATIONS
	walletOperation := wallet.Group("/operation")
	walletOperation.POST("/deposit", DepositHandler)
	walletOperation.POST("/withdraw", WithdrawHandler)

	// DASHBOARD
	dashboard := common.Router.Group("/dashboard")
	dashboard.GET("", MakeDashboardHandler)
}
