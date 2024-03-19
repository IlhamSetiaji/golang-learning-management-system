package utils

import "github.com/gin-gonic/gin"

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data,omitempty"`
}

func FormatResponse(c *gin.Context, code int, status string, message string, data interface{}) {
	c.JSON(code, Response{
		Meta: Meta{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, status string, message string) {
	c.JSON(code, Response{
		Meta: Meta{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: nil,
	})
}

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	FormatResponse(c, 200, "success", message, data)
}
