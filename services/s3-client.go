package services

import (
	"bytes"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/kawakattsun/sam-face-auth/domains"
)

type S3Client struct {
	accessKeyID     string
	secretAccessKey string
	region          string
	BucketName      string
}

var instance *S3Client
var once sync.Once

func GetS3Client() *S3Client {
	once.Do(func() {
		instance = &S3Client{
			accessKeyID:     os.Getenv("ACCSESS_KEY_ID"),
			secretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
			region:          os.Getenv("REGION"),
			BucketName:      os.Getenv("BUCKET_NAME"),
		}
	})
	return instance
}

func (s *S3Client) Put(image domains.Image) (string, error) {
	creds := credentials.NewStaticCredentials(s.accessKeyID, s.secretAccessKey, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(s.region)},
	)
	if err != nil {
		return "", err
	}
	client := s3.New(sess)
	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(image.GetName()),
		Body:   bytes.NewReader(image.GetBody()),
	})
	if err != nil {
		return "", err
	}
	return "ok", nil
}
