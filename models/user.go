package models

import "time"

//User dados do usuario
type User struct {
	ID        string    `json:"_id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_At"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

//Users lista de usuarios
type Users []*User
