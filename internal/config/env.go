package config

type AppConfig struct {
	AppEnv   string // "development", "production", v.v.
	DBUrl    string
	Port     string
	RedisUrl string
	DefaultLanguage string // Mặc định là "vi"
	LogLevel string // "debug", "info", "warn", "error"
}

var Config AppConfig
