package routes

import (
	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/handlers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", handlers.Ping)
	r.POST("/loan", handlers.Auth, handlers.AccountExecutiveApproval, handlers.AccountManagerApproval, handlers.BankManagerApproval)

	return r
}
