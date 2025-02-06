package models

type User struct {
	ID    uint   `json:"idd"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
