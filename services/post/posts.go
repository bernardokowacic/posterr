package post

import (
	"errors"
	"posterr/database/entity"
	"posterr/repository"
	"time"
)

type PostServiceInterface interface {
	Index(owner string, page uint64, pageSize uint64, startDateFormated time.Time, endDateFormated time.Time) string
	Insert(userUuid string, content string, repost uint64, quotePost uint64) error
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

func (p *postService) Insert(userUuid string, content string, repostID uint64, quotePostID uint64) error {
	user, err := p.UserRepository.GetUserData(userUuid)
	if err != nil {
		return err
	}

	currentDayPostsCount := p.PostRepository.CurrentDayTotalPosts(user.ID)
	if currentDayPostsCount > 5 {
		err := errors.New("maximum number of posts reached")
		return err
	}

	var repost *uint64
	var quotePost *uint64
	if repostID != 0 {
		if !p.PostRepository.CheckUniqueRepost(repostID) {
			err := errors.New("the current reposted post reached the maximum number of reposts")
			return err
		}

		repost = &repostID
	}
	if quotePostID != 0 {
		if !p.PostRepository.CheckUniqueQuotes(repostID) {
			err := errors.New("the current quoted post reached the maximum number of quotes")
			return err
		}

		quotePost = &quotePostID
	}

	post := entity.Post{
		UserID:   user.ID,
		User:     user,
		Content:  content,
		RepostID: repost,
		QuoteID:  quotePost,
	}

	insertErr := p.PostRepository.Insert(post)
	if insertErr != nil {
		return insertErr
	}

	return nil
}
