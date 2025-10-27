// internal/app/actor/repository.go
package actor

import "gorm.io/gorm"

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetTopActors(limit int) ([]Actor, error) {
	var actors []Actor

	query := `
		SELECT 
			p.id,
			p.name_ru AS name,
			p.photo_url,
			COUNT(fp.id) AS film_count
		FROM persons p
		JOIN film_persons fp ON p.id = fp.person_id
		WHERE fp.role_type = 'actor'
		GROUP BY p.id
		ORDER BY film_count DESC
		LIMIT ?
	`

	err := r.DB.Raw(query, limit).Scan(&actors).Error
	return actors, err
}