package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// ErrorResponse struct: contains a resposne with error message
type ErrorResponse struct {
	Message string `json:"message"`
}
type stubDynamoDB struct {
	dynamodbiface.DynamoDBAPI
}

func createDynamodbSession() dynamodbiface.DynamoDBAPI {
	if isTesting {
		// if it's testing
		return &stubDynamoDB{}
	} else {
		// if it's not testing and have to be real
		sess := session.Must(session.NewSession())
		return dynamodb.New(sess)
	}
}

// createDevice function saves a device in the database in exchange for receiving the desired inputs
// if the id device has already been saved. Edits its information.
// @param request Request
// @return Response, error
func createDevice(request Request) (Response, error) {
	// define validation rules
	var validationRules = make([]Rule, 1)
	validationRules = append(validationRules, Rule{Field: "id", Message: "device should have valid id", Pattern: "exists"})
	validationRules = append(validationRules, Rule{Field: "deviceModel", Message: "device should have valid device model", Pattern: "exists"})
	validationRules = append(validationRules, Rule{Field: "model", Message: "device should have valid model", Pattern: "exists"})
	validationRules = append(validationRules, Rule{Field: "name", Message: "device should have valid name", Pattern: "exists"})
	validationRules = append(validationRules, Rule{Field: "serial", Message: "device should have valid serial", Pattern: "exists"})

	// execute the validation
	var success, response = validateRequest(request, validationRules)

	if !success { // validation failed
		return response, nil
	}

	// convert request body input into a raw (unsaved) device object
	var body = request.Body
	var inputs = Device{}
	json.Unmarshal([]byte(body), &inputs)

	var svc = createDynamodbSession()
	device, _ := storeDevice(svc, inputs)

	// convert created device into json string
	result, _ := json.Marshal(device)

	// returns stored device as response to client
	return Response{Body: string(result), StatusCode: 201}, nil
}

// getDevice function retrives a device by passed id parameter in request url
// @param request Request
// @return Response, error
func getDevice(request Request) (Response, error) {
	// gets id from request parameters
	var id = request.PathParameters["id"]

	if id == "" { // if client didn't passed device id
		errorResponse := ErrorResponse{
			Message: fmt.Sprintf("requested endpoint not found"),
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 404}, nil
	}

	// retrives device from database
	var svc = createDynamodbSession()
	device, err := retriveDevice(svc, id)
	if err != nil {
		if err.Error() == "notFound" { // when device id not found
			errorResponse := ErrorResponse{
				Message: fmt.Sprintf("device with id %s not found", id),
			}
			result, _ := json.Marshal(errorResponse)
			return Response{Body: string(result), StatusCode: 404}, nil
		}
		return Response{Body: string(err.Error()), StatusCode: 500}, nil // other unexpected errors

	}

	// convert retrived device to json
	result, _ := json.Marshal(device)

	return Response{Body: string(result), StatusCode: 200}, nil
}
