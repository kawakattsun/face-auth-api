package searchface

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/kawakattsun/face-auth-api/pkg/rekognitionservice"
	"github.com/kawakattsun/face-auth-api/pkg/responder"
	"github.com/kawakattsun/face-auth-api/pkg/s3service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/service/rekognition"
)

type searchFaceRequest struct {
	Image []byte `json:"image"`
}

func requestBody(request events.APIGatewayProxyRequest) (*searchFaceRequest, error) {
	r := new(searchFaceRequest)
	if err := json.Unmarshal([]byte(request.Body), r); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *searchFaceRequest) imageBuffer() *bytes.Buffer {
	return bytes.NewBuffer(r.Image)
}

type searchFaceResponse struct {
	Rekognition *rekognition.SearchFacesByImageOutput `json:"rekognition"`
}

func responseBody(output *rekognition.SearchFacesByImageOutput) *searchFaceResponse {
	r := &searchFaceResponse{
		Rekognition: output,
	}

	return r
}

func Handler(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	req, err := requestBody(request)
	if err != nil {
		log.Println("[ERROR] Request error occured.")
		return responder.Response500(err)
	}
	if err := s3service.Put("test.jpeg", req.imageBuffer()); err != nil {
		log.Println("[ERROR] S3 put error.")
		return responder.Response500(err)
	}
	output, err := rekognitionservice.SearchFace(req.Image)
	if err != nil {
		log.Println("[ERROR] SerchFace error.")
		return responder.Response500(err)
	}

	return responder.Response200(responseBody(output))
}
