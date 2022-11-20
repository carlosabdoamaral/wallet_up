package api

import "github.com/gin-gonic/gin"

func NewWalletHandler(c *gin.Context)     {}
func WalletDetailsHandler(c *gin.Context) {}
func EditWalletHandler(c *gin.Context)    {}
func DeleteWalletHandler(c *gin.Context)  {}

func DepositHandler(c *gin.Context)  {}
func WithdrawHandler(c *gin.Context) {}

func NewWalletCategoryHandler(c *gin.Context)     {}
func WalletCategoryDetailsHandler(c *gin.Context) {}
func EditWalletCategoryHandler(c *gin.Context)    {}
func DeleteWalletCategoryHandler(c *gin.Context)  {}
