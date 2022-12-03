package utils

import (
	"time"

	"github.com/J-Obog/tasket/src/api"
)

func GenerateUUID() string {
	return ""
}

func TimeNow() int64 {
	return time.Now().Unix()
}

func GenerateHash(s string) []byte {
	return nil
}

func CurrentUserId(req api.RestRequest) string {
	return ""
}
