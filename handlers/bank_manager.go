package handlers

import (
	"net/http"

	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func BankManagerApproval(c *gin.Context) {
	form := model.LoanForm{}
	c.ShouldBindBodyWith(&form, binding.JSON)

	c.AbortWithStatusJSON(http.StatusOK, &model.Loan{
		Form:     form,
		Message:  "loan ok!",
		SignedBy: "Bank Manager",
	})
}
