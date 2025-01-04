package service

import (
	database "KeyVault-Manager/Database"
	utils "KeyVault-Manager/Utils"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	signupUserStmt = "INSERT INTO users (username,password) VALUES ($1,$2)"
	CheckuserExist = "select * from users where username = '%s'"
	DeleteAllRows  = "Delete from users"
	DeleteUserStmt = "Delete from users where username = %s"
)

func SignupUserController(c *gin.Context, usr utils.CreateAccountInput) error {
	return database.DbManager.InsertToDb(c, signupUserStmt, usr.UserName, usr.Password)
}

func CheckUserExistService(c *gin.Context, name string) (bool, error) {
	return database.DbManager.GetUserFromDB(c, fmt.Sprintf(CheckuserExist, name))
}

func GetUserService(c *gin.Context, name string) (usr utils.CreateAccountInput, err error) {
	return database.DbManager.GetUserFromDBService(c, fmt.Sprintf(CheckuserExist, name))
}
