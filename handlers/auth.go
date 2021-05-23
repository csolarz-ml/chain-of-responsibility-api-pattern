package handlers

import (
	"net/http"

	"github.com/csolarz-ml/chain-of-responsibility-api-pattern/model"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Auth(c *gin.Context) {
	key := c.Request.Header.Get("Secret-Api-Key")

	if key != "" {
		c.Header("Powered-By", "GNB")

		loan := &model.LoanRequest{}
		if err := c.ShouldBindBodyWith(&loan, binding.JSON); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "invalid loan request!",
			})
		}

		c.Next()
	}

	c.AbortWithStatus(http.StatusUnauthorized)
}
