package routes

import (
	account "KeyVault-Manager/Controller/Account"
	middlewares "KeyVault-Manager/Middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	r.POST("/token", middlewares.GenerateTokenController)
	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/signup", account.SignUpUserController)
		accountGroup.POST("/login", account.UserLoginController)
	}
}
