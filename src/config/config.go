package config

type AppConfig struct {
	RedisUrl      string `json:"redisUrl"`
	RedisUsername string `json:"redisUsername"`
	RedisPassword string `json:"redisPassword"`
}
