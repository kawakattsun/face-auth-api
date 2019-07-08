package actions

import (
	"encoding/base64"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kawakattsun/sam-face-auth/domains"
	res "github.com/kawakattsun/sam-face-auth/responders"
	"github.com/kawakattsun/sam-face-auth/services"
)

func PostFace(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	req, err := domains.GetRequestPostFace(request.Body)
	if err != nil {
		return res.Response500(err)
	}
	b64data := req.Body[strings.IndexByte(req.Body, ',')+1:]
	body, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return res.Response500(err)
	}
	face := domains.GetFace(req.Name, body)
	client := services.GetS3Client()
	_, err = client.Put(face)
	if err != nil {
		return res.Response500(err)
	}

	return res.Response200OK()
}
