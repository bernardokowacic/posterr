package controllers

import (
	"net/http"
	"posterr/repository"
	"posterr/services/post"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPosts(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := struct {
			Authorization *string `header:"Authorization" binding:"required"`
		}{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization is required"})
			return
		}

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
		var startDateFormated *time.Time
		startDate := c.Query("start_date")
		if startDate != "" {
			date, err := time.Parse("2006-01-02 15:04:05", startDate)
			startDateFormated = &date
			if err != nil {
				c.JSON(http.StatusNotAcceptable, gin.H{"message": "Start date must be a date"})
				return
			}
		}
		var endDateFormated *time.Time
		endDate := c.Query("end_date")
		if endDate != "" {
			date, err := time.Parse("2006-01-02 15:04:05", endDate)
			endDateFormated = &date
			if err != nil {
				c.JSON(http.StatusNotAcceptable, gin.H{"message": "End date must be a date"})
				return
			}
		}

		pageNumber, err := strconv.ParseUint(page, 0, 64)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": "Page must be an unsigned integer"})
			return
		}

		postService := post.NewPost(postRepository, userRepository)
		posts, err := postService.Index(*header.Authorization, owner, pageNumber, pageSizeNumber, startDateFormated, endDateFormated)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, posts)
	}
}

func InsertPost(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := struct {
			Authorization *string `header:"Authorization" binding:"required"`
		}{}
		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "Authorization is required"})
			return
		}

		postData := struct {
			Content   string
			Repost    uint64
			QuotePost uint64
		}{}
		err := c.BindJSON(&postData)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
			return
		}

		postService := post.NewPost(postRepository, userRepository)
		insertErr := postService.Insert(*header.Authorization, postData.Content, postData.Repost, postData.QuotePost)
		if insertErr != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"message": insertErr.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	}
}
