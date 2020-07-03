package main

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Device struct: containing a device fields
type Device struct {
	ID          string `json:"id"`
	DeviceModel string `json:"deviceModel"`
	Name        string `json:"name"`
	Model       string `json:"model"`
	Serial      string `json:"serial"`
}

// ErrorResponse struct: contains a resposne with error message
type ErrorResponse struct {
	Message string `json:"message"`
}

// storeDevice writes a new device in table
func storeDevice(device Device) (Device, error) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	deviceItem, err := dynamodbattribute.MarshalMap(device)
	if err != nil {
		fmt.Println("Got error marshalling map:")
		fmt.Println(err.Error())
		return device, err
	}

	// write device in db and return the object
	_, err = svc.PutItem(&dynamodb.PutItemInput{
		Item:      deviceItem,
		TableName: aws.String("devices"),
	})
	return device, err

}

// retriveDevice writes a new device in table
func retriveDevice(id string) (Device, error) {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	device := Device{}

	// Perform the query
	fmt.Println("Trying to read from table: ", "devices")
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("devices"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
		return device, err
	}
	if _, ok := result.Item["id"]; !ok {
		return device, errors.New("notFound")
	}
	// Unmarshall the result in to an Device
	err = dynamodbattribute.UnmarshalMap(result.Item, &device)
	if err != nil {
		fmt.Println(err.Error())
		return device, err
	}

	return device, nil

}
