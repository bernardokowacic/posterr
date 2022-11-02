package post_test

import (
	"errors"
	"posterr/database/entity"
	mockrepository "posterr/mocks/repository"
	"posterr/repository"
	"posterr/services/post"
	"reflect"
	"testing"
	"time"
)

func Test_postService_Index(t *testing.T) {
	currentTime := time.Now().Round(0)
	startTime := currentTime.Add(-time.Hour * 1).Round(0)
	endTime := currentTime.Add(time.Hour * 1).Round(0)
	foundUser := entity.User{
		ID:         1,
		Username:   "Test User",
		Uuid:       "53d83955-16f0-445a-bfe6-a431c1789994",
		Created_at: currentTime,
		TotalPosts: 1,
	}
	type fields struct {
		PostRepository repository.PostRepositoryInterface
		UserRepository repository.UserRepositoryInterface
	}
	type args struct {
		userUuid          string
		owner             string
		page              uint64
		pageSize          uint64
		startDateFormated *time.Time
		endDateFormated   *time.Time
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		want           []entity.Post
		wantErr        bool
	}{
		{
			name: "Without time filter",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:          "53d83955-16f0-445a-bfe6-a431c1789994",
				owner:             "all",
				page:              uint64(1),
				pageSize:          uint64(10),
				startDateFormated: nil,
				endDateFormated:   nil,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Get", foundUser.ID, a.owner, a.page, a.pageSize, a.startDateFormated, a.endDateFormated).Return([]entity.Post{
					{
						ID:     1,
						UserID: 1,
					},
					{
						ID:     2,
						UserID: 1,
					},
				}, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			want: []entity.Post{
				{
					ID:     1,
					UserID: 1,
				},
				{
					ID:     2,
					UserID: 1,
				},
			},
			wantErr: false,
		},
		{
			name: "With time filter",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:          "53d83955-16f0-445a-bfe6-a431c1789994",
				owner:             "all",
				page:              uint64(1),
				pageSize:          uint64(10),
				startDateFormated: &startTime,
				endDateFormated:   &endTime,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Get", foundUser.ID, a.owner, a.page, a.pageSize, a.startDateFormated, a.endDateFormated).Return([]entity.Post{
					{
						ID:         2,
						UserID:     1,
						Created_at: currentTime,
					},
				}, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			want: []entity.Post{
				{
					ID:         2,
					UserID:     1,
					Created_at: currentTime,
				},
			},
			wantErr: false,
		},
		{
			name: "SQL Error Get user data",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:          "53d83955-16f0-445a-bfe6-a431c1789994",
				owner:             "all",
				page:              uint64(1),
				pageSize:          uint64(10),
				startDateFormated: &startTime,
				endDateFormated:   &endTime,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "SQL Error get posts",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:          "53d83955-16f0-445a-bfe6-a431c1789994",
				owner:             "all",
				page:              uint64(1),
				pageSize:          uint64(10),
				startDateFormated: &startTime,
				endDateFormated:   &endTime,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Get", foundUser.ID, a.owner, a.page, a.pageSize, a.startDateFormated, a.endDateFormated).Return([]entity.Post{}, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "page is 0",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:          "53d83955-16f0-445a-bfe6-a431c1789994",
				owner:             "all",
				page:              uint64(0),
				pageSize:          uint64(10),
				startDateFormated: nil,
				endDateFormated:   nil,
			},
			mockBehavior: func(f fields, a args) {
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := post.NewService(tt.fields.PostRepository, tt.fields.UserRepository)

			got, err := service.Index(tt.args.userUuid, tt.args.owner, tt.args.page, tt.args.pageSize, tt.args.startDateFormated, tt.args.endDateFormated)
			if (err != nil) != tt.wantErr {
				t.Errorf("postService.Index() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postService.Index() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}

func Test_postService_Insert(t *testing.T) {
	currentTime := time.Now().Round(0)
	id := uint64(1)
	foundUser := entity.User{
		ID:         1,
		Username:   "Test User",
		Uuid:       "53d83955-16f0-445a-bfe6-a431c1789994",
		Created_at: currentTime,
		TotalPosts: 0,
	}
	postInsert := entity.Post{
		UserID:   1,
		User:     foundUser,
		Content:  "content",
		RepostID: nil,
		QuoteID:  nil,
	}
	repost := entity.Post{
		UserID:   1,
		User:     foundUser,
		Content:  "content",
		RepostID: &id,
		QuoteID:  &id,
	}
	type fields struct {
		PostRepository repository.PostRepositoryInterface
		UserRepository repository.UserRepositoryInterface
	}
	type args struct {
		userUuid    string
		content     string
		repostID    uint64
		quotePostID uint64
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		wantErr        bool
	}{
		{
			name: "Add new post",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    0,
				quotePostID: 0,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Insert", postInsert).Return(nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: false,
		},
		{
			name: "Add new post get user data error",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    0,
				quotePostID: 0,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(entity.User{}, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post more than 5 posts in a day",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    0,
				quotePostID: 0,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(6), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post currentDayTotalPosts error",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    0,
				quotePostID: 0,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post with repost",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueQuotes", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Insert", repost).Return(nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: false,
		},
		{
			name: "Add new post CheckUniqueRepost error",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(false, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post with repeated repost",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(false, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post with repeated quote",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueQuotes", a.repostID).Return(false, nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post with repeated quote",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueQuotes", a.repostID).Return(false, errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
		{
			name: "Add new post insert error",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid:    "53d83955-16f0-445a-bfe6-a431c1789994",
				content:     "content",
				repostID:    1,
				quotePostID: 1,
			},
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CurrentDayTotalPosts", foundUser.ID).Return(int64(0), nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueRepost", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CheckUniqueQuotes", a.repostID).Return(true, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("Insert", repost).Return(errors.New("error"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := post.NewService(tt.fields.PostRepository, tt.fields.UserRepository)

			if err := service.Insert(tt.args.userUuid, tt.args.content, tt.args.repostID, tt.args.quotePostID); (err != nil) != tt.wantErr {
				t.Errorf("postService.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}
