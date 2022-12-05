package mocks

import (
	"context"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/stretchr/testify/mock"
)

type MockRedisClient struct {
	mock.Mock
	redis.Client
}

func (this *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	ret := this.Called(ctx, key)

	var r0 *redis.StringCmd
	if rf, ok := ret.Get(0).(func() *redis.StringCmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StringCmd)
		}
	}

	return r0
}

func (this *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	ret := this.Called(ctx, key, value, expiration)

	var r0 *redis.StatusCmd
	if rf, ok := ret.Get(0).(func() *redis.StatusCmd); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*redis.StatusCmd)
		}
	}

	return r0
}
