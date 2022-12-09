package queue

type IQueue interface {
	Push(serializable interface{}) error
	Pull() ([]byte, error)
}
