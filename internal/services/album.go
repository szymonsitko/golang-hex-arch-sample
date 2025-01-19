package services

import (
	"github.com/ssitko/hex-domain/internal/domain"
	"github.com/ssitko/hex-domain/internal/repositories"
)

// Service Layer
// Orchestrates the business logic and interacts with the repository.
type AlbumService struct {
	repo repositories.AlbumRepository
}

func NewAlbumService(repo repositories.AlbumRepository) *AlbumService {
	return &AlbumService{repo: repo}
}

func (s *AlbumService) GetAllAlbums() ([]domain.Album, error) {
	return s.repo.GetAll()
}

func (s *AlbumService) GetAlbumByID(id int) (domain.Album, error) {
	return s.repo.GetByID(id)
}

func (s *AlbumService) CreateAlbum(album domain.Album) (domain.Album, error) {
	return s.repo.Create(album)
}

func (s *AlbumService) UpdateAlbum(album domain.Album) (domain.Album, error) {
	return s.repo.Update(album)
}

func (s *AlbumService) DeleteAlbum(id int) error {
	return s.repo.Delete(id)
}
