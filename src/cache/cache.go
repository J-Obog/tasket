package cache

type Cache interface {
	Get(key string) (*string, error)
	Set(key string, value string, expr int64) error
}
