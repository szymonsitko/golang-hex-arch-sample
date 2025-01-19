package repositories

import (
	album "github.com/ssitko/hex-domain/internal/domain"
	"github.com/ssitko/hex-domain/internal/infrastructure/persistence"
)

// TODO: add comments
type AlbumRepository interface {
	GetAll() ([]album.Album, error)
	GetByID(id int) (album.Album, error)
	Create(album album.Album) (album.Album, error)
	Update(album album.Album) (album.Album, error)
	Delete(id int) error
}

type GormAlbumRepository struct {
	db persistence.DB
}

func NewGormAlbumRepository(db persistence.DB) *GormAlbumRepository {
	return &GormAlbumRepository{db: db}
}

func (r *GormAlbumRepository) GetAll() ([]album.Album, error) {
	var albums []album.Album
	if err := r.db.Find(&albums); err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *GormAlbumRepository) GetByID(id int) (album.Album, error) {
	var album album.Album
	if err := r.db.First(&album, id); err != nil {
		return album, err
	}
	return album, nil
}

func (r *GormAlbumRepository) Create(albumEntity album.Album) (album.Album, error) {
	if err := r.db.Create(&albumEntity); err != nil {
		return album.Album{}, err
	}
	return albumEntity, nil
}

func (r *GormAlbumRepository) Update(albumEntity album.Album) (album.Album, error) {
	if err := r.db.Save(&albumEntity); err != nil {
		return album.Album{}, err
	}
	return albumEntity, nil
}

func (r *GormAlbumRepository) Delete(id int) error {
	if err := r.db.Delete(&album.Album{}, id); err != nil {
		return err
	}
	return nil
}
