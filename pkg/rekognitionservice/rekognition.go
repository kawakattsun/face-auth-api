package rekognitionservice

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

const (
	faceMatchThreshold = 99.4
)

var reko *rekognition.Rekognition

func init() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	reko = rekognition.New(sess)
}

func SearchFace(image []byte) (*rekognition.SearchFacesByImageOutput, error) {
	input := &rekognition.SearchFacesByImageInput{
		CollectionId:       aws.String(os.Getenv("REKOGNITION_COLLECTION_ID")),
		FaceMatchThreshold: aws.Float64(faceMatchThreshold),
		Image: &rekognition.Image{
			Bytes: image,
		},
	}
	output, err := reko.SearchFacesByImage(input)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func CollectFace(bucket, key string) (string, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	reko := rekognition.New(sess)
	input := &rekognition.IndexFacesInput{
		CollectionId: aws.String(os.Getenv("REKOGNITION_COLLECTION_ID")),
		Image: &rekognition.Image{
			S3Object: &rekognition.S3Object{
				Bucket: aws.String(bucket),
				Name:   aws.String(key),
			},
		},
	}
	output, err := reko.IndexFaces(input)
	if err != nil {
		return "", err
	}

	return *output.FaceRecords[0].Face.FaceId, err
}
