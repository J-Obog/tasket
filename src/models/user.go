package models

type User struct {
	Id        string
	Email     string
	Password  []byte
	CreatedAt int64
	UpdatedAt int64
}

type UserRequest struct {
	Email    string
	Password string
}
