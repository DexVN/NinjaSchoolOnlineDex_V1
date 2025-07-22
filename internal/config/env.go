package config

type AppConfig struct {
	AppEnv   string // "development", "production", v.v.
	DBUrl    string
	Port     string
	RedisUrl string
}

var Config AppConfig
