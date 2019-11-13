package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kawakattsun/face-auth-api/internal/searchface"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return searchface.Handler(request), nil
}

func main() {
	log.Print("Start searchface lambda function.\n")
	lambda.Start(handler)
}
