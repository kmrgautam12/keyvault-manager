package account

import (
	utils "Authentication-Go/Utils"

	"github.com/gin-gonic/gin"
)

func CreateUserAccount(c *gin.Context) {
	var user CreateAccountInput
	err := c.ShouldBindJSON(&user)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
}
