package controllers

import (
	"net/http"
	"posterr/services/user"

	"github.com/gin-gonic/gin"
)

func GetUser(userService user.UserServiceInteface) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := struct {
			Authorization *string `header:"Authorization" binding:"required"`
		}{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization is required"})
			return
		}

		user, err := userService.Find(*header.Authorization)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
