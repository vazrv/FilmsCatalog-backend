// internal/app/auth/model.go
package auth

type User struct {
	ID         uint   `json:"id"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	PasswordHash string `json:"-"` // не отправляется в JSON
	AvatarURL  string `json:"avatar_url,omitempty"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at,omitempty"`
	IsAdmin    bool   `json:"is_admin"`
	// Удалены: FavoritesCount, ReviewsCount, RatingsCount — их нет в БД
}

type RegisterRequest struct {
	Username        string `json:"username" binding:"required,min=3"`
	Email           string `json:"email" binding:"required,email"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirm_password" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}