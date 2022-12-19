package store

import "github.com/J-Obog/tasket/src/model"

type UserStore interface {
	Get(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Insert(user model.User) error
	Delete(id string) error
}
