package queue

type Queue interface {
	Push(serializable interface{}) error
	Pull() ([]byte, error)
}
