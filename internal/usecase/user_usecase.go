package usecase

import (
	"context"
	"time"

	"github.com/IlhamSetiaji/go-lms/internal/entity"
	"github.com/IlhamSetiaji/go-lms/internal/messaging"
	"github.com/IlhamSetiaji/go-lms/internal/model"
	"github.com/IlhamSetiaji/go-lms/internal/repository"
	"github.com/IlhamSetiaji/go-lms/internal/request"
	"github.com/IlhamSetiaji/go-lms/internal/response"
	"github.com/IlhamSetiaji/go-lms/utils"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCase struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository repository.UserRepositoryInterface
	Producer       *messaging.EmailProducer
	Viper          *viper.Viper
}

type UserUseCaseInterface interface {
	Login(ctx context.Context, request *request.UserLoginRequest) (*response.UserLoginResponse, error)
	Me(ctx context.Context, userId uint) (*response.UserMeResponse, error)
	Register(ctx context.Context, request *request.RegisterUserRequest) (*response.UserMeResponse, error)
}

func NewUserUseCase(db *gorm.DB, log *logrus.Logger, validate *validator.Validate, userRepository repository.UserRepositoryInterface, producer *messaging.EmailProducer, viper *viper.Viper) UserUseCaseInterface {
	return &UserUseCase{
		DB:             db,
		Log:            log,
		Validate:       validate,
		UserRepository: userRepository,
		Producer:       producer,
		Viper:          viper,
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
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
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
		Expires:   time.Now().Add(time.Hour * 72).Unix(),
	}, nil
}

func (c *UserUseCase) Me(ctx context.Context, userId uint) (*response.UserMeResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	user, err := c.UserRepository.GetMyProfile(tx, &entity.User{}, userId)
	if err != nil {
		c.Log.Errorf("Error when finding user by id: %v", err)
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}
	var roles []response.RoleResponse
	for _, role := range user.Roles {
		roles = append(roles, response.RoleResponse{
			ID:   role.ID,
			Name: role.Name,
		})
	}
	return &response.UserMeResponse{
		ID:              user.ID,
		Username:        user.Username,
		Email:           user.Email,
		Name:            user.Name,
		EmailVerifiedAt: user.EmailVerifiedAt.String(),
		Roles:           roles,
	}, nil
}

func (c *UserUseCase) Register(ctx context.Context, request *request.RegisterUserRequest) (*response.UserMeResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	err := c.Validate.Struct(request)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Errorf("Error when hashing password: %v", err)
		return nil, err
	}
	request.Password = string(hashedPassword)

	user, err := c.UserRepository.CreateUser(tx, &entity.User{}, request)
	if err != nil {
		c.Log.Errorf("Error when creating user: %v", err)
		return nil, err
	}
	err = c.UserRepository.AssignRole(tx, user.ID, request.RoleID)

	if err != nil {
		c.Log.Errorf("Error when assigning role: %v", err)
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	// Publish an EmailEvent to RabbitMQ
	emailEvent := &model.EmailEvent{
		From:    "your-email@gmail.com",
		To:      user.Email,
		Subject: "User Registration Notification",
		Body:    "A new user has been registered: " + user.Username,
	}
	if err = c.Producer.Send(ctx, emailEvent); err != nil {
		c.Log.Errorf("Error when sending email event: %v", err)
		return nil, err
	}

	return c.Me(ctx, user.ID)
}
