package managers

import "github.com/J-Obog/tasket/src/model"

type UserManager interface {
	Get(id string) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	Create(email string, password string) error
	Delete(id string) error
}
