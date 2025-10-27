// internal/app/actor/model.go
package actor

type Actor struct {
	ID    uint   `json:"id" gorm:"column:id"`
	Name  string `json:"name" gorm:"column:name"`
	Photo string `json:"photo_url" gorm:"column:photo_url"`
}
