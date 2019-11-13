package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/kawakattsun/face-auth-api/internal/saveface"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return saveface.Handler(request), nil
}

func main() {
	log.Print("Start saveface lambda function.\n")
	lambda.Start(handler)
}
