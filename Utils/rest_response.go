package utils

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	StatusCode int         `json:"status"`
	Response   interface{} `json:"response"`
}

func SentSuccessResponse200(c *gin.Context, r interface{}) {
	c.JSON(200, SuccessResponse{StatusCode: 200, Response: r})
}
func SentErrorResponse500(c *gin.Context, r interface{}) {
	c.JSON(200, SuccessResponse{StatusCode: 500, Response: r})
}
func SentErrorResponse400(c *gin.Context, r interface{}) {
	c.JSON(200, SuccessResponse{StatusCode: 400, Response: r})
}
