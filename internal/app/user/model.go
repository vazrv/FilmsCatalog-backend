// internal/app/user/model.go
package user

// UserProfileFlat — плоская модель для маппинга из SQL
type UserProfileFlat struct {
	ID               uint   `json:"id"`
	Username         string `json:"username"`
	Email            string `json:"email"`
	AvatarURL        string `json:"avatar_url,omitempty"`
	CreatedAt        string `json:"created_at"`
	IsAdmin          bool   `json:"is_admin"`
	FavoritesCount   int    `json:"favorites_count"` // ⚠️ Прямое соответствие JSON
	ReviewsCount     int    `json:"reviews_count"`
	RatingsCount     int    `json:"ratings_count"`
}