package types

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
