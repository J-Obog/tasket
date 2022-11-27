package models

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  []byte `json:"password"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type UserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
