package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

const ENVInDocker = "IN_DOCKER"

type Config struct {
	Env            string `env:"JWT_SECRET" envDefault:"dev"`
	JWTSecret      string `env:"JWT_SECRET"`
	ServerHostPort string `env:"SERVER_HOST_PORT" envDefault:"localhost"`

	Listener ListenerConfig `envPrefix:"LISTENER_"`
	Postgres PostgresConfig `envPrefix:"POSTGRES_"`
	Minio    MinioConfig    `envPrefix:"MINIO_"`
	Redis    RedisConfig    `envPrefix:"REDIS_"`
}

type ListenerConfig struct {
	Host         string        `env:"HOST" envDefault:"0.0.0.0"`
	Port         int32         `env:"PORT" envDefault:"8081"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"5m"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"1m"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" envDefault:"5s"`
}

type PostgresConfig struct {
	HostPort          string        `env:"HOST_PORT,required"`
	Username          string        `env:"USER,required"`
	Password          string        `env:"PASSWORD,required"`
	DBName            string        `env:"DB_NAME" envDefault:"guide"`
	DBMaxConn         int           `env:"DB_MAX_CONN" envDefault:"10"`
	DBMaxConnLifeTime time.Duration `env:"DB_MAX_CONN_LIFE_TIME" envDefault:"5m"`
	DBMaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" envDefault:"1m"`
	DBTimeout         time.Duration `env:"DB_TIMEOUT" envDefault:"5s"`
}

type MinioConfig struct {
	HostPort string `env:"HOST_PORT,required"`
	Username string `env:"ROOT_USER,required"`
	Password string `env:"ROOT_PASSWORD,required"`
	UseSSL   bool   `env:"USE_SSL" envDefault:"false"`
}

type RedisConfig struct {
	HostPort string `env:"HOST_PORT,required"`
	Password string `env:"PASSWORD,required"`
}

func New() (*Config, error) {
	cfg := &Config{}

	if os.Getenv(ENVInDocker) == "" {
		if err := godotenv.Load("./local.env"); err != nil {
			return nil, err
		}
	}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
