package s3service

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	collectionFolder = "/collection/"
)

var (
	awSession  *session.Session
	backetName string
)

func init() {
	config := aws.NewConfig().WithRegion(os.Getenv("AWS_REGION"))
	if os.Getenv("USE_DOCKER") == "1" {
		config.WithEndpoint(os.Getenv("MINIO_ENDPOINT")).WithS3ForcePathStyle(true)
		backetName = "fixture"
	} else {
		backetName = os.Getenv("S3_BUCKET_NAME_COLLECTION")
	}
	awSession = session.Must(session.NewSession(config))
}

func PutCollection(filename string, body io.Reader) error {
	return Put(collectionFolder+filename, body)
}

func Put(filename string, body io.Reader) error {
	uploader := s3manager.NewUploader(awSession)
	input := &s3manager.UploadInput{
		Body:   aws.ReadSeekCloser(body),
		Bucket: aws.String(backetName),
		Key:    aws.String(filename),
	}
	if _, err := uploader.Upload(input); err != nil {
		return err
	}

	return nil
}
