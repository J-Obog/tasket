package model

type LogSource uint

const (
	LogSource_UNKNOWN LogSource = 0
	LogSource_TASK    LogSource = 1
	LogSource_SERVER  LogSource = 2
)

type Log struct {
	Id        string    `json:"id"`
	TaskId    string    `json:"taskId"`
	Source    LogSource `json:"source"`
	Content   []byte    `json:"content"`
	CreatedAt int64     `json:"createdAt"`
}

type LogOptions struct {
	Source        *LogSource `json:"source"`
	CreatedAfter  *int64     `json:"createdAfter"`
	CreatedBefore *int64     `json:"createdBefore"`
}
