package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// Device struct: containing a device fields
type Device struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Model       string `json:"model"`
	Serial      string `json:"serial"`
}

// storeDevice writes a new device in table
func storeDevice(svc dynamodbiface.DynamoDBAPI, device Device) (Device, error) {

	// maps the input device object into a dynamodb object
	deviceItem, _ := dynamodbattribute.MarshalMap(device)

	// write device in db and return the object
	_, err := svc.PutItem(&dynamodb.PutItemInput{
		Item:      deviceItem,
		TableName: aws.String("devices"),
	})
	return device, err

}

// retriveDevice writes a new device in table
func retriveDevice(svc dynamodbiface.DynamoDBAPI, id string) (Device, error) {
	device := Device{}

	// perform the query to get the requested device by id
	result, _ := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("devices"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": { // query parameter for db -> just id
				S: aws.String(id),
			},
		},
	})

	// in case if device not found
	if retrivedId, ok := result.Item["id"]; !ok || *retrivedId.S == "" {
		return device, errors.New("notFound")
	}
	// unmarshall the result in to a Device object
	dynamodbattribute.UnmarshalMap(result.Item, &device)

	return device, nil

}
