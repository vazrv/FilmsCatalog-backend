// internal/app/user/service.go
package user

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userID uint) (*UserProfileFlat, error) {
	return s.repo.GetProfileWithStats(userID)
}
