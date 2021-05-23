package handlers

import (
	"net/http"

	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func AccountExecutiveApproval(c *gin.Context) {
	form := model.LoanForm{}
	c.ShouldBindBodyWith(&form, binding.JSON)

	if form.Amount < 100 {
		c.AbortWithStatusJSON(http.StatusOK, &model.Loan{
			Form:     form,
			Message:  "loan ok!",
			SignedBy: "Account Executive",
		})
	}

	c.Next()
}
