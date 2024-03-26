package controller

import (
	"github.com/IlhamSetiaji/go-lms/internal/http/middleware"
	"github.com/IlhamSetiaji/go-lms/internal/request"
	"github.com/IlhamSetiaji/go-lms/internal/usecase"
	"github.com/IlhamSetiaji/go-lms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase usecase.UserUseCaseInterface
	Viper   *viper.Viper
}

func NewUserController(log *logrus.Logger, useCase usecase.UserUseCaseInterface) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *UserController) Login(ctx *gin.Context) {
	payload := new(request.UserLoginRequest)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		c.Log.Errorf("Error when binding request: %v", err)
		utils.ErrorResponse(ctx, 400, "error", err.Error())
		return
	}
	response, err := c.UseCase.Login(ctx, payload)
	if err != nil {
		c.Log.Errorf("Error when login: %v", err)
		utils.ErrorResponse(ctx, 500, "error", err.Error())
		return
	}
	utils.SuccessResponse(ctx, "Login success", response)
}

func (c *UserController) Me(ctx *gin.Context) {
	user := middleware.GetUser(ctx)
	if user == nil {
		utils.ErrorResponse(ctx, 401, "error", "Unauthorized")
		return
	}
	// userId := fmt.Sprintf("%v", user["userId"])
	userId := user["userId"].(uint)
	response, err := c.UseCase.Me(ctx, userId)
	if err != nil {
		c.Log.Errorf("Error when getting user: %v", err)
		utils.ErrorResponse(ctx, 500, "error", err.Error())
		return
	}
	utils.SuccessResponse(ctx, "Success", response)
}

func (c *UserController) Register(ctx *gin.Context) {
	payload := new(request.RegisterUserRequest)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		c.Log.Errorf("Error when binding request: %v", err)
		utils.ErrorResponse(ctx, 400, "error", err.Error())
		return
	}
	response, err := c.UseCase.Register(ctx, payload)
	if err != nil {
		c.Log.Errorf("Error when registering user: %v", err)
		utils.ErrorResponse(ctx, 500, "error", err.Error())
		return
	}
	utils.SuccessResponse(ctx, "User registered", response)
}
