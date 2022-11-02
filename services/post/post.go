package post

import (
	"errors"
	"posterr/database/entity"
	"posterr/repository"
	"time"
)

type PostServiceInterface interface {
	Index(userUuid string, owner string, page uint64, pageSize uint64, startDateFormated *time.Time, endDateFormated *time.Time) ([]entity.Post, error)
	Insert(userUuid string, content string, repost uint64, quotePost uint64) error
}

type postService struct {
	PostRepository repository.PostRepositoryInterface
	UserRepository repository.UserRepositoryInterface
}

func NewService(postRepository repository.PostRepositoryInterface, userRepository repository.UserRepositoryInterface) PostServiceInterface {
	return &postService{PostRepository: postRepository, UserRepository: userRepository}
}

func (p *postService) Index(userUuid string, owner string, page uint64, pageSize uint64, startDateFormated *time.Time, endDateFormated *time.Time) ([]entity.Post, error) {
	if page == 0 {
		return nil, errors.New("page can not be 0")
	}
	user, err := p.UserRepository.GetUserData(userUuid)
	if err != nil {
		return nil, err
	}

	posts, err := p.PostRepository.Get(user.ID, owner, page, pageSize, startDateFormated, endDateFormated)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (p *postService) Insert(userUuid string, content string, repostID uint64, quotePostID uint64) error {
	user, err := p.UserRepository.GetUserData(userUuid)
	if err != nil {
		return err
	}

	currentDayPostsCount, err := p.PostRepository.CurrentDayTotalPosts(user.ID)
	if err != nil {
		return err
	}
	if currentDayPostsCount > 5 {
		err := errors.New("maximum number of posts reached")
		return err
	}

	var repost *uint64
	var quotePost *uint64
	if repostID != 0 {
		uniqueRepost, err := p.PostRepository.CheckUniqueRepost(repostID)
		if err != nil {
			return err
		}
		if !uniqueRepost {
			err := errors.New("the current reposted post reached the maximum number of reposts")
			return err
		}

		repost = &repostID
	}
	if quotePostID != 0 {
		uniqueQuote, err := p.PostRepository.CheckUniqueQuotes(repostID)
		if err != nil {
			return err
		}
		if !uniqueQuote {
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
