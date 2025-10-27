// internal/app/film/model.go
package film

// понять какие данные пойдут на фронт

type Genre struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// FilmFlat — плоская модель для RAW-запроса (без GORM-связей)
type FilmFlat struct {
	ID              uint    `json:"id"`
	TitleRu         string  `json:"title_ru"`
	TitleOriginal   string  `json:"title_original,omitempty"`
	Year            int     `json:"year"`
	PosterURL       string  `json:"poster_url"`
	KinopoiskRating float64 `json:"kinopoisk_rating"`

	// Поле для результата json_agg — будет заполнено напрямую
	GenresJSON []byte `json:"genres"` // json_agg вернёт JSON-массив, сохраняем как []byte
}
