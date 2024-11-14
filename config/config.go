package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	HttpPort      string
	RedisPassword string
	RedisHost     string
	RedisPort     string
	DbHost        string
	DbPort        string
	DbName        string
	DbUser        string
	DbPassword    string
	S3Endpoint    string
	S3BucketName  string
	S3Region      string
	S3AcessKeyID  string
	S3SecretKey   string
	S3Arn         string
}

func GetConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := &Config{
		HttpPort:      os.Getenv("HTTP_PORT"),
		RedisPassword: os.Getenv("REDIS_PASSWORD"),
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPort:     os.Getenv("REDIS_PORT"),
		DbHost:        os.Getenv("DB_HOST"),
		DbPort:        os.Getenv("DB_PORT"),
		DbName:        os.Getenv("DB_NAME"),
		DbUser:        os.Getenv("DB_USER"),
		DbPassword:    os.Getenv("DB_PASSWORD"),
		S3Endpoint:    os.Getenv("S3_ENDPOINT"),
		S3BucketName:  os.Getenv("S3_BUCKET_NAME"),
		S3Region:      os.Getenv("S3_REGION"),
		S3AcessKeyID:  os.Getenv("S3_ACCESS_KEY_ID"),
		S3SecretKey:   os.Getenv("S3_SECRET_KEY"),
		S3Arn:         os.Getenv("S3_ARN"),
	}

	return config
}
