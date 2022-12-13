package types

type ICache interface {
	Get(key string) (*string, error)
	Set(key string, value string, expr int64) error
}

type ILogger interface {
	Info(message string)
	Debug(message string)
	Warn(message string)
	Error(message string)
}

type IQueue interface {
	Push(message EventMessage) error
	Pull() (*EventMessage, error)
}

type IServer interface {
	Start(host string, port int) error
	Stop() error
}

type ITimeProvider interface {
	Now() int64
}

type IHashProvider interface {
}

type IUUIDProvider interface {
	UUID() string
}

type ILogManager interface {
	GetLogsByTask(taskId string) ([]Log, error)
	GetLogsByFilter(taskId string, filter LogOptions) ([]Log, error)
	CreateLog(taskId string, newLog NewLog)
}

type ITaskManager interface {
	GetTaskById(id string) (*Task, error)
	GetTasksByFilter(userId string, options TaskOptions) ([]Task, error)
	UpdateTask(id string, updatedTask UpdatedTask) error
	UpdateTaskStatus(id string, status TaskStatus) error
	StopTask(id string) error
	CreateTask(userId string, newTask NewTask) error
}

type IUserManager interface {
	GetUserById(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	CreateUser(userReq NewUser) error
	DeleteUser(id string) error
}

type IUserResource interface {
	GetUser(req RestRequest) RestResponse
	DeleteUser(req RestRequest) RestResponse
	CreateUser(req RestRequest) RestResponse
}

type ITaskResource interface {
	GetTask(req RestRequest) RestResponse
	GetTasks(req RestRequest) RestResponse
	GetTaskLogs(req RestRequest) RestResponse
	StopTask(req RestRequest) RestResponse
	UpdateTask(req RestRequest) RestResponse
	CreateTask(req RestRequest) RestResponse
}

type IUserStore interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Insert(user User) error
	Delete(id string) error
}

type ITaskStore interface {
	GetById(id string) (*Task, error)
	GetByFilter(userId string, options TaskOptions) ([]Task, error)
	Update(id string, updatedTask UpdatedTask, updatedAt int64) error
	UpdateStatus(id string, status TaskStatus, updatedAt int64) error
	Insert(task Task) error
}

type ILogStore interface {
	GetByTask(taskId string) ([]Log, error)
	GetByFilter(taskId string, options LogOptions) ([]Log, error)
	Insert(log Log) error
}
