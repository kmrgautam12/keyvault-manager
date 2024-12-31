package routes

import (
	middlewares "Authentication-Go/Middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRouters(r *gin.Engine) {
	r.POST("/token", middlewares.GenerateTokenController)
	accountGroup := r.Group("/account")
	{
		accountGroup.POST("/signup")
	}
}
