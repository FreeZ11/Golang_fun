package types

import "time"

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByLogin(email string) (*GetUserPayload, error)
	CreateUser(user User) error
}

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	Login     string `json:"login" validate:"required"`
}

type GetUserPayload struct {
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	Login     string    `json:"login" validate:"required"`
	CreatedAt time.Time `json:"createdAt"`
}

type User struct {
	Id        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Login     string    `json:"login"`
	CreatedAt time.Time `json:"createdAt"`
}
