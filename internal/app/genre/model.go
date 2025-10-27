// internal/app/genre/model.go
package genre

type Genre struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}