package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/J-Obog/tasket/src/api"
	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
	"github.com/J-Obog/tasket/src/utils"
)

type UserResource struct {
	userManager managers.UserManager
}

func NewUserResource(userManager managers.UserManager) *UserResource {
	return &UserResource{
		userManager: userManager,
	}
}

func (this *UserResource) GetUser(req api.RestRequest) (api.RestResponse, int) {
	id := utils.CurrentUserId(req)

	user, err := this.userManager.GetUserById(id)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	if user == nil {
		res := api.RestResponse{
			Error: api.RestError{
				Type:   "resource not found",
				Detail: fmt.Sprintf("User could not be found"),
			},
		}

		return res, http.StatusNotFound
	}

	res := api.RestResponse{
		Data: &user,
	}

	return res, http.StatusOK
}

func (this *UserResource) DeleteUser(req api.RestRequest) (api.RestResponse, int) {
	id := utils.CurrentUserId(req)

	err := this.userManager.DeleteUser(id)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Message: "User deleted successfully",
	}

	return res, http.StatusOK
}

func (this *UserResource) CreateUser(req api.RestRequest) (api.RestResponse, int) {
	var newUser models.NewUser

	err := json.Unmarshal(req.Body, &newUser)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	//validate user request

	err = this.userManager.CreateUser(newUser)

	if err != nil {
		return api.ServerErrorResponse(), http.StatusInternalServerError
	}

	res := api.RestResponse{
		Message: "User created successfully",
	}

	return res, http.StatusOK
}
