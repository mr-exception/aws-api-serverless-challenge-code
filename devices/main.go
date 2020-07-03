package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Device struct: containing a device fields
type Device struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Model       string `json:"model"`
	Serial      string `json:"serial"`
}

// Response type alias from aws events
type Response events.APIGatewayProxyResponse

// Request type alias from aws events
type Request events.APIGatewayProxyRequest

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(request Request) (Response, error) {
	var response = Response{Body: "request endpoint not found", StatusCode: 404}
	if request.HTTPMethod == "GET" {
		return getDevice(request)
	}
	if request.HTTPMethod == "POST" {
		return createDevice(request)
	}
	return response, nil
}

func main() {
	lambda.Start(Handler)
}

func createDevice(request Request) (Response, error) {
	var body = request.Body
	var device = Device{}
	json.Unmarshal([]byte(body), &device)

	result, error := json.Marshal(device)

	if error != nil {
		return Response{StatusCode: 404}, error
	}
	return Response{Body: string(result), StatusCode: 201}, nil
}

func getDevice(request Request) (Response, error) {
	var id = request.PathParameters["id"]

	device := Device{ID: id, Name: "test device", DeviceModel: "test-device-model", Model: "test-model", Serial: "test-serial"}
	result, error := json.Marshal(device)

	if error != nil {
		return Response{StatusCode: 404}, error
	}
	return Response{Body: string(result), StatusCode: 201}, nil
}
