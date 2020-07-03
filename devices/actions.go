package main

import (
	"encoding/json"
	"fmt"
)

// ErrorResponse struct: contains a resposne with error message
type ErrorResponse struct {
	Message string `json:"message"`
}

// createDevice function saves a device in the database in exchange for receiving the desired inputs
// if the id device has already been saved. Edits its information.
// @param request Request
// @return Response, error
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

	// now data is validated and ready to store in database
	device, error := storeDevice(inputs)

	if error != nil { // in case of unexpected errors
		return Response{Body: error.Error(), StatusCode: 500}, error
	}
	// convert created device into json string
	result, error := json.Marshal(device)
	if error != nil {
		return Response{Body: error.Error(), StatusCode: 500}, error
	}

	return Response{Body: string(result), StatusCode: 201}, nil
}

// getDevice function retrives a device by passed id parameter in request url
// @param request Request
// @return Response, error
func getDevice(request Request) (Response, error) {
	// gets id from request parameters
	var id = request.PathParameters["id"]

	if id == "" {
		errorResponse := ErrorResponse{
			Message: fmt.Sprintf("requested endpoint not found"),
		}
		result, _ := json.Marshal(errorResponse)
		return Response{Body: string(result), StatusCode: 404}, nil
	}

	// retrives device from database
	device, err := retriveDevice(id)
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
	result, err := json.Marshal(device)
	if err != nil {
		return Response{Body: err.Error(), StatusCode: 500}, nil
	}

	return Response{Body: string(result), StatusCode: 201}, nil
}
