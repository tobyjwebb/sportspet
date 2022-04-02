package settings

import "os"

type Config struct {
	FrontendAddr string
	RedisAddr    string
}

func GetConfig() *Config {
	frontendAddr := os.Getenv("TC_FRONTEND_ADDR")
	if frontendAddr == "" {
		frontendAddr = ":8081"
	}
	redisAddr := os.Getenv("TC_REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = ":6379"
	}
	return &Config{
		FrontendAddr: frontendAddr,
		RedisAddr:    redisAddr,
	}
}
