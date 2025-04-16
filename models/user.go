package models

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Temp DB
var Users = []User{
	{ID: 1, Email: "alice@example.com", Password: "123456"},
	{ID: 2, Email: "bob@example.com", Password: "password"},
}
