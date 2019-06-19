package services

import (
	"os"
	"github.com/kawakattsun/sam-face-auth/domains"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
)

type s3Client struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	BucketName      string
}

func getClient() s3Client {
	return *s3Client{
		AccessKeyId: os.Getenv("ACCSESS_KEY_ID"),
		SecretAccessKey: os.Getenv("SECRET_ACCESS_KEY"),
		Region: os.Getenv("REGION"),
		BucketName: os.Getenv("BUCKET_NAME")
	}
}

func (s *S3Client) Put(image Image) result, error{
	client := s3.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s.AccessKeyId, s.SecretAccessKey, ""),
		Region:      s.Region,
	})
	return cli.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(s.BucketName),
        Key:    aws.String(image.getName()),
        Body:   image.getBody(),
    })
}
