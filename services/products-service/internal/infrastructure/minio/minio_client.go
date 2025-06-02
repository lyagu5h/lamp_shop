package minio

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	endpoint := os.Getenv("MINIO_ENDPOINT")       
	accessKey := os.Getenv("MINIO_ACCESS_KEY")     
	secretKey := os.Getenv("MINIO_SECRET_KEY")     
	useSSL := false                                

	var client *minio.Client
	var err error

	timeout := time.After(30 * time.Second)
	tick := time.Tick(2 * time.Second)

	bucket := os.Getenv("MINIO_BUCKET_PRODUCTS")
	if bucket == "" {
		log.Fatal("MINIO_BUCKET_PRODUCTS must be set")
	}
	region := "us-east-1"

	for {
		select {
		case <-timeout:
			log.Fatalf("could not initialize MinIO client or bucket: %v", err)
		case <-tick:
			client, err = minio.New(endpoint, &minio.Options{
				Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
				Secure: useSSL,
			})
			if err != nil {
				log.Printf("Waiting for MinIO to be ready (client init): %v", err)
				continue
			}

			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			exists, errBucket := client.BucketExists(ctx, bucket)
			if errBucket != nil {
				log.Printf("Waiting for MinIO to be ready (bucket check): %v", errBucket)
				continue
			}
			if !exists {
				if errMake := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: region}); errMake != nil {
					log.Printf("Failed to create bucket, retrying: %v", errMake)
					continue
				}
				log.Printf("Bucket %s created", bucket)
			}
			log.Printf("Connected to MinIO and bucket %s is ready", bucket)
			return client
		}
	}
}