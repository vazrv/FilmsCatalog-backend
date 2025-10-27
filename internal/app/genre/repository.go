// internal/app/genre/repository.go
package genre

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetAll() ([]Genre, error) {
	var genres []Genre
	err := r.DB.Table("genres").Find(&genres).Error
	return genres, err
}