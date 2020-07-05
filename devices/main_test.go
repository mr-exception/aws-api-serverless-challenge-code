package main

import (
	"testing"
)

func TestCreateValidationFail0(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{

		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 400 {
		t.Errorf("response status code has to be 400")
	}
	if response.Body != `{"message":"device should have valid id"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestCreateValidationFail1(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{
			"id": "valid-id"
		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 400 {
		t.Errorf("response status code has to be 400")
	}
	if response.Body != `{"message":"device should have valid device model"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestCreateValidationFail2(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{
			"id": "valid-id",
			"deviceModel": "valid-device-model"
		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 400 {
		t.Errorf("response status code has to be 400")
	}
	if response.Body != `{"message":"device should have valid model"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestCreateValidationFail3(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{
			"id": "valid-id",
			"deviceModel": "valid-device-model",
			"model": "valid-model"
		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 400 {
		t.Errorf("response status code has to be 400")
	}
	if response.Body != `{"message":"device should have valid name"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestCreateValidationFail4(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{
			"id": "valid-id",
			"deviceModel": "valid-device-model",
			"model": "valid-model",
			"name": "valid-name"
		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 400 {
		t.Errorf("response status code has to be 400")
	}
	if response.Body != `{"message":"device should have valid serial"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestCreateSuccess(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "POST",
		Body: `{
			"id": "valid-id",
			"deviceModel": "valid-device-model",
			"model": "valid-model",
			"name": "valid-name",
			"serial": "valid-serial"
		}`,
	}
	var response, _ = Handler(request)
	if response.StatusCode != 201 {
		t.Errorf("response status code has to be 201")
	}
}

func TestGetFailure0(t *testing.T) {
	isTesting = true

	var params = make(map[string]string)
	params["id"] = "invalid-id"

	var request = Request{
		Path:           "/api/devices",
		PathParameters: params,
		HTTPMethod:     "GET",
	}
	var response, _ = Handler(request)

	if response.StatusCode != 404 {
		t.Errorf("response status code has to be 404 but is %d", response.StatusCode)
	}
}

func TestGetFailure1(t *testing.T) {
	isTesting = true

	var params = make(map[string]string)
	params["id"] = ""

	var request = Request{
		Path:           "/api/devices",
		PathParameters: params,
		HTTPMethod:     "GET",
	}
	var response, _ = Handler(request)

	if response.StatusCode != 404 {
		t.Errorf("response status code has to be 404 but is %d", response.StatusCode)
	}
}

func TestGetSuccess(t *testing.T) {
	isTesting = true

	var params = make(map[string]string)
	params["id"] = "valid-id"

	var request = Request{
		Path:           "/api/devices",
		PathParameters: params,
		HTTPMethod:     "GET",
	}
	var response, _ = Handler(request)
	if response.StatusCode != 200 {
		t.Errorf("response status code has to be 200 but is %d", response.StatusCode)
	}
}

func TestFailedEndpoint0(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "PUT",
	}
	var response, _ = Handler(request)
	if response.StatusCode != 404 {
		t.Errorf("response status code has to be 404 but is %d", response.StatusCode)
	}
	if response.Body != `{"message":"requested endpoint not found"}` {
		t.Errorf("body is: %s", response.Body)
	}
}

func TestFailedEndpoint1(t *testing.T) {
	isTesting = true
	var request = Request{
		Path:       "/api/devices",
		HTTPMethod: "PUT",
	}
	var response, _ = Handler(request)
	if response.StatusCode != 404 {
		t.Errorf("response status code has to be 404 but is %d", response.StatusCode)
	}
	if response.Body != `{"message":"requested endpoint not found"}` {
		t.Errorf("body is: %s", response.Body)
	}
}
