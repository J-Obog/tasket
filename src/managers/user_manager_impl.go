package managers

import (
	"github.com/J-Obog/tasket/src/model"
	"github.com/J-Obog/tasket/src/store"
	"github.com/J-Obog/tasket/src/types"
)

type UserManagerImpl struct {
	userStore    store.UserStore
	timeProvider types.ITimeProvider
	uuidProvider types.IUUIDProvider
}

func NewUserManager(userStore store.UserStore, timeProvider types.ITimeProvider, uuidProvider types.IUUIDProvider) *UserManagerImpl {
	return &UserManagerImpl{
		userStore:    userStore,
		timeProvider: timeProvider,
		uuidProvider: uuidProvider,
	}
}

func (this *UserManagerImpl) Get(id string) (*model.User, error) {
	return this.userStore.Get(id)
}

func (this *UserManagerImpl) GetByEmail(email string) (*model.User, error) {
	return this.userStore.GetByEmail(email)
}

func (this *UserManagerImpl) Create(email string, password string) error {
	now := this.timeProvider.Now()

	user := model.User{
		Id:        this.uuidProvider.UUID(),
		Email:     email,
		Password:  password,
		CreatedAt: now,
	}

	return this.userStore.Insert(user)
}

func (this *UserManagerImpl) Delete(id string) error {
	return nil
}
