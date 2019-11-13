package dynamodbservice

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const (
	CollectionFacesTable = "collection_faces"
)

type CollectionFacesEvent struct {
	FaceID string `dynamo:"face_id"`
	UserID string `dynamo:"user_id"`
	Date   time.Time
}

func SaveFaceID(userID, faceID string) error {
	db := dynamo.New(session.New(), &aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	table := db.Table(CollectionFacesTable)
	event := CollectionFacesEvent{
		FaceID: faceID,
		UserID: userID,
		Date:   time.Now().UTC(),
	}

	return table.Put(event).Run()
}
