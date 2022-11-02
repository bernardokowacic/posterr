package routes

import (
	"posterr/controllers"
	"posterr/services/post"
	"posterr/services/user"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, postService post.PostServiceInterface, userService user.UserServiceInteface) {
	router.GET("/user", controllers.GetUser(userService))

	router.GET("/posts", controllers.GetPosts(postService))
	router.POST("/post", controllers.InsertPost(postService))
}

func StartAPI(postService post.PostServiceInterface, userService user.UserServiceInteface) *gin.Engine {
	router := gin.Default()

	GetRoutes(router, postService, userService)
	return router
}
