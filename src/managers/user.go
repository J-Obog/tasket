package managers

import (
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/utils"
)

type UserManager struct {
	userStore store.UserStore
}

func NewUserManager(userStore store.UserStore) *UserManager {
	return &UserManager{
		userStore: userStore,
	}
}

func (this *UserManager) GetUserById(id string) (*models.User, error) {
	return this.userStore.GetById(id)
}

func (this *UserManager) GetUserByEmail(email string) (*models.User, error) {
	return this.userStore.GetByEmail(email)
}

func (this *UserManager) CreateUser(userReq models.NewUser) error {
	now := utils.TimeNow()
	id := utils.GenerateUUID()
	hashedPswd := utils.GenerateHash(userReq.Password)

	user := models.User{
		Id:        id,
		Email:     userReq.Email,
		Password:  hashedPswd,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return this.userStore.Insert(user)
}

func (this *UserManager) DeleteUser(id string) error {
	return nil
}
