package models

type Log struct {
	Id        string
	TaskId    string
	Source    LogRequest
	Content   []byte
	CreatedAt int64
}

type LogRequest struct {
	TaskId    string
	Source    string
	Content   []byte
	CreatedAt int64
}

type LogFilter struct {
	Source        string
	Contains      []byte
	CreatedAfter  int64
	CreatedBefore int64
	First         int64
	Last          int64
}
