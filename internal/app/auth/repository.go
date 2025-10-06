// internal/app/auth/repository.go
package auth

import "gorm.io/gorm"

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
