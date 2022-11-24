package models

type LogSource uint

const (
	LogSource_UNKNOWN LogSource = 0
	LogSource_TASK              = 1
	LogSource_SERVER            = 2
)

type Log struct {
	Id        string
	TaskId    string
	Source    LogSource
	Content   []byte
	CreatedAt int64
}

type LogRequest struct {
	TaskId    string
	Source    LogSource
	Content   []byte
	CreatedAt int64
}

type LogFilter struct {
	Source        LogSource
	Contains      []byte
	CreatedAfter  int64
	CreatedBefore int64
	First         int64
	Last          int64
}
