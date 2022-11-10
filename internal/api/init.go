package api

import (
	"fmt"

	"github.com/carlosabdoamaral/wallet_up/common"
	"github.com/carlosabdoamaral/wallet_up/internal/handlers"
	"github.com/carlosabdoamaral/wallet_up/internal/middlewares"
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
	account.POST("/new", handlers.NewAccountHandler)
	account.POST("/details", handlers.AccountDetailsHandler)
	account.PUT("/edit", handlers.EditAccountHandler)
	account.DELETE("/soft-delete", handlers.SoftDeleteAccountHandler)
	account.POST("/restore", handlers.RestoreAccountHandler)

	// SETTINGS
	accountSettings := account.Group("/settings")
	accountSettings.GET("/details", handlers.AccountSettingsDetailsHandler)
	accountSettings.PUT("/new", handlers.EditAccountSettingsHandler)

	// WALLET
	wallet := common.Router.Group("/wallet")
	wallet.POST("/new", handlers.NewWalletHandler)
	wallet.GET("/details", handlers.WalletDetailsHandler)
	wallet.PUT("/edit", handlers.EditWalletHandler)
	wallet.DELETE("/delete", handlers.DeleteWalletHandler)

	// OPERATIONS
	walletOperation := wallet.Group("/operation")
	walletOperation.POST("/deposit", handlers.DepositHandler)
	walletOperation.POST("/withdraw", handlers.WithdrawHandler)

	// DASHBOARD
	dashboard := common.Router.Group("/dashboard")
	dashboard.GET("", handlers.MakeDashboardHandler)
}
