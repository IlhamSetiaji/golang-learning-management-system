package controller

import (
	"github.com/IlhamSetiaji/go-lms/internal/request"
	"github.com/IlhamSetiaji/go-lms/internal/usecase"
	"github.com/IlhamSetiaji/go-lms/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase usecase.UserUseCaseInterface
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
