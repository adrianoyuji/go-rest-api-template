package controllers

import (
	"net/http"
	"os"

	"github.com/adrianoyuji/go-rest-api-template/internal/models"
	"github.com/adrianoyuji/go-rest-api-template/internal/repositories"
	"github.com/adrianoyuji/go-rest-api-template/internal/services"
	"github.com/adrianoyuji/go-rest-api-template/pkg/jwt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	Service *services.AuthService
	Repo    *repositories.UserRepository
	DB      *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	repo := repositories.NewUserRepository(db)
	return &AuthController{
		Service: services.NewAuthService(repo),
		Repo:    repo,
		DB:      db,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := a.Service.Register(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"user": input})
}

func (a *AuthController) Login(c *gin.Context) {
	var creds struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := a.Service.Authenticate(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	secret := os.Getenv("JWT_SECRET")
	token, err := jwt.GenerateToken(user.ID, secret, 24*7) // 7 days
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
