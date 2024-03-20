package config

import (
	"github.com/IlhamSetiaji/go-lms/internal/http/controller"
	"github.com/IlhamSetiaji/go-lms/internal/http/middleware"
	"github.com/IlhamSetiaji/go-lms/internal/http/route"
	"github.com/IlhamSetiaji/go-lms/internal/messaging"
	"github.com/IlhamSetiaji/go-lms/internal/repository"
	"github.com/IlhamSetiaji/go-lms/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/rabbitmq/amqp091-go"
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
	Producer *amqp091.Channel
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.DB, config.Log)

	// setup producers
	userProducer := messaging.NewUserProducer(config.Producer, config.Log)

	// setup usecases
	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository, userProducer)

	// setup controllers
	userController := controller.NewUserController(config.Log, userUseCase)

	// setup middlewares
	authMiddleware := middleware.NewAuth(config.Config)

	// setup routes
	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.SetupRoutes()
}
