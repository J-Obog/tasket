package queue

type Queue interface {
	Push(message EventMessage) error
	Pull() (*EventMessage, error)
}
