package models

type LogSource uint

const (
	LogSource_UNKNOWN LogSource = 0
	LogSource_TASK              = 1
	LogSource_SERVER            = 2
)

type Log struct {
	Id        string    `json:"id"`
	TaskId    string    `json:"taskId"`
	Source    LogSource `json:"source"`
	Content   []byte    `json:"content"`
	CreatedAt int64     `json:"createdAt"`
}

type NewLog struct {
	TaskId  string
	Source  LogSource
	Content []byte
}

type LogOptions struct {
	Source        *LogSource `json:"source"`
	Contains      []byte     `json:"contains"`
	CreatedAfter  *int64     `json:"createdAfter"`
	CreatedBefore *int64     `json:"createdBefore"`
	First         *int64     `json:"first"`
	Last          *int64     `json:"last"`
}
