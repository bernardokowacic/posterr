package user

import (
	"posterr/database/entity"
	"posterr/repository"
)

type UserServiceInteface interface {
	Find(userUuid string) (entity.User, error)
}

type UserService struct {
	PostRepository repository.PostRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewService(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) UserServiceInteface {
	return &UserService{PostRepository: postRepository, UserRepository: userRepository}
}

func (u *UserService) Find(userUuid string) (entity.User, error) {
	var err error
	user, err := u.UserRepository.GetUserData(userUuid)
	if err != nil {
		return entity.User{}, err
	}
	user.TotalPosts, err = u.PostRepository.CountUserTotalPosts(user.ID)
	if err != nil {
		return user, err
	}

	return user, nil
}
