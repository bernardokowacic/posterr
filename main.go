package main

import (
	"log"
	"posterr/database"
	"posterr/repository"
	"posterr/routes"
	"posterr/services/post"
	"posterr/services/user"

	"github.com/joho/godotenv"
)

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

	router := routes.StartAPI(postService, userService)
	router.Run()
}
