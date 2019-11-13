package collectface

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/kawakattsun/face-auth-api/pkg/dynamodbservice"
	"github.com/kawakattsun/face-auth-api/pkg/rekognitionservice"
	"github.com/kawakattsun/face-auth-api/pkg/responder"
)

func Handler(bucket, key string) interface{} {
	// todo: rekognition collectionから同一Userは削除してから登録する？
	faceID, err := rekognitionservice.CollectFace(bucket, key)
	if err != nil {
		log.Println("[ERROR] CollectFace error.")
		return responder.Response(500)
	}
	userID := filepath.Base(key[:strings.Index(key, "-")])
	if err := dynamodbservice.SaveFaceID(userID, faceID); err != nil {
		// todo: rekognition collectionから削除
		log.Println("[ERROR] SaceFaceID error.")
		return responder.Response(500)
	}

	return responder.Response(200)
}
