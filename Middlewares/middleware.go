package middlewares

import (
	service "Authentication-Go/Database/Service"
	utils "Authentication-Go/Utils"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Ip       string `json:"ip"`
}

type GenerateJWTTokenOutput struct {
	Token      string `json:"token"`
	TokenId    string `json:"token_id"`
	ValidUntil int64  `json:"valid_until"`
}

var secretString = "secret-string"

func Middleware(r *gin.Engine) {

}
func GenerateJWTToken(c *gin.Context) {
	var user User
	var t GenerateJWTTokenOutput
	err := c.ShouldBindBodyWithJSON(&user)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
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

	service.GenerateJWTTokenService(c)
	utils.SentSuccessResponse200(c, t)
}

func SignUpUser(c *gin.Context) {
	var usr User
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}

	// Check whether user is present in the db/not

	service.

		// hashing the password of the user

		hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)

	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}

}
