package handlers

import (
	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AccountExecutiveApproval(c *gin.Context) {
	loan := model.LoanRequest{}
	c.ShouldBindBodyWith(&loan, binding.JSON)

	if loan.Amount < 100 {
		c.AbortWithStatusJSON(200, gin.H{
			"message":   "loan ok!",
			"signed_by": "Account Executive",
		})
	}

	c.Next()
}
