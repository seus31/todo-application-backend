package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/seus31/todo-application-backend/middleware"
	"github.com/seus31/todo-application-backend/routes"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	dsn := "host=db user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("データベース接続に失敗しました")
	}

	api := app.Group("/api")
	auth := api.Group("/auth")
	v1 := api.Group("/v1")
	v1.Use(middleware.AuthMiddleware)

	users := v1.Group("/users")
	tasks := v1.Group("/tasks")
	categories := v1.Group("/categories")

	routes.SetUpAuthRoutes(auth, db)
	routes.SetUpTaskRoutes(tasks, db)
	routes.SetUpUserRoutes(users, db)
	routes.SetUpCategoryRoutes(categories, db)

	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
