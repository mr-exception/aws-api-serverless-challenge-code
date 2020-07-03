package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response type alias from aws events
type Response events.APIGatewayProxyResponse

// Request type alias from aws events
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	var response = Response{Body: "request endpoint not found", StatusCode: 404}
	if request.HTTPMethod == "GET" { // so request is for getDevice function
		return getDevice(request)
	}
	if request.HTTPMethod == "POST" { // so request is for createDevice function
		return createDevice(request)
	}
	return response, nil
}

// main function of project
func main() {
	lambda.Start(Handler)
}
