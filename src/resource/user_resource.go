package resource

type UserResource interface {
	GetUser(req RestRequest) RestResponse
	DeleteUser(req RestRequest) RestResponse
	CreateUser(req RestRequest) RestResponse
}
