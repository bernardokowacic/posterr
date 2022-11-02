package main

import (
	"log"
	"posterr/database"
	"posterr/repository"
	"posterr/routes"
	"posterr/services/post"
	"posterr/services/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router = gin.Default()

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	dbConn, err := database.CreatePGConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	database.Migrate(dbConn)
	database.Seed(dbConn)

	postRepository := repository.NewPostRepository(dbConn)
	userRepository := repository.NewUserRepository(dbConn)
	postService := post.NewService(postRepository, userRepository)
	userService := user.NewService(postRepository, userRepository)

	routes.GetRoutes(router, postService, userService)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
