package routes

import (
	"posterr/controllers"
	"posterr/repository"

	"github.com/gin-gonic/gin"
)

func GetRoutes(router *gin.Engine, postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) {
	router.GET("/user", controllers.GetUser(postRepository, userRepository))

	router.GET("/posts", controllers.GetPosts(postRepository, userRepository))
	router.POST("/post", controllers.InsertPost(postRepository, userRepository))
}
