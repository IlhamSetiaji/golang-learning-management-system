package usecase

import (
	"context"

	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/IlhamSetiaji/go-lms/internal/messaging"
	"github.com/IlhamSetiaji/go-lms/internal/repository"
	"github.com/IlhamSetiaji/go-lms/internal/request"
	"github.com/IlhamSetiaji/go-lms/internal/response"
	"github.com/IlhamSetiaji/go-lms/utils"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepositoryInterface
	Producer       *messaging.UserProducer
}

type UserUseCaseInterface interface {
	Login(ctx context.Context, request *request.UserLoginRequest) (*response.UserLoginResponse, error)
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository repository.UserRepositoryInterface, producer *messaging.UserProducer) UserUseCaseInterface {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		UserRepository: userRepository,
		Producer:       producer,
	}
}

func (c *UserUseCase) Login(ctx context.Context, request *request.UserLoginRequest) (*response.UserLoginResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	err := c.Validate.Struct(request)
	if err != nil {
		return nil, err
	}
	user, err := c.UserRepository.FindFirstByField(tx, &entity.User{}, "username", request.Username)
	if err != nil {
		c.Log.Errorf("Error when finding user by username: %v", err)
		return nil, err
	}
	if user == nil {
		c.Log.Errorf("User not found")
		return nil, err
	}
	checkedPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if checkedPassword != nil {
		c.Log.Errorf("Error when comparing password: %v", checkedPassword)
		return nil, checkedPassword
	}
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.Log.Errorf("Error when generating token: %v", err)
		return nil, err
	}
	return &response.UserLoginResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		TokenType: "Bearer",
		Token:     token,
	}, nil
}
