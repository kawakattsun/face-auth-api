package saveface

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"log"

	"github.com/kawakattsun/face-auth-api/pkg/responder"
	"github.com/kawakattsun/face-auth-api/pkg/s3service"

	"github.com/aws/aws-lambda-go/events"
)

type saveFaceRequest struct {
	UserID string `json:"user_id"`
	Image  []byte `json:"image"`
}

func requestBody(request events.APIGatewayProxyRequest) (*saveFaceRequest, error) {
	r := new(saveFaceRequest)
	if err := json.Unmarshal([]byte(request.Body), r); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *saveFaceRequest) imageBuffer() *bytes.Buffer {
	return bytes.NewBuffer(r.Image)
}

func (r *saveFaceRequest) imageHash() string {
	hash := md5.New()
	hash.Write(r.Image)

	return hex.EncodeToString(hash.Sum(nil))
}

func Handler(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	req, err := requestBody(request)
	if err != nil {
		log.Println("[ERROR] Request error occured.")
		return responder.Response500(err)
	}
	filename := req.UserID + "-" + req.imageHash() + ".jpeg"
	if err := s3service.PutCollection(filename, req.imageBuffer()); err != nil {
		log.Println("[ERROR] S3 put error.")
		return responder.Response500(err)
	}

	return responder.Response200("ok")
}
