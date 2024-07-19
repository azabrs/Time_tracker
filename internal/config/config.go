package config

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	DB     *DB
	Server *Server
}

type DB struct {
	Name       string `envconfig:"DATABASE_NAME" required:"true"`
	DriverName string `envconfig:"DATABASE_DRIVER_NAME" required:"true"`
	DSN        string `envconfig:"DATABASE_DSN" required:"true"`
}

type Server struct {
	Host string `envconfig:"HOST" required:"true"`
	Port string `envconfig:"PORT" required:"true"`
}

func LoadConfig(ctx context.Context) *Config {
	for _, fileName := range []string{".env"} {
		err := godotenv.Load(fileName)
		if err != nil {
			log.Fatalf("error loading %s fileName : %v", fileName, err)
		}
	}

	cfg := Config{}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("cannot process envs: %v", err)
	} else {
		log.Println("Config initialized")
	}

	return &cfg
}
