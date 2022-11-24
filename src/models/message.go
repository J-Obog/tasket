package models

type Message struct {
	Id        string
	Data      []byte
	Timestamp int64
}

type ScheduledTask struct {
	TaskId     string
	TaskConfig TaskConfig
}
