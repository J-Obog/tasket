package types

type ITimeProvider interface {
	Now() int64
}
