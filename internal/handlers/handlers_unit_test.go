package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ssitko/hex-domain/internal/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAlbumService is a mock implementation of the AlbumService interface (contained in the Album domain)
type MockAlbumService struct {
	mock.Mock
}

func (m *MockAlbumService) GetAllAlbums() ([]domain.Album, error) {
	args := m.Called()
	return args.Get(0).([]domain.Album), args.Error(1)
}

func (m *MockAlbumService) GetAlbumByID(id int) (domain.Album, error) {
	args := m.Called(id)
	return args.Get(0).(domain.Album), args.Error(1)
}

func (m *MockAlbumService) CreateAlbum(album domain.Album) (domain.Album, error) {
	args := m.Called(album)
	return args.Get(0).(domain.Album), args.Error(1)
}

func (m *MockAlbumService) UpdateAlbum(album domain.Album) (domain.Album, error) {
	args := m.Called(album)
	return args.Get(0).(domain.Album), args.Error(1)
}

func (m *MockAlbumService) DeleteAlbum(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func setupTestRouter(service *MockAlbumService) *gin.Engine {
	r := gin.Default()
	handler := NewAlbumHandler(service)
	r.GET("/albums", handler.GetAlbums)
	r.GET("/albums/:id", handler.GetAlbumByID)
	r.POST("/albums", handler.CreateAlbum)
	r.PUT("/albums/:id", handler.UpdateAlbum)
	r.DELETE("/albums/:id", handler.DeleteAlbum)
	return r
}

func TestHandlers(t *testing.T) {
	mockService := new(MockAlbumService)
	r := setupTestRouter(mockService)

	t.Run("GET :: /albums endpoint", func(t *testing.T) {
		albums := []domain.Album{{ID: 1, Title: "Test Album", Artist: "Test Artist", Price: 9.99}}
		mockService.On("GetAllAlbums").Return(albums, nil)

		req, _ := http.NewRequest("GET", "/albums", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var responseAlbums []domain.Album
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbums)
		assert.Nil(t, err)
		assert.Equal(t, albums, responseAlbums)
		mockService.AssertExpectations(t)
	})

	t.Run("GET :: /albums/:id endpoint", func(t *testing.T) {
		album := domain.Album{ID: 1, Title: "Test Album", Artist: "Test Artist", Price: 9.99}
		mockService.On("GetAlbumByID", 1).Return(album, nil)

		req, _ := http.NewRequest("GET", "/albums/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var responseAlbum domain.Album
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		assert.Nil(t, err)
		assert.Equal(t, album, responseAlbum)
		mockService.AssertExpectations(t)
	})

	t.Run("POST :: /albums endpoint", func(t *testing.T) {
		album := domain.Album{Title: "Test Album", Artist: "Test Artist", Price: 9.99}
		createdAlbum := domain.Album{ID: 1, Title: "Test Album", Artist: "Test Artist", Price: 9.99}
		mockService.On("CreateAlbum", album).Return(createdAlbum, nil)

		jsonValue, _ := json.Marshal(album)
		req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var responseAlbum domain.Album
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		assert.Nil(t, err)
		assert.Equal(t, createdAlbum, responseAlbum)
		mockService.AssertExpectations(t)
	})

	t.Run("PUT :: /albums/:id endpoint", func(t *testing.T) {
		album := domain.Album{ID: 1, Title: "Updated Album", Artist: "Updated Artist", Price: 19.99}
		mockService.On("UpdateAlbum", album).Return(album, nil)

		jsonValue, _ := json.Marshal(album)
		req, _ := http.NewRequest("PUT", "/albums/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var responseAlbum domain.Album
		err := json.Unmarshal(w.Body.Bytes(), &responseAlbum)
		assert.Nil(t, err)
		assert.Equal(t, album, responseAlbum)
		mockService.AssertExpectations(t)
	})

	t.Run("DELETE :: /albums/:id endpoint", func(t *testing.T) {
		mockService.On("DeleteAlbum", 1).Return(nil)

		req, _ := http.NewRequest("DELETE", "/albums/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
		mockService.AssertExpectations(t)
	})
}
