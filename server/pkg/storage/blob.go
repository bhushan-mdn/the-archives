package storage

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/bhushan-mdn/the-archives/pkg/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var bucketName string
var region string

func Setup() {
	// ctx := context.Background()
	bucketName = config.Storage.BucketName
	region = config.Storage.Region
	endpoint := config.Storage.Endpoint
	accessKeyID := config.Storage.AccessKeyID
	secretAccessKey := config.Storage.SecretAccessKey
	useSSL := config.Storage.SSL

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	if err != nil {
		log.Fatalln(err)
	}

	MinioClient = minioClient
}

func ListAllBuckets() {
	buckets, err := MinioClient.ListBuckets(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bucket := range buckets {
		fmt.Println(bucket)
	}
}

func MakeNewBucket(bucketName string) {
	ctx := context.Background()

	err := MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: region})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := MinioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func FPutObject(fileName string, filePath string, contentType string) {
	ctx := context.Background()

	objectName := fileName

	info, err := MinioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}

func PutObject(file io.Reader, filePath string, fileType string, fileSize int64) error {
	ctx := context.Background()

	objectName := filePath
	objectSize := fileSize

	info, err := MinioClient.PutObject(ctx, bucketName, objectName, file, objectSize, minio.PutObjectOptions{
		ContentType: fileType,
	})
	if err != nil {
		log.Println(err)
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}

func GetObject(location string) (*minio.Object, error) {
	ctx := context.Background()

	objectName := location

	object, err := MinioClient.GetObject(ctx, bucketName, objectName, minio.GetObjectOptions{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	stat, err := object.Stat()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Printf("Successfully downloaded %s of size %d\n", objectName, stat.Size)
	return object, nil
}

func RemoveObject(location string) error {
	ctx := context.Background()

	objectName := location

	err := MinioClient.RemoveObject(ctx, bucketName, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
