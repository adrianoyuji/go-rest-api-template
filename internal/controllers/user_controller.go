package controllers

import (
	"net/http"

	"github.com/adrianoyuji/go-rest-api-template/internal/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	Repo *repositories.UserRepository
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{Repo: repositories.NewUserRepository(db)}
}

func (u *UserController) ListUsers(c *gin.Context) {
	users, err := u.Repo.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": users})
}
