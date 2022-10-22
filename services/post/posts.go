package post

import (
	"posterr/database/entity"
	"posterr/repository"
	"time"
)

type PostServiceInterface interface {
	Index(owner string, page uint64, pageSize uint64, startDateFormated time.Time, endDateFormated time.Time) string
	Insert(userUuid string, content entity.Post) (entity.Post, error)
}

type postService struct {
	PostRepository repository.PostRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewPost(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) PostServiceInterface {
	return &postService{PostRepository: postRepository, UserRepository: userRepository}
}

func (p *postService) Index(owner string, page uint64, pageSize uint64, startDateFormated time.Time, endDateFormated time.Time) string {

	return "pong"
}

func (p *postService) Insert(userUuid string, content entity.Post) (entity.Post, error) {
	user, err := p.UserRepository.GetUserData(userUuid)

	if err != nil {
		return entity.Post{}, err
	}
	return entity.Post{}, nil
}
