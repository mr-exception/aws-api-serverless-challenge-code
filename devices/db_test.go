package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

type stubDynamoDB struct {
	dynamodbiface.DynamoDBAPI
}

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

// test to get an item that exists
func TestSuccessGetItemFound(t *testing.T) {
	svc := &stubDynamoDB{}
	device, err := retriveDevice(svc, "valid-id")
	if err != nil {
		t.Errorf("Error: calling Dynamodb %d", err)
	}
	if device.ID == "" || device.Name == "" || device.Model == "" || device.DeviceModel == "" || device.Serial == "" {
		t.Errorf("Error: device not found")
	}
}

// test to get an item that doesnt exists
func TestSuccessGetItemNotFound(t *testing.T) {
	svc := &stubDynamoDB{}
	device, err := retriveDevice(svc, "invalid-id")
	if err != nil {
		t.Errorf("Error calling Dynamodb %d", err)
	}
	if device.ID != "" {
		t.Errorf("Error in fetch device")
	}
}
