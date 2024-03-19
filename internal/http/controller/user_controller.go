package controller

import (
	"github.com/IlhamSetiaji/go-lms/internal/usecase"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(log *logrus.Logger, useCase *usecase.UserUseCase) *UserController {
	return &UserController{
		Log:     log,
		UseCase: useCase,
	}
}

func (c *UserController) Login() {}
