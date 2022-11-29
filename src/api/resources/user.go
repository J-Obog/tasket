package resources

import (
	"encoding/json"
	"net/http"

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

func (this *UserResource) GetUser(req models.RestRequest) models.RestResponse {
	id := utils.CurrentUserId(req)

	user, err := this.userManager.GetUserById(id)

	if err != nil {
		return utils.MakeServerError()
	}
	return models.RestResponse{
		Object: map[string]interface{}{
			"data": &user,
		},
		Status: http.StatusOK,
	}
}

func (this *UserResource) DeleteUser(req models.RestRequest) models.RestResponse {
	id := utils.CurrentUserId(req)

	err := this.userManager.DeleteUser(id)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Object: map[string]interface{}{
			"message": "User deleted successfully",
		},
		Status: http.StatusOK,
	}
}

func (this *UserResource) CreateUser(req models.RestRequest) models.RestResponse {
	var newUser models.NewUser

	err := json.Unmarshal(req.Body, &newUser)

	if err != nil {
		return utils.MakeServerError()
	}

	//validate user request

	err = this.userManager.CreateUser(newUser)

	if err != nil {
		return utils.MakeServerError()
	}

	return models.RestResponse{
		Object: map[string]interface{}{
			"message": "User created successfully",
		},
		Status: http.StatusOK,
	}
}
