package services

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/config"
	"github.com/seus31/todo-application-backend/dto/requests/users"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
	"github.com/seus31/todo-application-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	userRepo interfaces.UserRepositoryInterface
}

func NewAuthService(repo interfaces.UserRepositoryInterface) *AuthService {
	return &AuthService{
		userRepo: repo,
	}
}

func (s *AuthService) Register(ctx *fiber.Ctx) error {
	var req users.CreateUserRequest
	contextData := utils.GetContextFromFiber(ctx)

	if err := ctx.BodyParser(&req); err != nil {
		return ErrFailedToParseRequest
	}

	validate := users.CreateUserRequestValidator()
	if err := utils.ValidateStruct(validate, &req); err != nil {
		return err
	}

	_, err := s.userRepo.FindUserByName(contextData, req.Name)

	if err == nil {
		return ErrDuplicateName
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrFailedToRegisterUser
	}

	_, err = s.userRepo.FindUserByEmail(contextData, req.Email)
	if err == nil {
		return ErrDuplicateEmail
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrFailedToRegisterUser
	}

	if utils.CheckPasswordAndConfirmPassword(req.Password, req.ConfirmPassword) == false {
		return ErrPasswordMismatch
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return ErrFailedToHashPassword
	}

	user := &models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashPassword,
	}

	err = s.userRepo.Create(contextData, user)
	if err != nil {
		return ErrFailedToRegisterUser
	}

	return nil
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindUserByName(ctx, username)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Name
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return t, nil
}
