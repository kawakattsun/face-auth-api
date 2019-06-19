package actions

import (
	"base64"

	"github.com/aws/aws-lambda-go/events"
	"github.com/kawakattsun/domains"
	res "github.com/kawakattsun/responders/api-gateway-proxy"
)

func PostFace(request events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
	req, err := domains.GetRequestPostFace(request.Body)
	if err != nil {
		return res.Response500(err)
	}
	body, err := base64.StdEncoding.DecodeString(req.Body)
	if err != nil {
		return res.Response500(err)
	}
	face := &domains.Face{
		Name: req.Name,
		Body: body,
	}
	client := s3.getClient()
	result, err = client.Put(&face)
	if err != nil {
		return res.Response500(err)
	}

	return res.Response201(result)
}
