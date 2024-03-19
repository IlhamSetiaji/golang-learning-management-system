package route

import (
	"github.com/IlhamSetiaji/go-lms/internal/http/controller"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController *controller.UserController
}

func (c *RouteConfig) SetupRoutes() {
	// user routes
	c.App.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	c.SetupAuthRoutes()
}

func (c *RouteConfig) SetupAuthRoutes() {
	userRoutes := c.App.Group("/auth")
	{
		userRoutes.POST("/login", c.UserController.Login)
	}
}
