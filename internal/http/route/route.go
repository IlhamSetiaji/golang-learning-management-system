package route

import (
	"github.com/IlhamSetiaji/go-lms/internal/http/controller"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	App            *gin.Engine
	UserController *controller.UserController
	AuthMiddleware gin.HandlerFunc
}

func (c *RouteConfig) SetupRoutes() {
	// user routes
	c.App.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	c.SetupAuthRoutes()
	c.SetupMustAuthRoutes()
}

func (c *RouteConfig) SetupAuthRoutes() {
	authRoutes := c.App.Group("/auth")
	{
		authRoutes.POST("/login", c.UserController.Login)
		authRoutes.POST("/register", c.UserController.Register)
	}
}

func (c *RouteConfig) SetupMustAuthRoutes() {
	userRoutes := c.App.Group("/users")
	userRoutes.Use(c.AuthMiddleware)
	{
		userRoutes.GET("/profile", c.UserController.Me)
	}
}
