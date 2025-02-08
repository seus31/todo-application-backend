package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/dto/responses"
	"github.com/seus31/todo-application-backend/interfaces"
	"github.com/seus31/todo-application-backend/utils"
	"gorm.io/gorm"
)

type UserInfoService struct {
	userRepo interfaces.UserRepositoryInterface
}

func NewUserInfoService(repo interfaces.UserRepositoryInterface) *UserInfoService {
	return &UserInfoService{
		userRepo: repo,
	}
}

func (s *UserInfoService) GetUserInfo(ctx *fiber.Ctx) (*responses.UserInfoResponse, error) {
	contextData := utils.GetContextFromFiber(ctx)
	user, err := s.userRepo.GetUserByID(contextData, ctx.Locals("userID").(uint))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrRecordNotFound
		}
		return nil, ErrUnexpectedError
	}

	response := &responses.UserInfoResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	return response, nil
}
