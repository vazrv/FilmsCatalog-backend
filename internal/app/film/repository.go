// internal/app/film/repository.go
package film

// исключительно работа с базой

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetPopularFilms(limit int) ([]FilmFlat, error) {
	var films []FilmFlat

	query := `
		SELECT 
			f.id,
			f.title_ru,
			f.title_original,
			f.year,
			f.poster_url,
			f.kinopoisk_rating,
			json_agg(json_build_object('id', g.id, 'name', g.name)) FILTER (WHERE g.id IS NOT NULL) AS genres
		FROM films f
		LEFT JOIN film_genres fg ON f.id = fg.film_id
		LEFT JOIN genres g ON fg.genre_id = g.id
		WHERE f.kinopoisk_rating IS NOT NULL
		GROUP BY f.id
		ORDER BY f.kinopoisk_rating DESC
		LIMIT ?
	`

	err := r.DB.Raw(query, limit).Scan(&films).Error
	if err != nil {
		return nil, err
	}

	return films, nil
}
