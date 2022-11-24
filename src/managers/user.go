package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
)

type UserManager struct {
	userStore store.UserStore
}

func NewUserManager(userStore store.UserStore) *UserManager {
	return &UserManager{
		userStore: userStore,
	}
}

func (this *UserManager) GetUserById(id string) *models.User {
	return nil
}

func (this *UserManager) CreateUser(userReq models.UserRequest) {

}

func (this *UserManager) DeleteUser(id string) {

}
