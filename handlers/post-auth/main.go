package main

import (
	"github.com/kawakattsun/sam-face-auth/actions"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return actions.PostAuth(request), nil
}

func main() {
	lambda.Start(handler)
}
