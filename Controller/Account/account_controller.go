package account

import "github.com/gin-gonic/gin"

func SignUpUserController(c *gin.Context) {
	SignUpUser(c)
}
func UserLoginController(c *gin.Context) {
	UserLogin(c)
}
