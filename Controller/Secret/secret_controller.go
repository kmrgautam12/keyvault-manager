package secret

import "github.com/gin-gonic/gin"

func CreateSecretController(c *gin.Context) {
	CreateSecret(c)
}
