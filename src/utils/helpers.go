package utils

import (
	"time"

	"github.com/segmentio/ksuid"
)

func GenerateUUID() string {
	return ksuid.New().String()
}

func TimeNow() int64 {
	return time.Now().Unix()
}

func GenerateHash(s string) []byte {
	return nil
}
