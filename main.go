package main

import (
	database "KeyVault-Manager/Database"
	routes "KeyVault-Manager/Routes"

	"github.com/gin-gonic/gin"
)

func init() {
	database.InitApplicationLayer()
}

func main() {
	r := gin.Default()
	routes.RegisterRouters(r)
	r.Run(":8080")
}
