// internal/app/user/repository.go
package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetProfileWithStats(userID uint) (*UserProfileFlat, error) {
	var profile UserProfileFlat

	query := `
        SELECT 
            u.id,
            u.username,
            u.email,
            u.avatar_url,
            u.created_at,
            u.is_admin,
            COALESCE(fav.count, 0) AS favorites_count,
            COALESCE(rev.count, 0) AS reviews_count,
            COALESCE(rat.count, 0) AS ratings_count
        FROM users u
        LEFT JOIN (SELECT user_id, COUNT(*) AS count FROM favorites GROUP BY user_id) fav ON u.id = fav.user_id
        LEFT JOIN (SELECT user_id, COUNT(*) AS count FROM reviews GROUP BY user_id) rev ON u.id = rev.user_id
        LEFT JOIN (SELECT user_id, COUNT(*) AS count FROM ratings GROUP BY user_id) rat ON u.id = rat.user_id
        WHERE u.id = ?
    `

	err := r.DB.Raw(query, userID).Scan(&profile).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}