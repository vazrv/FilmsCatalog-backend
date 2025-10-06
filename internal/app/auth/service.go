// internal/app/auth/service.go
package auth

type AuthService struct {
	repo *UserRepository
}

func NewAuthService(repo *UserRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *AuthService) GetUserByEmail(email string) (*User, error) {
	return s.repo.GetUserByEmail(email)
}
