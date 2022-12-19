package resource

import (
	"github.com/J-Obog/tasket/src/rest"
)

type UserResource interface {
	Get(req rest.RestRequest) rest.RestResponse
	Delete(req rest.RestRequest) rest.RestResponse
	Create(req rest.RestRequest) rest.RestResponse
}
