package controllers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"posterr/database/entity"
	postServiceMock "posterr/mocks/services/post"
	userServiceMock "posterr/mocks/services/user"
	"posterr/routes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPostsRoute(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      nil,
		endDate:        nil,
	}

	router := routes.StartAPI(postService, userService)
	postService.On("Index", args.uuid, args.owner, args.pageNumber, args.pageSizeNumber, args.startDate, args.endDate).Return([]entity.Post{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all&page=1&page_size=10", nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPostsRouteWithError(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      nil,
		endDate:        nil,
	}

	router := routes.StartAPI(postService, userService)
	postService.On("Index", args.uuid, args.owner, args.pageNumber, args.pageSizeNumber, args.startDate, args.endDate).Return([]entity.Post{}, errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all&page=1&page_size=10", nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsPageSizeWithstartDate(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	var currentTime *time.Time
	var startTime *time.Time
	formatedTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	currentTime = &formatedTime
	formatedStartTime := currentTime.Add(-time.Hour * 1)
	startTime = &formatedStartTime
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      startTime,
		endDate:        nil,
	}
	url := fmt.Sprintf("/posts?owner=%s&page=%d&page_size=%d&start_date=%s", args.owner, args.pageNumber, args.pageSizeNumber, args.startDate.Format("2006-01-02 15:04:05"))

	router := routes.StartAPI(postService, userService)
	postService.On("Index", args.uuid, args.owner, args.pageNumber, args.pageSizeNumber, args.startDate, args.endDate).Return([]entity.Post{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPostsPageSizeWithstartDateString(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      nil,
		endDate:        nil,
	}
	url := fmt.Sprintf("/posts?owner=%s&page=%d&page_size=%d&start_date=%s", args.owner, args.pageNumber, args.pageSizeNumber, "mydate")

	router := routes.StartAPI(postService, userService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsPageSizeWithendDateString(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      nil,
		endDate:        nil,
	}
	url := fmt.Sprintf("/posts?owner=%s&page=%d&page_size=%d&end_date=%s", args.owner, args.pageNumber, args.pageSizeNumber, "mydate")

	router := routes.StartAPI(postService, userService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsPageSizeWithendDate(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	var currentTime *time.Time
	var endTime *time.Time
	formatedTime, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	currentTime = &formatedTime
	formatedEndTime := currentTime.Add(-time.Hour * 1)
	endTime = &formatedEndTime
	args := struct {
		uuid           string
		owner          string
		pageNumber     uint64
		pageSizeNumber uint64
		startDate      *time.Time
		endDate        *time.Time
	}{
		uuid:           "53d83955-16f0-445a-bfe6-a431c1789994",
		owner:          "all",
		pageNumber:     1,
		pageSizeNumber: 10,
		startDate:      nil,
		endDate:        endTime,
	}
	url := fmt.Sprintf("/posts?owner=%s&page=%d&page_size=%d&end_date=%s", args.owner, args.pageNumber, args.pageSizeNumber, args.endDate.Format("2006-01-02 15:04:05"))

	router := routes.StartAPI(postService, userService)
	postService.On("Index", args.uuid, args.owner, args.pageNumber, args.pageSizeNumber, args.startDate, args.endDate).Return([]entity.Post{}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("Authorization", args.uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetPostsPageSizeString(t *testing.T) {
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all&page=1&page_size=ten", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsPageString(t *testing.T) {
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all&page=one&page_size=10", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsWithoutPageSize(t *testing.T) {
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all&page=1", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsWithoutPage(t *testing.T) {
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts?owner=all", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsWithoutOwner(t *testing.T) {
	uuid := "53d83955-16f0-445a-bfe6-a431c1789994"
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
	req.Header.Set("Authorization", uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestGetPostsWithoutAuthorization(t *testing.T) {
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/posts", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}

func TestInsertPostRoute(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		Uuid      string
		Content   string
		Repost    uint64
		QuotePost uint64
	}{
		Uuid:      "53d83955-16f0-445a-bfe6-a431c1789994",
		Content:   "content",
		Repost:    0,
		QuotePost: 0,
	}
	body, _ := json.Marshal(args)

	router := routes.StartAPI(postService, userService)
	postService.On("Insert", args.Uuid, args.Content, args.Repost, args.QuotePost).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewReader(body))
	req.Header.Set("Authorization", args.Uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestInsertPostRouteWithError(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		Uuid      string
		Content   string
		Repost    uint64
		QuotePost uint64
	}{
		Uuid:      "53d83955-16f0-445a-bfe6-a431c1789994",
		Content:   "content",
		Repost:    0,
		QuotePost: 0,
	}
	body, _ := json.Marshal(args)

	router := routes.StartAPI(postService, userService)
	postService.On("Insert", args.Uuid, args.Content, args.Repost, args.QuotePost).Return(errors.New("error"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewReader(body))
	req.Header.Set("Authorization", args.Uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestInsertPostRouteWithoutBody(t *testing.T) {
	userService := &userServiceMock.UserServiceInteface{}
	postService := &postServiceMock.PostServiceInterface{}
	args := struct {
		Uuid         string
		Content      string
		Repost       uint64
		AnotherThing uint64
	}{
		Uuid:         "53d83955-16f0-445a-bfe6-a431c1789994",
		Content:      "content",
		Repost:       0,
		AnotherThing: 0,
	}
	body, _ := json.Marshal(args)

	router := routes.StartAPI(postService, userService)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/post", bytes.NewReader(body))
	req.Header.Set("Authorization", args.Uuid)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotAcceptable, w.Code)
}

func TestInsertPostWithoutAuthorization(t *testing.T) {
	router := routes.StartAPI(&postServiceMock.PostServiceInterface{}, &userServiceMock.UserServiceInteface{})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/post", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusForbidden, w.Code)
}
