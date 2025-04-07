package storage

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageInitializer struct {
	client *minio.Client
	logger *slog.Logger
}

func NewStorageInitializer(logger *slog.Logger, endpoint, accessKey, secretKey string, useSSL bool) (*StorageInitializer, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return &StorageInitializer{client: client, logger: logger}, nil
}

func (s *StorageInitializer) InitBuckets(ctx context.Context, buckets []string) error {
	for _, bucket := range buckets {
		exists, err := s.client.BucketExists(ctx, bucket)
		if err != nil {
			return fmt.Errorf("checking bucket %s: %w", bucket, err)
		}
		if !exists {
			err = s.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{})
			if err != nil {
				return fmt.Errorf("creating bucket %s: %w", bucket, err)
			}
			s.logger.Info("Created bucket", "bucket", bucket)
		} else {
			s.logger.Info("Bucket already exists:", "bucket", bucket)
		}
	}
	s.logger.Info("Initialized buckets:", "buckets", buckets)

	return nil
}

func (s *StorageInitializer) UploadFile(ctx context.Context, bucket, objectName, filePath string) error {
	_, err := s.client.FPutObject(ctx, bucket, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return fmt.Errorf("uploading file to bucket %s: %w", bucket, err)
	}
	s.logger.Info("Uploaded file:", "object", objectName, "to bucket:", bucket)

	return nil
}

func (s *StorageInitializer) ListObjects(ctx context.Context, bucket string) ([]minio.ObjectInfo, error) {
	objects := []minio.ObjectInfo{}

	objectCh := s.client.ListObjects(ctx, bucket, minio.ListObjectsOptions{Recursive: true})
	for object := range objectCh {
		if object.Err != nil {
			return nil, fmt.Errorf("listing objects in bucket %s: %w", bucket, object.Err)
		}
		objects = append(objects, object)
	}
	s.logger.Info("Listed objects in bucket:", "bucket", bucket)

	return objects, nil
}

func (s *StorageInitializer) DeleteObject(ctx context.Context, bucket, objectName string) error {
	err := s.client.RemoveObject(ctx, bucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("deleting object %s from bucket %s: %w", objectName, bucket, err)
	}
	s.logger.Info("Deleted object:", "object", objectName, "from bucket:", bucket)

	return nil
}

func (s *StorageInitializer) GetObjectURL(ctx context.Context, bucketName, objectName string) (string, error) {
	presignedURL, err := s.client.PresignedGetObject(ctx, bucketName, objectName, 24*time.Hour, nil)
	if err != nil {
		return "", fmt.Errorf("getting presigned URL for object %s in bucket %s: %w", objectName, bucketName, err)
	}
	s.logger.Info("Generated presigned URL for object:", "object", objectName, "in bucket:", bucketName)

	return presignedURL.String(), nil
}
