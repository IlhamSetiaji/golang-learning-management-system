package config

import (
	"github.com/IlhamSetiaji/go-lms/internal/http/controller"
	"github.com/IlhamSetiaji/go-lms/internal/http/route"
	"github.com/IlhamSetiaji/go-lms/internal/messaging"
	"github.com/IlhamSetiaji/go-lms/internal/repository"
	"github.com/IlhamSetiaji/go-lms/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *gin.Engine
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *messaging.Producer
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Log)

	// setup usecases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, config.Producer)

	// setup controllers
	userController := controller.NewUserController(config.Log, userUseCase)

	// setup routes
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.SetupRoutes()
}
