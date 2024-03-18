package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func FormatResponse(c *gin.Context, code int, status string, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}

func ErrorResponse(c *gin.Context, code int, status string, message string) {
	c.JSON(code, gin.H{
		"status":  status,
		"message": message,
	})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	FormatResponse(c, 200, "success", message, data)
}
