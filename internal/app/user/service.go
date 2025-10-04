package user

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetProfile(userID uint) (*UserProfile, error) {
	return s.repo.GetProfileWithStats(userID)
}
