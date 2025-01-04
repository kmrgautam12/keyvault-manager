package utils

import (
	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	StatusCode int         `json:"status"`
	Response   interface{} `json:"response"`
}
type ErrorResponse struct {
	StatusCode    int `json:"status"`
	ErrorResponse BuildErrorResponse
}
type BuildErrorResponse struct {
	ErrorCode    string `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	Status       int    `json:"status"`
}

func SentSuccessResponse200(c *gin.Context, r interface{}) {
	c.JSON(200, SuccessResponse{StatusCode: 200, Response: r})
}
func SentErrorResponse500(c *gin.Context, r error) {
	c.JSON(200, ErrorResponse{StatusCode: 500, ErrorResponse: BuildErrorResponse{
		ErrorCode:    "500",
		ErrorMessage: r.Error(),
		Status:       500,
	}})
}
func SentErrorResponse400(c *gin.Context, r error) {
	c.JSON(200, ErrorResponse{StatusCode: 400, ErrorResponse: BuildErrorResponse{
		ErrorCode:    "400",
		ErrorMessage: r.Error(),
		Status:       400,
	}})
}
