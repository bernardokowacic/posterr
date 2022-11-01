package user

import (
	"posterr/database/entity"
	"posterr/repository"
)

type UserServiceInteface interface {
	Find(userUuid string) entity.User
}

type UserService struct {
	PostRepository repository.PostRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewUser(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) UserServiceInteface {
	return &UserService{PostRepository: postRepository, UserRepository: userRepository}
}

func (u *UserService) Find(userUuid string) entity.User {
	user := u.UserRepository.GetUserData(userUuid)
	user.TotalPosts = u.PostRepository.CountUserTotalPosts(user.ID)

	return user
}
