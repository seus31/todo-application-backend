package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/seus31/todo-application-backend/database/seeders"
	"github.com/seus31/todo-application-backend/middleware"
	"github.com/seus31/todo-application-backend/routes"
	admin_routes "github.com/seus31/todo-application-backend/routes/admin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	dsn := "host=db user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Tokyo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("データベース接続に失敗しました")
	}

	err = seeders.AdminSeed(db)
	if err != nil {
		panic("シーダーの実行に失敗しました")
	}

	api := app.Group("/api")
	adminAuth := api.Group("/admin/auth")
	auth := api.Group("/auth")
	v1 := api.Group("/v1")
	v1.Use(middleware.AuthMiddleware)

	admin := v1.Group("/admin")
	users := admin.Group("/users")

	user := v1.Group("/user")
	tasks := v1.Group("/tasks")
	categories := v1.Group("/categories")

	// Admin Routes
	admin_routes.SetUpAuthRoutes(adminAuth, db)
	admin_routes.SetUpUserRoutes(users, db)

	// User Routes
	routes.SetUpAuthRoutes(auth, db)
	routes.SetUpUserInfoRoutes(user, db)
	routes.SetUpTaskRoutes(tasks, db)
	admin_routes.SetUpUserRoutes(users, db)
	routes.SetUpCategoryRoutes(categories, db)

	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
