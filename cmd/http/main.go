package main

import (
	"context"
	"database/sql"
	"go-forum/internal/adapter/handler"
	"go-forum/internal/adapter/storage"
	"go-forum/pkg/config"
	"go-forum/pkg/lib/logger"
	"log"
	"os"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()
	logger := logger.SetupPrettySlog(os.Stdout)

	// Database connection
	db, err := sql.Open("postgres", cfg.MakeConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize Minio storage
	storage, err := storage.NewStorageInitializer(logger, cfg.Minio.Endpoint, cfg.Minio.AccessKey, cfg.Minio.SecretKey, cfg.Minio.UserSSL)
	if err != nil {
		log.Fatal(err)
	}
	if err = storage.InitBuckets(context.Background(), cfg.Minio.Buckets); err != nil {
		log.Fatal(err)
	}

	httpSrv := handler.NewAPIServer(
		"0.0.0.0:8080",
		db,
		storage,
		logger,
	)
	httpSrv.Run()
}
