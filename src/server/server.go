package server

type Server interface {
	Start(host string, port int) error
	Stop() error
}
