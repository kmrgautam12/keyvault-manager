package middlewares

import (
	service "KeyVault-Manager/Database/Service"
	utils "KeyVault-Manager/Utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var secretString = "secret-string"

func GenerateJWTToken(c *gin.Context) {

	var user utils.GenerateJWTInput
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}

	exist, err := service.CheckUserExistService(c, user.UserName)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}
	if !exist {
		utils.SentErrorResponse500(c, fmt.Errorf("user %s doesn't exist", user.UserName))
		return
	}
	t, err := CreateClaimsAndToken(user.UserName)
	if err != nil {
		utils.SentErrorResponse500(c, fmt.Errorf("error generating user token, user : %s", user.UserName))
	}

	utils.SentSuccessResponse200(c, t)
}

func CreateClaimsAndToken(user string) (t utils.GenerateJWTTokenOutput, err error) {
	t.TokenId = uuid.NewString()
	t.ValidUntil = time.Now().Add((JwtClaimExpire)).Unix()

	tClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        uuid.NewString(),
		ExpiresAt: int64(JwtClaimExpire),
		IssuedAt:  time.Now().Unix(),
		Issuer:    user,
		Audience:  "public",
	})

	t.Token, err = tClaim.SignedString([]byte(secretString))
	if err != nil {
		return t, err
	}
	return t, nil

}
