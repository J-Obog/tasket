package utils

import (
	"net/http"

	"github.com/J-Obog/tasket/src/models"
)

func GenerateUUID() string {
	return ""
}

func TimeNow() int64 {
	return 0
}

func GenerateHash(s string) []byte {
	return nil
}

func CurrentUserId(req models.RestRequest) string {
	return ""
}

func MakeServerError() models.RestResponse {
	return models.RestResponse{
		Body:   make([]byte, 0),
		Status: http.StatusInternalServerError,
	}
}

func TaskIdParam() string {
	return ""
}
