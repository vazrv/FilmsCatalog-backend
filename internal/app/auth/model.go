// internal/app/auth/model.go
package auth

type User struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	PasswordHash   string `json:"-"`
	AvatarURL      string `json:"avatar_url,omitempty"`
	BackdropURL    string `json:"backdrop_url,omitempty"`
	CreatedAt      string `json:"created_at"`
	IsAdmin        bool   `json:"is_admin"`
	FavoritesCount int    `json:"favorites_count"`
	ReviewsCount   int    `json:"reviews_count"`
	RatingsCount   int    `json:"ratings_count"`
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
