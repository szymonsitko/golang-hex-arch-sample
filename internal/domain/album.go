package domain

// Domain Layer
// Represents the core business logic and entities.
type Album struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
