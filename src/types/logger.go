package types

type ILogger interface {
	Info(message string)
	Debug(message string)
	Warn(message string)
	Error(message string)
}
