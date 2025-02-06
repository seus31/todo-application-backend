package admin_services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/config"
	"github.com/seus31/todo-application-backend/dto/requests/admin/auth"
	admin_interfaces "github.com/seus31/todo-application-backend/interfaces/admin"
	"github.com/seus31/todo-application-backend/services"
	"github.com/seus31/todo-application-backend/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	adminRepo admin_interfaces.AdminRepositoryInterface
}

func NewAuthService(repo admin_interfaces.AdminRepositoryInterface) *AuthService {
	return &AuthService{
		adminRepo: repo,
	}
}

func (s *AuthService) Login(ctx *fiber.Ctx) (string, error) {
	var req admin_auth.LoginRequest
	contextData := utils.GetContextFromFiber(ctx)

	if err := ctx.BodyParser(&req); err != nil {
		return "", services.ErrFailedToParseRequest
	}

	admin, err := s.adminRepo.FindAdminByEmail(contextData, req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", services.ErrInvalidCredentials
		}
		return "", services.ErrUnexpectedError
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(req.Password)); err != nil {
		return "", services.ErrInvalidCredentials
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = admin.ID
	claims["email"] = admin.Email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET_KEY")))
	if err != nil {
		return "", services.ErrUnexpectedError
	}

	return t, nil
}
