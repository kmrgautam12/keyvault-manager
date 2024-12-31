package service

import (
	database "Authentication-Go/Database"

	"github.com/gin-gonic/gin"
)

func SignupUserController() {}

func GenerateJWTTokenService(c *gin.Context) {
	database.DbManager.InsertToDb(c)
}
