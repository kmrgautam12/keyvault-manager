package main

import (
	database "Authentication-Go/Database"
	routes "Authentication-Go/Routes"

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
