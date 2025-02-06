package seeders

import (
	"fmt"
	admin_models "github.com/seus31/todo-application-backend/models/admin"
	"github.com/seus31/todo-application-backend/utils"
	"gorm.io/gorm"
)

func AdminSeed(db *gorm.DB) error {
	var count int64

	db.Model(&admin_models.Admin{}).Count(&count)
	if count > 0 {
		return nil
	}

	hashPassword, err := utils.HashPassword("password")
	if err != nil {
		panic("予期せぬエラーが発生しました。")
	}
	admins := admin_models.Admin{Name: "Yamada Taro", Email: "test@sample.com", Password: hashPassword}

	if err := db.Create(&admins).Error; err != nil {
		fmt.Printf("%+v", err)
	}

	return nil
}
