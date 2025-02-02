package services

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/seus31/todo-application-backend/config"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/models"
	"golang.org/x/crypto/bcrypt"
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

func (s *AuthService) Register(ctx context.Context, name, email, hashedPassword string) error {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return s.userRepo.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.userRepo.FindUserByUsername(ctx, username)
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

func (s *AuthService) CheckUserByName(ctx context.Context, name string) bool {
	if _, err := s.userRepo.FindUserByUsername(ctx, name); err == nil {
		return false
	}
	return true
}

func (s *AuthService) CheckUserByEmail(ctx context.Context, email string) bool {
	if _, err := s.userRepo.FindUserByEmail(ctx, email); err == nil {
		return false
	}
	return true
}
