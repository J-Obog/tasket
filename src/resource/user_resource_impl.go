package resource

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/types"
)

type UserResourceImpl struct {
	userManager types.IUserManager
}

func NewUserResource(userManager types.IUserManager) *UserResource {
	return &UserResource{
		userManager: userManager,
	}
}

func (this *UserResource) GetUser(req types.RestRequest) types.RestResponse {
	id := UserId(req.Metadata)

	user, err := this.userManager.GetUserById(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if user == nil {
		resErr := MakeError("resource not found", fmt.Sprintf("User could not be found"))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	return MakeResponse(&user, http.StatusOK)
}

func (this *UserResource) DeleteUser(req types.RestRequest) types.RestResponse {
	id := UserId(req.Metadata)

	err := this.userManager.DeleteUser(id)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage("User deleted successfully")
	return MakeResponse(&msg, http.StatusOK)
}

func (this *UserResource) CreateUser(req types.RestRequest) types.RestResponse {
	var newUser types.NewUser

	err := json.Unmarshal(req.Body, &newUser)

	if err != nil {
		return ServerErrorResponse()
	}

	//validate user request

	err = this.userManager.CreateUser(newUser)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage("User created successfully")
	return MakeResponse(&msg, http.StatusOK)
}
