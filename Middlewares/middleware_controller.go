package middlewares

import "github.com/gin-gonic/gin"

func GenerateTokenController(c *gin.Context) {
	GenerateJWTToken(c)
}
