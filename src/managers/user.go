package managers

import "github.com/J-Obog/tasket/src/types"

type UserManager struct {
	userStore    types.IUserStore
	timeProvider types.ITimeProvider
	uuidProvider types.IUUIDProvider
}

func NewUserManager(userStore types.IUserStore, timeProvider types.ITimeProvider, uuidProvider types.IUUIDProvider) *UserManager {
	return &UserManager{
		userStore:    userStore,
		timeProvider: timeProvider,
		uuidProvider: uuidProvider,
	}
}

func (this *UserManager) GetUserById(id string) (*types.User, error) {
	return this.userStore.GetById(id)
}

func (this *UserManager) GetUserByEmail(email string) (*types.User, error) {
	return this.userStore.GetByEmail(email)
}

//todo implement hashing
func (this *UserManager) CreateUser(userReq types.NewUser) error {
	now := this.timeProvider.Now()
	id := this.uuidProvider.UUID()
	hashedPswd := []byte{} //utils.GenerateHash(userReq.Password)

	user := types.User{
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
