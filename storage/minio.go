package storage

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type MinioClient interface {
}

type minioRestClient struct {
	mc *minio.Client
}

func NewMinioRestClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) MinioClient {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &minioRestClient{mc: minioClient}
}
