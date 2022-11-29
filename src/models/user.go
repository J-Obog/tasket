package models

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
