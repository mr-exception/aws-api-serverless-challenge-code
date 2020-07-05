package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// GetItem stub for dynamodb
func (m *stubDynamoDB) GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	// Make response
	var requestedID string = *input.Key["id"].S

	id := dynamodb.AttributeValue{}
	deviceModel := dynamodb.AttributeValue{}
	name := dynamodb.AttributeValue{}
	serial := dynamodb.AttributeValue{}
	model := dynamodb.AttributeValue{}
	resp := make(map[string]*dynamodb.AttributeValue)
	resp["id"] = &id
	resp["name"] = &name
	resp["deviceModel"] = &deviceModel
	resp["model"] = &model
	resp["serial"] = &serial
	if requestedID == "valid-id" {
		id.SetS("valid-id")
		deviceModel.SetS("deviceModel")
		name.SetS("name")
		serial.SetS("serial")
		model.SetS("model")
	} else {
		id.SetS("")
		deviceModel.SetS("")
		name.SetS("")
		serial.SetS("")
		model.SetS("")
	}
	// Returned canned response
	output := &dynamodb.GetItemOutput{
		Item: resp,
	}
	return output, nil
}

// PutItem sub for dynamodb
func (m *stubDynamoDB) PutItem(input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	// Make response

	id := dynamodb.AttributeValue{}
	deviceModel := dynamodb.AttributeValue{}
	name := dynamodb.AttributeValue{}
	serial := dynamodb.AttributeValue{}
	model := dynamodb.AttributeValue{}
	resp := make(map[string]*dynamodb.AttributeValue)

	resp["id"] = &id
	id.SetS("valid-id")

	resp["name"] = &name
	name.SetS("name")

	resp["deviceModel"] = &deviceModel
	deviceModel.SetS("deviceModel")

	resp["model"] = &model
	model.SetS("model")

	resp["serial"] = &serial
	serial.SetS("serial")

	// Returned canned response
	output := &dynamodb.PutItemOutput{}
	return output, nil
}

// test to get an item that exists
func TestSuccessGetItemFound(t *testing.T) {
	isTesting = true
	svc := createDynamodbSession()
	device, err := retriveDevice(svc, "valid-id")
	if err != nil {
		t.Errorf("Error: calling Dynamodb %d", err)
	}

	// fetched device is not empty
	if device.ID == "" || device.Name == "" || device.Model == "" || device.DeviceModel == "" || device.Serial == "" {
		t.Errorf("Error: device not found")
	}
}

// test to get an item that doesnt exists
func TestSuccessGetItemNotFound(t *testing.T) {
	isTesting = true
	svc := createDynamodbSession()
	device, err := retriveDevice(svc, "invalid-id")
	if err == nil || err.Error() != "notFound" {
		t.Errorf("not found item get failed")
	}

	// fetched device have to be empty
	if device.ID != "" {
		t.Errorf("Error in fetch device")
	}
}

// test to store an item
func TestSuccessPutItemSuccess(t *testing.T) {
	isTesting = true
	svc := createDynamodbSession()
	var device = Device{
		ID:          "valid-id",
		Name:        "name",
		Serial:      "serial",
		DeviceModel: "deviceModel",
		Model:       "model",
	}
	device, err := storeDevice(svc, device)
	if err != nil {
		t.Errorf("Error: calling Dynamodb %d", err)
	}

	// fetched device is not empty
	if device.ID == "" || device.Name == "" || device.Model == "" || device.DeviceModel == "" || device.Serial == "" {
		t.Errorf("Error: device not found")
	}
}
