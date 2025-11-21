package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// AppConfig 汇总所有配置段供全局使用。
type AppConfig struct {
	App       AppSection       `mapstructure:"app"`
	Database  DatabaseSection  `mapstructure:"database"`
	Redis     RedisSection     `mapstructure:"redis"`
	TaskQueue TaskQueueSection `mapstructure:"task_queue"`
	Vector    VectorSection    `mapstructure:"vector_store"`
	Logging   LoggingSection   `mapstructure:"logging"`
	Security  SecuritySection  `mapstructure:"security"`
}

// AppSection 基础应用配置。
type AppSection struct {
	Name    string      `mapstructure:"name"`
	Env     string      `mapstructure:"env"`
	Version string      `mapstructure:"version"`
	HTTP    HTTPSection `mapstructure:"http"`
}

// HTTPSection HTTP服务配置。
type HTTPSection struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

// DatabaseSection 数据库连接配置。
type DatabaseSection struct {
	Driver string `mapstructure:"driver"`
	DSN    string `mapstructure:"dsn"`
}

// RedisSection Redis配置。
type RedisSection struct {
	Addr string `mapstructure:"addr"`
	DB   int    `mapstructure:"db"`
}

// TaskQueueSection 异步任务配置。
type TaskQueueSection struct {
	RedisDSN    string `mapstructure:"redis_dsn"`
	Concurrency int    `mapstructure:"concurrency"`
}

// VectorSection 向量数据库配置。
type VectorSection struct {
	Provider string `mapstructure:"provider"`
	Endpoint string `mapstructure:"endpoint"`
}

// LoggingSection 日志配置。
type LoggingSection struct {
	Level    string `mapstructure:"level"`
	Encoding string `mapstructure:"encoding"`
	Filename string `mapstructure:"filename"`
}

// SecuritySection 安全配置。
type SecuritySection struct {
	JWTSecret string `mapstructure:"jwt_secret"`
}

// Load 通过viper加载配置。
func Load(path string) (*AppConfig, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	v.SetEnvPrefix("ROBOT_MANAGE")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("读取配置失败: %w", err)
	}

	var cfg AppConfig
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("解析配置失败: %w", err)
	}

	return &cfg, nil
}
