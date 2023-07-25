// minio/minio.go
package minio

import (
	"log"
	// "fmt"
	"os"
	"strings"
	"context"

	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MinioConnection() (*minio.Client, error) {
	// Inisialisasi koneksi dengan server MinIO
	// Load environment variables from .env file
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	endpoint := os.Getenv("MINIO_ENDPOINT")
	accessKeyID := os.Getenv("MINIO_ACCESS_KEY")
	secretAccessKey := os.Getenv("MINIO_SECRET_KEY")
	useSSL := strings.ToLower(os.Getenv("MINIO_USE_SSL")) == "true"
	bucketName := os.Getenv("MINIO_BUCKET_NAME")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}
	// Create a new bucket if it doesn't exist.
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalf("Error checking if bucket exists: %s", err)
	}
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalf("Error creating bucket: %s", err)
		}
		log.Printf("Bucket '%s' created successfully.\n", bucketName)
	} else {
		log.Printf("Bucket '%s' already exists.\n", bucketName)
	}
	// log.Printf("%#v\n", minioClient)
	return minioClient, nil
}
