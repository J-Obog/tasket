package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/rest"
)

type UserResource struct {
	userManager managers.UserManager
}

func NewUserResource(userManager managers.UserManager) *UserResource {
	return &UserResource{
		userManager: userManager,
	}
}

func (this *UserResource) GetUser(req rest.RestRequest) (rest.RestResponse, int) {
	id := rest.UserId(req.Metadata)

	user, err := this.userManager.GetUserById(id)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	if user == nil {
		res := rest.RestResponse{
			Error: rest.RestError{
				Type:   "resource not found",
				Detail: fmt.Sprintf("User could not be found"),
			},
		}

		return res, http.StatusNotFound
	}

	res := rest.RestResponse{
		Data: &user,
	}

	return res, http.StatusOK
}

func (this *UserResource) DeleteUser(req rest.RestRequest) (rest.RestResponse, int) {
	id := rest.UserId(req.Metadata)

	err := this.userManager.DeleteUser(id)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Message: "User deleted successfully",
	}

	return res, http.StatusOK
}

func (this *UserResource) CreateUser(req rest.RestRequest) (rest.RestResponse, int) {
	var newUser models.NewUser

	err := json.Unmarshal(req.Body, &newUser)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate user request

	err = this.userManager.CreateUser(newUser)

	if err != nil {
		return rest.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := rest.RestResponse{
		Message: "User created successfully",
	}

	return res, http.StatusOK
}
