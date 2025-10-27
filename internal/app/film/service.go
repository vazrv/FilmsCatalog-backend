// internal/app/film/service.go
package film

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetPopularFilms(limit int) ([]FilmFlat, error) {
	return s.repo.GetPopularFilms(limit)
}

// получается из данные репозитория и обрабатывает их
