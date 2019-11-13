package responder

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func Response200(body interface{}) events.APIGatewayProxyResponse {
	b, err := json.Marshal(body)
	if err != nil {
		return Response500(err)
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(b),
		Headers:    commonHeaders(),
	}
}

func Response500(err error) events.APIGatewayProxyResponse {
	log.Printf("%+v\n", err)
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Headers:    commonHeaders(),
		Body:       `{"message":"サーバエラーが発生しました。"}`,
	}
}
