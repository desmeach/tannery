package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Env    string      `yaml:"env"`
	Mongo  MongoConfig `yaml:"mongo"`
	Http   HttpConfig  `yaml:"http"`
	JWT    JWTConfig   `yaml:"jwt"`
	Redis  RedisConfig `yaml:"redis"`
	Common CommonConfig
}

type CommonConfig struct {
	GracefulTimeout time.Duration `env:"GRACEFUL_TIMEOUT" envDefault:"30s"`
}

type MongoConfig struct {
	URI      string `yaml:"uri" env:"MONGO_URI"`
	Username string `yaml:"username" env:"MONGO_USERNAME"`
	Password string `yaml:"password" env:"MONGO_PASSWORD"`
	CoreDB   string `yaml:"core_db" env:"MONGO_CORE_DB"`
	AuthDB   string `yaml:"auth_db" env:"MONGO_AUTH_DB"`
	MaxPool  int    `yaml:"max_pool" env:"MONGO_MAX_POOL"`
	MinPool  int    `yaml:"min_pool" env:"MONGO_MIN_POOL"`
}

type HttpConfig struct {
	Host           string        `yaml:"host" env:"HTTP_HOST"`
	Port           string        `yaml:"port" env:"HTTP_PORT"`
	ReadTimeout    time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT"`
	WriteTimeout   time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT"`
	MaxHeaderBytes int           `yaml:"max_header_bytes" env:"MAX_HEADER_BYTES"`
	EnableCORS     bool          `yaml:"enable_cors" env:"ENABLE_CORS"`
}

type JWTConfig struct {
	Secret     string        `yaml:"secret" env:"JWT_SECRET"`
	AccessTTL  time.Duration `yaml:"access_ttl" env:"JWT_ACCESS_TTL"`
	RefreshTTL time.Duration `yaml:"refresh_ttl" env:"JWT_REFRESH_TTL"`
	Issuer     string        `yaml:"issuer" env:"JWT_ISSUER"`
}

type RedisConfig struct {
	Host     string `yaml:"host" env:"REDIS_HOST"`
	Port     string `yaml:"port" env:"REDIS_PORT"`
	Password string `yaml:"password" env:"REDIS_PASSWORD"`
	DB       int    `yaml:"db" env:"REDIS_DB"`
	PoolSize int    `yaml:"pool_size" env:"REDIS_POOL_SIZE"`
}

func MustLoad() *Config {
	if err := godotenv.Load(); err != nil {
		panic("error loading .env file: " + err.Error())
	}

	path := fetchConfigPath()

	if path == "" {
		panic("config file path is empty")
	}

	return MustLoadWithPath(path)
}

func MustLoadWithPath(path string) *Config {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var path string

	flag.StringVar(&path, "config_path", "", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
		fmt.Println("CONFIG_PATH", path)
	}

	return path
}
