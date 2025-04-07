package conf

import (
	"marine-backend/cmd/flags"
	"path/filepath"
)

type Database struct {
	Host        string `json:"host" env:"HOST"`
	Port        int    `json:"port" env:"PORT"`
	User        string `json:"user" env:"USER"`
	Password    string `json:"password" env:"PASS"`
	Name        string `json:"name" env:"NAME"`
	DBFile      string `json:"db_file" env:"FILE"`
	TablePrefix string `json:"table_prefix" env:"TABLE_PREFIX"`
	SSLMode     string `json:"ssl_mode" env:"SSL_MODE"`
	DSN         string `json:"dsn" env:"DSN"`
}

type LogConfig struct {
	Enable     bool   `json:"enable" env:"LOG_ENABLE"`
	Name       string `json:"name" env:"LOG_NAME"`
	MaxSize    int    `json:"max_size" env:"MAX_SIZE"`
	MaxBackups int    `json:"max_backups" env:"MAX_BACKUPS"`
	MaxAge     int    `json:"max_age" env:"MAX_AGE"`
	Compress   bool   `json:"compress" env:"COMPRESS"`
}

type Config struct {
	Database       Database  `json:"database" envPrefix:"DB_"`
	Address        string    `json:"address" env:"ADDR"`
	HttpPort       int       `json:"http_port" env:"HTTP_PORT"`
	Log            LogConfig `json:"log"`
	DelayedStart   int       `json:"delayed_start" env:"DELAYED_START"`
	MaxConnections int       `json:"max_connections" env:"MAX_CONNECTIONS"`
	MaxConcurrency int       `json:"max_concurrency" env:"MAX_CONCURRENCY"`
}

func DefaultConfig() *Config {
	logPath := filepath.Join(flags.DataDir, "log/log.log")
	return &Config{
		Address:  "0.0.0.0",
		HttpPort: 8080,
		Database: Database{
			DSN: "root:123456@tcp(127.0.0.1:3306)/port?charset=utf8mb4&parseTime=True&loc=Local",
		},
		Log: LogConfig{
			Enable:     true,
			Name:       logPath,
			MaxSize:    50,
			MaxBackups: 30,
			MaxAge:     28,
		},
		MaxConnections: 0,
		MaxConcurrency: 64,
	}
}
