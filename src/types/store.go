package types

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

type IUserStore interface {
	GetById(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	Insert(user User) error
	Delete(id string) error
}
