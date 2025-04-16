package models

type Employee struct {
	ID       int    `json:"id"` // field json map
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
}
