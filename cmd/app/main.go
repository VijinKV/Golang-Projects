package main

import (
	_ "bank/internal"
	"bank/internal/constants"
	"bank/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/account/:"+constants.URL_PARAMS_ACC_NO, handlers.AccountDetailHandler)
	router.POST("/account/:"+constants.URL_PARAMS_ACC_NO+"/transfer", handlers.TransferMoneyHandler)

	router.Run(":8080")
}
