package secret

import (
	utils "KeyVault-Manager/Utils"
	"crypto/sha256"
	"encoding/hex"

	"github.com/gin-gonic/gin"
)

func CreateSecret(c *gin.Context) {
	var secretInput utils.CreateSecretInput
	err := c.ShouldBindBodyWithJSON(&secretInput)
	if err != nil {
		panic(err)
	}

	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(secretInput.Value))
	hash := sha256Hash.Sum(nil)
	h := hex.EncodeToString(hash)
	utils.SentSuccessResponse200(c, h)

}
