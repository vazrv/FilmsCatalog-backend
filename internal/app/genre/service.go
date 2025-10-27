// internal/app/genre/service.go
package genre

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetAll() ([]Genre, error) {
	return s.repo.GetAll()
}