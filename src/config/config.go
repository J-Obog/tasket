package config

const (
	taskQueueName string = "tasks"
)

type RedisConfig struct {
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type RabbitMqConfig struct {
	Address       string `json:"address"`
	TaskQueueName string
}

type AppConfig struct {
	Redis    RedisConfig    `json:"redis"`
	RabbitMq RabbitMqConfig `json:"rabbitMq"`
}
