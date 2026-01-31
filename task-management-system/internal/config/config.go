package config

import (
    "fmt"
    "os"
    
    "github.com/joho/godotenv" // 从 .env 文件加载环境变量
    "github.com/spf13/viper" // 读取 config.yaml 文件并支持环境变量替换
)

type Config struct {
    Server    ServerConfig    `mapstructure:"server"`
    Database  DatabaseConfig  `mapstructure:"database"`
    Redis     RedisConfig     `mapstructure:"redis"`
    JWT       JWTConfig       `mapstructure:"jwt"`
    RateLimit RateLimitConfig `mapstructure:"rate_limit"`
}

type ServerConfig struct {
    Port string `mapstructure:"port"`
    Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    User     string `mapstructure:"user"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    Charset  string `mapstructure:"charset"`
}

type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Password string `mapstructure:"password"`
    DB       int    `mapstructure:"db"`
}

type JWTConfig struct {
    Secret string `mapstructure:"secret"`
    Expire string `mapstructure:"expire"`
}

type RateLimitConfig struct {
    RequestsPerMinute int `mapstructure:"requests_per_minute"`
}

func LoadConfig(path string) (*Config, error) {
    err := godotenv.Load(path + "/.env")
    if err != nil {
        fmt.Println("Warning: .env file not found, using system environment variables")
    }
    
    viper.AutomaticEnv()
    viper.SetEnvPrefix("")
    
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(path + "/configs") // 可能有问题
    
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("failed to read config file: %w", err)
    }
    
    var config Config
    if err := viper.Unmarshal(&config); err != nil {
        return nil, fmt.Errorf("failed to unmarshal config: %w", err)
    }
    
    return &config, nil
}

func GetEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}
