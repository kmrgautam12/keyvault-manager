package account

import (
	service "KeyVault-Manager/Database/Service"
	middlewares "KeyVault-Manager/Middlewares"
	utils "KeyVault-Manager/Utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUpUser(c *gin.Context) {
	var usr utils.CreateAccountInput
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
	exist, err := service.CheckUserExistService(c, usr.UserName)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
	if exist {
		utils.SentErrorResponse400(c, fmt.Errorf("user already exist with username : %s", usr.UserName))
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	usr.Password = string(hash)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}
	err = service.SignupUserController(c, usr)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}
	utils.SentSuccessResponse200(c, fmt.Sprintf("user created with username %s", usr.UserName))
}

func UserLogin(c *gin.Context) {

	var usr utils.CreateAccountInput
	err := c.ShouldBindJSON(&usr)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
	exist, err := service.CheckUserExistService(c, usr.UserName)
	if err != nil {
		utils.SentErrorResponse400(c, err)
		return
	}
	if !exist {
		utils.SentErrorResponse400(c, fmt.Errorf("no user with username : %s", usr.UserName))
		return
	}

	//retrive password from
	dbResp, err := service.GetUserService(c, usr.UserName)
	if err != nil {
		utils.SentErrorResponse500(c, err)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(dbResp.Password), []byte(usr.Password))
	if err != nil {
		utils.SentErrorResponse500(c, fmt.Errorf("your password is not valid"))
		return
	}
	t, err := middlewares.CreateClaimsAndToken(usr.UserName)
	if err != nil {
		utils.SentErrorResponse500(c, fmt.Errorf("login failed !! Please check username and password"))
		return
	}
	utils.SentSuccessResponse200(c, utils.LoginOutputStruct{
		Username: usr.UserName,
		UserIp:   usr.Ip,
		Token:    t,
	})
}
