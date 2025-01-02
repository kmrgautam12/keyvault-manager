package service

import (
	account "Authentication-Go/Controller/Account"
	database "Authentication-Go/Database"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	signupUserStmt = "INSERT INTO users (username,password) VALUES (%s,%s)"
	CheckuserExist = "select * from users where username = '%s'"
)

func SignupUserController(c *gin.Context, usr account.CreateAccountInput) {
	database.DbManager.InsertToDb(c, fmt.Sprintf(signupUserStmt, usr.UserName, usr.Password))
}

func GenerateJWTTokenService(c *gin.Context, name string) (bool, error) {
	return database.DbManager.GetUserFromDB(c, fmt.Sprintf(CheckuserExist, name))
}
