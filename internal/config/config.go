package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server       ServerConfig
	Database     DatabaseConfig
	Logging      LoggingConfig
	Auth         AuthConfig
	FeatureFlags FeatureFlagsConfig
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host          string
	Port          int
	User          string
	Password      string
	Name          string
	SslMode       string
	MaxOpenConns  int
	MaxIdleConns  int
	ConnectionURL string // 生成的连接字符串
}

type LoggingConfig struct {
	Level string
	File  string
}

type AuthConfig struct {
	JWTSecret       string
	TokenExpiration time.Duration
}

type FeatureFlagsConfig struct {
	EnableFeatureX bool
	EnableFeatureY bool
}

func LoadConfig() (Config, error) {
	var config Config

	// 设置配置文件名和路径
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../../internal/config") // 可选的其他路径

	// 读取环境变量
	viper.AutomaticEnv()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return config, fmt.Errorf("failed to read config file: %w", err)
	}

	// 解析配置到结构体
	if err := viper.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	// 生成数据库连接字符串
	config.Database.ConnectionURL = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.SslMode,
	)

	// 解析时间持续时间
	if duration, err := time.ParseDuration(viper.GetString("auth.token_expiration")); err == nil {
		config.Auth.TokenExpiration = duration
	} else {
		log.Fatalf("Invalid token expiration: %s", viper.GetString("auth.token_expiration"))
	}

	return config, nil
}
