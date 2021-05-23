package main

import (
	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//a middleware is a first handler for chain routes
	r.Use(handlers.Auth)

	r.GET("/ping", handlers.Ping)
	r.POST("/loan", handlers.AccountExecutiveApproval, handlers.AccountManagerApproval, handlers.BankManagerApproval)

	r.Run()
}
