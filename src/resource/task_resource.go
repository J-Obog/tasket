package resource

import (
	"github.com/J-Obog/tasket/src/rest"
)

type TaskResource interface {
	Get(req rest.RestRequest) rest.RestResponse
	GetBy(req rest.RestRequest) rest.RestResponse
	GetLogs(req rest.RestRequest) rest.RestResponse
	Stop(req rest.RestRequest) rest.RestResponse
	Update(req rest.RestRequest) rest.RestResponse
	Create(req rest.RestRequest) rest.RestResponse
}
