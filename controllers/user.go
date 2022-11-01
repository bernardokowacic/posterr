package controllers

import (
	"net/http"
	"posterr/repository"
	"posterr/services/user"

	"github.com/gin-gonic/gin"
)

func GetUser(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := struct {
			Authorization *string `header:"Authorization" binding:"required"`
		}{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization is required"})
			return
		}

		userService := user.NewUser(postRepository, userRepository)
		user := userService.Find(*header.Authorization)

		c.JSON(http.StatusOK, user)
	}
}
