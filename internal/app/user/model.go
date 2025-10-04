package user

type ProfileStats struct {
	FavoritesCount int `json:"favorites_count"`
	ReviewsCount   int `json:"reviews_count"`
	RatingsCount   int `json:"ratings_count"`
}

type UserProfile struct {
	ID          uint         `json:"id"`
	Username    string       `json:"username"`
	Email       string       `json:"email"`
	AvatarURL   string       `json:"avatar_url,omitempty"`
	BackdropURL string       `json:"backdrop_url,omitempty"`
	CreatedAt   string       `json:"created_at"`
	IsAdmin     bool         `json:"is_admin"`
	Stats       ProfileStats `json:"stats"`
}
