package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AWS struct {
		S3 struct {
			BucketName    string `envconfig:"S3_BUCKET_NAME"`
			ShareEndPoint string `envconfig:"S3_SHARE_ENDPOINT"`
			Directory     string `envconfig:"S3_DIRECTORY"`
		}
		AccessKey string `envconfig:"AWS_ACCESS_KEY"`
		SecretKey string `envconfig:"AWS_SECRET_KEY"`
		Region    string `envconfig:"AWS_REGION"`
	}

	DB struct {
		Host     string `envconfig:"DB_HOST"`
		UserName string `envconfig:"DB_USER"`
		Password string `envconfig:"DB_PASS"`
		Port     int    `envconfig:"DB_PORT"`
		Name     string `envconfig:"DB_NAME"`
	}

	Redis struct {
		Host     string `envconfig:"REDIS_HOST"`
		Port     int    `envconfig:"REDIS_PORT"`
		Password string `envconfig:"REDIS_PASSWORD"`
	}
}

func LoadConfig() (*Config, error) {
	// load default .env file, ignore the error
	_ = godotenv.Load()

	cfg := new(Config)
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, fmt.Errorf("load config error: %v", err)
	}

	return cfg, nil
}
