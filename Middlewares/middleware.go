package middlewares

import (
	account "KeyVault-Manager/Controller/Account"
	service "KeyVault-Manager/Database/Service"
	utils "KeyVault-Manager/Utils"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretString = "secret-string"

func GenerateJWTToken(c *gin.Context) {

	var user account.GenerateJWTInput
	var t account.GenerateJWTTokenOutput
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}

	exist, err := service.GenerateJWTTokenService(c, user.UserName)
	if err != nil {
		utils.SentErrorResponse500(c, err)
	}
	if !exist {
		utils.SentErrorResponse500(c, fmt.Errorf("user %s doesn't exist", user.UserName))
	}

	t.TokenId = uuid.NewString()
	t.ValidUntil = time.Now().Add((JwtClaimExpire)).Unix()

	tClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        uuid.NewString(),
		ExpiresAt: int64(JwtClaimExpire),
		IssuedAt:  time.Now().Unix(),
		Issuer:    user.UserName,
		Audience:  "public",
	})

	t.Token, err = tClaim.SignedString([]byte(secretString))
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}

	utils.SentSuccessResponse200(c, t)
}

func SignUpUser(c *gin.Context) {
	var usr account.CreateAccountInput
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
	// Check whether user is present in the db/not

	service.CheckuserExist()
	// hashing the password of the user

	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}

}
