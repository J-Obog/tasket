package resources

import (
	"github.com/J-Obog/tasket/src/managers"
	"github.com/J-Obog/tasket/src/models"
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
	return models.RestResponse{}
}

func (this *UserResource) DeleteUser(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}

func (this *UserResource) CreateUser(req models.RestRequest) models.RestResponse {
	return models.RestResponse{}
}
