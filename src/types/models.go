package types

type RestRequest struct {
	Url         string
	UrlParams   map[string]interface{}
	QueryParams map[string]interface{}
	Metadata    map[string]interface{}
	Body        []byte
}

type RestError struct {
	Type   string `json:"type"`
	Detail string `json:"detail"`
}

type RestMessage struct {
	Message string `json:"message"`
}

type RestResponse struct {
	Data   interface{}
	Status int64
}

type EventMessage struct {
	Type      EventType   `json:"type"`
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

type TaskScheduledMessage struct {
	TaskId     string     `json:"taskId"`
	TaskConfig TaskConfig `json:"taskConfig"`
}

type TaskStoppedMessage struct {
	TaskId string `json:"taskId"`
}

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

type TaskConfig struct {
	MaxTimeout int64             `json:"maxTimeout"`
	MaxMemory  int64             `json:"maxMemory"`
	Commands   []string          `json:"commands"`
	Environ    map[string]string `json:"environ"`
	SourceMeta map[string]string `json:"sourceMeta"`
}

type Task struct {
	Id          string     `json:"id"`
	UserId      string     `json:"userId"`
	Name        string     `json:"name"`
	Status      TaskStatus `json:"status"`
	Config      TaskConfig `json:"config"`
	StartedAt   *int64     `json:"startedAt"`
	CompletedAt *int64     `json:"completedAt"`
	CreatedAt   int64      `json:"createdAt"`
	UpdatedAt   int64      `json:"updatedAt"`
}

type NewTask struct {
	Name   string     `json:"name"`
	Config TaskConfig `json:"config"`
}

type UpdatedTask struct {
	Name string `json:"name"`
}

type TaskOptions struct {
	Name          *string     `json:"name"`
	Status        *TaskStatus `json:"status"`
	CreatedAfter  *int64      `json:"createdAfter"`
	CreatedBefore *int64      `json:"createdBefore"`
}

type User struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  []byte `json:"-"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

type NewUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
