package controllers

import (
	"net/http"
	"posterr/database/entity"
	"posterr/repository"
	"posterr/services/post"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		owner, exists := c.GetQuery("owner")
		if !exists || (owner != "all" && owner != "user") {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "Owner param is required and must be 'all' or 'user'"})
			return
		}
		page, exists := c.GetQuery("page")
		if !exists {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "You must inform the page"})
			return
		}
		pageSize, exists := c.GetQuery("page_size")
		if !exists {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "You must inform the quantity of items in page_size"})
			return
		}
		pageSizeNumber, err := strconv.ParseUint(pageSize, 0, 64)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "page_size must be an unsigned integer"})
			return
		}
		startDate := c.Query("start_date")
		endDate := c.Query("end_date")

		startDateFormated, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "Start date must be a date"})
			return
		}
		endDateFormated, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "End date must be a date"})
			return
		}

		pageNumber, err := strconv.ParseUint(page, 0, 64)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "Page must be an unsigned integer"})
			return
		}

		postService := post.NewPost(postRepository, userRepository)
		postService.Index(owner, pageNumber, pageSizeNumber, startDateFormated, endDateFormated)

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}

func InsertPost(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		var postData entity.Post
		c.BindJSON(&postData)
		userUiid := c.Request.Header["Authorization"][0]

		postService := post.NewPost(postRepository, userRepository)
		postService.Insert(userUiid, postData)

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	}
}
