package config

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

type Config struct {
	DB struct {
		Host     string `env:"DB_HOST"`
		User     string `env:"DB_USER"`
		Password string `env:"DB_PASSWORD"`
		Name     string `env:"DB_NAME"`
		Port     string `env:"DB_PORT"`
	}
	Minio struct {
		Endpoint  string   `env:"MINIO_ENDPOINT"`
		AccessKey string   `env:"MINIO_ACCESS_KEY"`
		SecretKey string   `env:"MINIO_SECRET_KEY"`
		UserSSL   bool     `env:"MINIO_USE_SSL"`
		Buckets   []string `env:"MINIO_BUCKETS" envSeparator:","`
	}
	Redis struct {
		URI      string `env:"REDIS_URI"`
		Password string `env:"REDIS_PASSWORD"`
		DB       int    `env:"REDIS_DB"`
	}
	JWTSecret string `env:"JWT_SECRET"`
}

var (
	cfg Config
)

func LoadConfig() *Config {
	// Load Database configuration
	cfg.DB.Host = getEnv("DB_HOST", "localhost")
	cfg.DB.User = getEnv("DB_USER", "postgres")
	cfg.DB.Password = getEnv("DB_PASSWORD", "0000")
	cfg.DB.Name = getEnv("DB_NAME", "frappuccino_db")
	cfg.DB.Port = getEnv("DB_PORT", "5432")

	// Load Minio configuration
	cfg.Minio.Endpoint = getEnv("MINIO_ENDPOINT", "localhost:9000")
	cfg.Minio.AccessKey = getEnv("MINIO_ACCESS_KEY", "minioadmin")
	cfg.Minio.SecretKey = getEnv("MINIO_SECRET_KEY", "minioadmin")
	cfg.Minio.UserSSL = getEnv("MINIO_USE_SSL", "false") == "true"
	cfg.Minio.Buckets = []string{
		getEnv("MINIO_BUCKETS", "bucket1"),
		getEnv("MINIO_BUCKETS", "bucket2"),
	}

	// Load Redis configuration
	// cfg.RedisURI = getEnv("REDIS_URI", "redis:6379")
	// cfg.RedisPassword = getEnv("REDIS_PASSWORD", "")
	// cfg.RedisDB, _ = strconv.Atoi(getEnv("REDIS_DB", "0"))

	// Load JWT secret
	// cfg.JWTSecret = CreateMd5Hash(getEnv("JWT_SECRET", "not-so-secret-now-is-it?"))

	return &cfg
}

func GetConfing() *Config {
	return &cfg
}

func (c *Config) MakeConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.Name,
	)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func CreateMd5Hash(text string) string {
	hasher := md5.New()
	_, err := io.WriteString(hasher, text)
	if err != nil {
		// panic(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))
}
