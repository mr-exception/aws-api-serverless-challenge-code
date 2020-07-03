package main

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse struct: contains a resposne with error message
type ErrorResponse struct {
	Message string `json:"message"`
}

func createDevice(request Request) (Response, error) {
	var body = request.Body
	var inputs = Device{}
	json.Unmarshal([]byte(body), &inputs)

	// validate the request data
	if inputs.ID == "" { // validation of id
		errorResponse := ErrorResponse{
			Message: "device should have a valid id",
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 400}, nil
	}

	if inputs.DeviceModel == "" { // validation of device model
		errorResponse := ErrorResponse{
			Message: "device should have a valid device model",
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 400}, nil
	}

	if inputs.Model == "" { // validation of model
		errorResponse := ErrorResponse{
			Message: "device should have a valid model",
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 400}, nil
	}

	if inputs.Name == "" { // validation of name
		errorResponse := ErrorResponse{
			Message: "device should have a valid name",
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 400}, nil
	}

	if inputs.Serial == "" { // validation of serial
		errorResponse := ErrorResponse{
			Message: "device should have a valid serial",
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 400}, nil
	}

	device, error := storeDevice(inputs)

	if error != nil {
		return Response{Body: error.Error(), StatusCode: 500}, error
	}

	result, error := json.Marshal(device)

	if error != nil {
		return Response{Body: error.Error(), StatusCode: 500}, error
	}
	return Response{Body: string(result), StatusCode: 201}, nil
}

func getDevice(request Request) (Response, error) {
	var id = request.PathParameters["id"]

	device, err := retriveDevice(id)

	if err != nil {
		if err.Error() == "notFound" {
			errorResponse := ErrorResponse{
				Message: fmt.Sprintf("device with id %s not found", id),
			}
			result, _ := json.Marshal(errorResponse)
			return Response{Body: string(result), StatusCode: 404}, nil
		}
		return Response{Body: string(err.Error()), StatusCode: 500}, nil

	}

	result, err := json.Marshal(device)

	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}
	return Response{Body: string(result), StatusCode: 201}, nil
}
