package main

import (
	"encoding/json"
	"fmt"
)

func createDevice(request Request) (Response, error) {
	var body = request.Body
	var inputs = Device{}
	json.Unmarshal([]byte(body), &inputs)

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
		} else {
			return Response{Body: string(err.Error()), StatusCode: 500}, nil
		}
	}

	result, err := json.Marshal(device)

	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}
	return Response{Body: string(result), StatusCode: 201}, nil
}
