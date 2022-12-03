package models

type Message struct {
	Id        string
	Data      []byte
	Timestamp int64
}

type TaskScheduledEvent struct {
	TaskId     string
	TaskConfig TaskConfig
}
