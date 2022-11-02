package controllers_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"posterr/database/entity"
	"posterr/routes"
	"testing"
	"time"

	postServiceMock "posterr/mocks/services/post"
	userServiceMock "posterr/mocks/services/user"

	"github.com/stretchr/testify/assert"
)

func TestGetUserRoute(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	foundUser := entity.User{
		ID:         1,
		Username:   "Test User",
		Uuid:       "53d83955-16f0-445a-bfe6-a431c1789994",
		Created_at: time.Now().Round(0),
		TotalPosts: 1,
	}

	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, userService)
	userService.On("Find", uuid).Return(foundUser, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserRouteWithError(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"

	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, userService)
	userService.On("Find", uuid).Return(entity.User{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetUserRouteWithoutAuthorization(t *testing.T) {
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
