package resource

import (
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/rest"
)

type UserResourceImpl struct {
	userManager managers.UserManager
}

func NewUserResource(userManager managers.UserManager) *UserResourceImpl {
	return &UserResourceImpl{
		userManager: userManager,
	}
}

func (this *UserResourceImpl) Get(req rest.RestRequest) rest.RestResponse {
	id := UserId(req.Metadata)

	user, err := this.userManager.Get(id)

	if err != nil {
		return ServerErrorResponse()
	}

	if user == nil {
		resErr := MakeError("resource not found", fmt.Sprintf("User could not be found"))
		return MakeResponse(&resErr, http.StatusNotFound)
	}

	return MakeResponse(&user, http.StatusOK)
}

func (this *UserResourceImpl) Delete(req rest.RestRequest) rest.RestResponse {
	id := UserId(req.Metadata)

	err := this.userManager.Delete(id)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage("User deleted successfully")
	return MakeResponse(&msg, http.StatusOK)
}

func (this *UserResourceImpl) Create(req rest.RestRequest) rest.RestResponse {
	/*var newUser types.NewUser

	err := json.Unmarshal(req.Body, &newUser)

	if err != nil {
		return ServerErrorResponse()
	}

	//validate user request

	err = this.userManager.CreateUser(newUser)

	if err != nil {
		return ServerErrorResponse()
	}

	msg := MakeMessage("User created successfully")*/
	return MakeResponse(nil, http.StatusOK)
}
