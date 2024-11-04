package config

import (
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type db struct {
	Host string `required:"true" envconfig:"DB_HOST"`
	Name string `required:"true" envconfig:"DB_NAME"`
	User string `required:"true" envconfig:"DB_USER"`
	Pass string `required:"true" envconfig:"DB_PASS"`
	Port string `required:"true" envconfig:"DB_PORT"`
}

type jwt struct {
	Key            string        `required:"true" envconfig:"JWT_KEY"`
	ExpirationTime time.Duration `required:"true" envconfig:"JWT_EXPIRATION_TIME"`
}

type Config struct {
	DB  db
	JWT jwt
}

var getWd = os.Getwd
var processEnv = envconfig.Process

func NewConfig(customPath *string) (*Config, error) {
	var newCfg Config

	wd, err := getWd()
	if err != nil {
		return nil, err
	}

	envPath := filepath.Join(wd, ".env")

	if customPath != nil {
		envPath = *customPath
	}

	_ = godotenv.Overload(envPath)
	if err = processEnv("", &newCfg); err != nil {
		return nil, err
	}

	return &newCfg, nil
}
