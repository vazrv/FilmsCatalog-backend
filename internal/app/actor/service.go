// internal/app/actor/service.go
package actor

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetTopActors(limit int) ([]Actor, error) {
	return s.repo.GetTopActors(limit)
}
