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
	account.POST("/create", NewAccount)
	account.POST("/details", AccountDetails)
	account.PUT("/edit", EditAccount)
	account.DELETE("/soft-delete", SoftDeleteAccount)
	account.POST("/restore", RestoreAccount)

	// SETTINGS
	accountSettings := account.Group("/settings")
	accountSettings.POST("/create", CreateConfig)
	accountSettings.PUT("/edit", UpdateConfig)
	accountSettings.POST("/details", ConfigDetails)
	accountSettings.DELETE("/delete", DeleteConfig)

	// WALLET
	wallet := common.Router.Group("/wallet")
	wallet.POST("/create", CreateWallet)
	wallet.GET("/details", WalletDetails)
	wallet.PUT("/edit", EditWallet)
	wallet.DELETE("/delete", DeleteWallet)

	// OPERATIONS
	walletOperation := wallet.Group("/operation")
	walletOperation.POST("/deposit", Deposit)
	walletOperation.POST("/withdraw", Withdraw)
	walletOperation.POST("/transaction/delete", DeleteTransaction)
	walletOperation.PUT("/transaction/edit", EditTransaction)

	// CATEGORY
	category := common.Router.Group("/category")
	category.POST("/create", CreateCategory)
	category.PUT("/edit", EditCategory)
	category.GET("/list", CategoryList)
	category.DELETE("/delete", DeleteCategory)

	// DASHBOARD
	dashboard := common.Router.Group("/dashboard")
	dashboard.GET("", DashboardDetails)
}
