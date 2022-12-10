package utils

import (
	"time"

	"github.com/segmentio/ksuid"
)

type ITimeProvider interface {
	Now() int64
}

func GenerateUUID() string {
	return ksuid.New().String()
}

func TimeNow() int64 {
	return time.Now().Unix()
}

func GenerateHash(s string) []byte {
	return nil
}
