package repository

import (
	"posterr/database/entity"
	"time"

	"gorm.io/gorm"
)

type PostRepositoryInterface interface {
	Get(userID uint64, owner string, page uint64, pageSize uint64, startDateFormated *time.Time, endDateFormated *time.Time) []entity.Post
	Find(postID uint64) *entity.Post
	Insert(post entity.Post) error
	CurrentDayTotalPosts(userID uint64) int64
	CountUserTotalPosts(userID uint64) uint64
	CheckUniqueRepost(postID uint64) bool
	CheckUniqueQuotes(postID uint64) bool
}

type postRepositoryStruct struct {
	DbConn *gorm.DB
}

func (p *postRepositoryStruct) Get(userID uint64, owner string, page uint64, pageSize uint64, startDateFormated *time.Time, endDateFormated *time.Time) []entity.Post {
	offset := int((page - 1) * pageSize)
	var posts []entity.Post
	query := p.DbConn.Model(&entity.Post{})
	if owner == "user" {
		query.Where("user_id = ?", userID)
	}
	if startDateFormated != nil {
		startDate := startDateFormated.Format("2006-01-02 15:04:05")
		query.Where("created_at >= ?", startDate)
	}
	if endDateFormated != nil {
		endDate := endDateFormated.Format("2006-01-02 15:04:05")
		query.Where("created_at <= ?", endDate)
	}
	query.Offset(offset).Limit(int(pageSize)).Find(&posts)

	for index, post := range posts {
		if post.RepostID != nil {
			var repost entity.Post
			p.DbConn.Model(&entity.Post{}).Where("id = ?", post.RepostID).First(&repost)
			posts[index].Repost = &repost
		}

		if post.QuoteID != nil {
			var quote entity.Post
			p.DbConn.Model(&entity.Post{}).Where("id = ?", post.QuoteID).First(&quote)
			posts[index].Quote = &quote
		}
	}

	return posts
}

func (p *postRepositoryStruct) Find(postID uint64) *entity.Post {
	var post entity.Post
	p.DbConn.Model(&entity.Post{}).Where("id = ?", postID).First(&post)
	if post.ID == 0 {
		return nil
	}
	return &post
}

func NewPostRepository(dbConn *gorm.DB) PostRepositoryInterface {
	return &postRepositoryStruct{DbConn: dbConn}
}

func (p *postRepositoryStruct) CurrentDayTotalPosts(userID uint64) int64 {
	currentDate := time.Now().Format("2006-01-02")
	var count int64
	p.DbConn.Model(&entity.Post{}).Where("date(created_at) = ? AND user_id = ?", currentDate, userID).Count(&count)

	return count
}

func (p *postRepositoryStruct) CheckUniqueRepost(postID uint64) bool {
	var count int64
	p.DbConn.Model(&entity.Post{}).Where("id = ? AND repost_id is null", postID).Count(&count)

	return count > 0
}

func (p *postRepositoryStruct) CheckUniqueQuotes(postID uint64) bool {
	var count int64
	p.DbConn.Model(&entity.Post{}).Where("id = ? AND quote_id = 0", postID).Count(&count)

	return count > 0
}

func (p *postRepositoryStruct) Insert(post entity.Post) error {
	result := p.DbConn.Model(&entity.Post{}).Create(&post)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (p *postRepositoryStruct) CountUserTotalPosts(userID uint64) uint64 {
	var count int64
	p.DbConn.Model(&entity.Post{}).Where("user_id = ?", userID).Count(&count)

	return uint64(count)
}
