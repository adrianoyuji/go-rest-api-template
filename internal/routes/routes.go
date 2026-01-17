package routes

import (
	"github.com/adrianoyuji/go-rest-api-template/internal/controllers"
	"github.com/adrianoyuji/go-rest-api-template/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	authController := controllers.NewAuthController(db)
	userController := controllers.NewUserController(db)

	r.POST("/auth/register", authController.Register)
	r.POST("/auth/login", authController.Login)

	protected := r.Group("/api")
	protected.Use(middleware.AuthGuard())
	{
		protected.GET("/users", userController.ListUsers)
	}

	return r
}
