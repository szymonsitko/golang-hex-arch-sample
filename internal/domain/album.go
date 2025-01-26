package domain

// Domain Layer
// Represents the core business logic and entities.
type Album struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Album service interface definition.
type AlbumService interface {
	CreateAlbum(album Album) (Album, error)
	DeleteAlbum(id int) error
	GetAlbumByID(id int) (Album, error)
	GetAllAlbums() ([]Album, error)
	UpdateAlbum(album Album) (Album, error)
}
