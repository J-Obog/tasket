package config

type AppConfig struct {
	RedisAddress    string `json:"redisAddress"`
	RedisUsername   string `json:"redisUsername"`
	RedisPassword   string `json:"redisPassword"`
	RabbitmqAddress string `json:"rabbitmqAddress"`
	TaskQueueName   string `json:"taskQueueName"`
}
