package user_test

import (
	"errors"
	"posterr/database/entity"
	mockrepository "posterr/mocks/repository"
	"posterr/repository"
	"posterr/services/user"
	"reflect"
	"testing"
	"time"
)

func TestUserService_Find(t *testing.T) {
	currentTime := time.Now().Round(0)
	foundUser := entity.User{
		ID:         1,
		Username:   "Test User",
		Uuid:       "53d83955-16f0-445a-bfe6-a431c1789994",
		Created_at: currentTime,
		TotalPosts: 1,
	}
	notFoundUser := entity.User{
		ID:         0,
		Username:   "",
		Uuid:       "",
		Created_at: currentTime,
		TotalPosts: 0,
	}
	type fields struct {
		PostRepository repository.PostRepositoryInterface
		UserRepository repository.UserRepositoryInterface
	}
	type args struct {
		userUuid string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		want           entity.User
		mockBehavior   func(f fields, a args)
		assertBehavior func(t *testing.T, f fields)
		wantErr        bool
	}{
		{
			name: "Valid user",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid: "53d83955-16f0-445a-bfe6-a431c1789994",
			},
			want:    foundUser,
			wantErr: false,
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CountUserTotalPosts", foundUser.ID).Return(uint64(1), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
		},
		{
			name: "Invalid user",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid: "this-is-a-uuid-trust-me-bro",
			},
			want:    notFoundUser,
			wantErr: false,
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(notFoundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CountUserTotalPosts", notFoundUser.ID).Return(uint64(0), nil)
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
		},
		{
			name: "SQL Error Get user data",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid: "53d83955-16f0-445a-bfe6-a431c1789994",
			},
			want:    entity.User{},
			wantErr: true,
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, errors.New("err"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
		},
		{
			name: "SQL Error count user posts",
			fields: fields{
				PostRepository: &mockrepository.PostRepositoryInterface{},
				UserRepository: &mockrepository.UserRepositoryInterface{},
			},
			args: args{
				userUuid: "53d83955-16f0-445a-bfe6-a431c1789994",
			},
			want: entity.User{
				ID:         1,
				Username:   "Test User",
				Uuid:       "53d83955-16f0-445a-bfe6-a431c1789994",
				Created_at: currentTime,
				TotalPosts: 0,
			},
			wantErr: true,
			mockBehavior: func(f fields, a args) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).On("GetUserData", a.userUuid).Return(foundUser, nil)
				f.PostRepository.(*mockrepository.PostRepositoryInterface).On("CountUserTotalPosts", foundUser.ID).Return(uint64(0), errors.New("err"))
			},
			assertBehavior: func(t *testing.T, f fields) {
				f.UserRepository.(*mockrepository.UserRepositoryInterface).AssertExpectations(t)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mockBehavior != nil {
				tt.mockBehavior(tt.fields, tt.args)
			}

			service := user.NewService(tt.fields.PostRepository, tt.fields.UserRepository)

			got, err := service.Find(tt.args.userUuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.Find() = %v, want %v", got, tt.want)
			}

			if tt.assertBehavior != nil {
				tt.assertBehavior(t, tt.fields)
			}
		})
	}
}
