package config

type Config struct {
	AppEnv          string `mapstructure:"app_env"`
	DbHost          string `mapstructure:"db_host"`
	DbPort          string `mapstructure:"db_port"`
	DbName          string `mapstructure:"db_name"`
	DbUser          string `mapstructure:"db_user"`
	DbPassword      string `mapstructure:"db_password"`
	DbSSL           string `mapstructure:"db_ssl"`
	Port            string `mapstructure:"app_port"`
	RedisUrl        string `mapstructure:"redis_url"`
	DefaultLanguage string `mapstructure:"default_language"`
	LogLevel        string `mapstructure:"log_level"`
	ServerCode      string `mapstructure:"server_code"`
	ServerPort      string `mapstructure:"server_port"`
}

