package handlers

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ssitko/hex-domain/config"
	album "github.com/ssitko/hex-domain/internal/domain"
	"github.com/ssitko/hex-domain/internal/handlers"
	"github.com/ssitko/hex-domain/internal/infrastructure/persistence"
	"github.com/ssitko/hex-domain/internal/repositories"
	"github.com/ssitko/hex-domain/internal/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var albumHandler *handlers.AlbumHandler

func init() {
	testEnvFilePath := os.Getenv("TEST_CONFIG_FILE_PATH")
	if testEnvFilePath == "" {
		log.Fatalf("invalid test config file path (.env): %s", testEnvFilePath)
	}

	err := config.LoadConfig(testEnvFilePath)
	if err != nil {
		log.Fatalf("invalid config provided %s", err)
	}

	dsn := config.GetConfigValue(config.DB_USER) + ":" + config.GetConfigValue(config.DB_PASSWORD) + "@tcp(" + config.GetConfigValue(config.DB_HOST) + ":" + config.GetConfigValue(config.DB_PORT) + ")/" + config.GetConfigValue(config.DB_NAME) + "?charset=utf8mb4&parseTime=True&loc=Local"
	mysqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}
	mysqlDB.AutoMigrate(&album.Album{})
	db := persistence.NewGormDBWrapper(mysqlDB)

	repo := repositories.NewGormAlbumRepository(db)
	service := services.NewAlbumService(repo)
	albumHandler = handlers.NewAlbumHandler(service)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/albums", albumHandler.GetAlbums)
	r.GET("/albums/:id", albumHandler.GetAlbumByID)
	r.POST("/albums", albumHandler.CreateAlbum)
	r.PUT("/albums/:id", albumHandler.UpdateAlbum)
	r.DELETE("/albums/:id", albumHandler.DeleteAlbum)
	return r
}

func TestAlbumHandlers(t *testing.T) {
	t.Run("GET :: /albums endpoint", func(t *testing.T) {
		r := setupRouter()

		req, _ := http.NewRequest("GET", "/albums", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var albums []album.Album
		err := json.Unmarshal(w.Body.Bytes(), &albums)
		assert.Nil(t, err)
		assert.NotEmpty(t, albums)
	})

	t.Run("GET :: /albums/1 endpoint", func(t *testing.T) {
		r := setupRouter()

		req, _ := http.NewRequest("GET", "/albums/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var album album.Album
		err := json.Unmarshal(w.Body.Bytes(), &album)
		assert.Nil(t, err)
		assert.Equal(t, uint(1), album.ID)
	})

	t.Run("POST :: /albums endpoint", func(t *testing.T) {
		r := setupRouter()

		albumEntity := album.Album{Title: "Test Album", Artist: "Test Artist", Price: 9.99}
		jsonValue, _ := json.Marshal(albumEntity)
		req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		var createdAlbum album.Album
		err := json.Unmarshal(w.Body.Bytes(), &createdAlbum)
		assert.Nil(t, err)
		assert.Equal(t, albumEntity.Title, createdAlbum.Title)
	})

	t.Run("PUT :: /albums/1 endpoint", func(t *testing.T) {
		r := setupRouter()

		albumEntity := album.Album{ID: 1, Title: "Updated Album", Artist: "Updated Artist", Price: 19.99}
		jsonValue, _ := json.Marshal(albumEntity)
		req, _ := http.NewRequest("PUT", "/albums/1", bytes.NewBuffer(jsonValue))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var updatedAlbum album.Album
		err := json.Unmarshal(w.Body.Bytes(), &updatedAlbum)
		assert.Nil(t, err)
		assert.Equal(t, albumEntity.Title, updatedAlbum.Title)
	})

	t.Run("DELETE :: /albums/1 endpoint", func(t *testing.T) {
		r := setupRouter()

		req, _ := http.NewRequest("DELETE", "/albums/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNoContent, w.Code)
	})
}
