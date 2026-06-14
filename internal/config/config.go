package config

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Mongo MongoConfig
	Http  HttpConfig
	JWT   JWTConfig
	Redis RedisConfig
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

type MongoConfig struct {
	URI      string
	Username string
	Password string
	CoreDB   string `mapstructure:"core_db"`
	AuthDB   string `mapstructure:"auth_db"`
	MaxPool  int    `mapstructure:"max_pool"`
	MinPool  int    `mapstructure:"min_pool"`
}

type HttpConfig struct {
	Host           string
	Port           string
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	MaxHeaderBytes int           `mapstructure:"max_header_bytes"`
	EnableCORS     bool          `mapstructure:"enable_cors"`
}

type JWTConfig struct {
	Secret     string
	AccessTTL  time.Duration `mapstructure:"access_ttl"`
	RefreshTTL time.Duration `mapstructure:"refresh_ttl"`
	Issuer     string        `mapstructure:"issuer"`
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
}

func Init(configPath string) (*Config, error) {
	// 1. load yaml config
	if err := loadYAMLConfig(configPath); err != nil {
		return nil, fmt.Errorf("failed to load YAML config: %w", err)
	}

	// 2. load .env
	if err := loadEnvFiles(); err != nil {
		return nil, fmt.Errorf("failed to load env files: %w", err)
	}

	// 3. parse config to struct
	cfg := &Config{}
	if err := unmarshalConfig(cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// 4. set variables from env
	if err := setFromEnv(cfg); err != nil {
		return nil, fmt.Errorf("failed to set env values: %w", err)
	}

	// 5. validate
	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func loadYAMLConfig(configPath string) error {
	viper.SetConfigName("main")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	return nil
}

func loadEnvFiles() error {
	envFiles := []string{
		".env",
		".env.local",
		fmt.Sprintf(".env.%s", getAppEnv()),
	}

	for _, envFile := range envFiles {
		if _, err := os.Stat(envFile); err == nil {
			if err := godotenv.Load(envFile); err != nil {
				return fmt.Errorf("error loading %s: %w", envFile, err)
			}
		}
	}

	return nil
}

func unmarshalConfig(cfg *Config) error {
	if err := viper.UnmarshalKey("app", &cfg.App); err != nil {
		return fmt.Errorf("failed to parse app config: %w", err)
	}

	if err := viper.UnmarshalKey("mongo", &cfg.Mongo); err != nil {
		return fmt.Errorf("failed to parse mongo config: %w", err)
	}

	if err := viper.UnmarshalKey("http", &cfg.Http); err != nil {
		return fmt.Errorf("failed to parse http config: %w", err)
	}

	if err := viper.UnmarshalKey("jwt", &cfg.JWT); err != nil {
		return fmt.Errorf("failed to parse jwt config: %w", err)
	}

	if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
		return fmt.Errorf("failed to parse redis config: %w", err)
	}

	return nil
}

func setFromEnv(cfg *Config) error {
	cfg.Mongo.URI = getEnv("MONGO_URI", "mongodb://localhost:27017")
	cfg.Mongo.Username = getEnv("MONGO_USERNAME", "")
	cfg.Mongo.Password = getEnv("MONGO_PASSWORD", "")

	cfg.Http.Host = getEnv("HTTP_HOST", cfg.Http.Host)
	cfg.Http.Port = getEnv("HTTP_PORT", cfg.Http.Port)

	cfg.JWT.Secret = getEnv("JWT_SECRET", "")

	cfg.Redis.Host = getEnv("REDIS_HOST", "localhost")
	cfg.Redis.Port = getEnv("REDIS_PORT", "6379")
	cfg.Redis.Password = getEnv("REDIS_PASSWORD", "")

	if env := getEnv("APP_ENV", ""); env != "" {
		cfg.App.Environment = env
	}

	return nil
}

func validateConfig(cfg *Config) error {
	if cfg.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}

	if cfg.Mongo.URI == "" {
		return fmt.Errorf("MONGO_URI is required")
	}

	return nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getAppEnv() string {
	env := os.Getenv("APP_ENV")
	if env == "" {
		return "development"
	}
	return env
}
