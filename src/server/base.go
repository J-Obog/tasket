package server

type IServer interface {
	Start(host string, port int) error
	Stop() error
}
