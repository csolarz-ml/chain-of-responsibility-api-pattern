package handlers

import (
	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AccountManagerApproval(c *gin.Context) {
	loan := model.LoanRequest{}
	c.ShouldBindBodyWith(&loan, binding.JSON)

	if loan.Amount < 10000 {
		c.AbortWithStatusJSON(200, gin.H{
			"message":   "loan ok!",
			"signed_by": "Account Manager",
		})
	}

	c.Next()
}
