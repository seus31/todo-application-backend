package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/services"
)

type UserInfoController struct {
	UserInfoService *services.UserInfoService
}

func NewUserInfoController(userInfoService *services.UserInfoService) *UserInfoController {
	return &UserInfoController{
		UserInfoService: userInfoService,
	}
}

func (uic *UserInfoController) Info(ctx *fiber.Ctx) error {
	userInfo, err := uic.UserInfoService.GetUserInfo(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	return ctx.Status(fiber.StatusOK).JSON(userInfo)
}
