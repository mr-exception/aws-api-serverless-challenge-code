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

// storeDevice writes a new device in table
func storeDevice(device Device) (Device, error) {
	// starts a new session to work with database
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)

	// maps the input device object into a dynamodb object
	deviceItem, err := dynamodbattribute.MarshalMap(device)
	if err != nil { // in case mapping failed
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

	// perform the query to get the requested device by id
	fmt.Println("Trying to read from table: ", "devices")
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("devices"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": { // query parameter for db -> just id
				S: aws.String(id),
			},
		},
	})
	// in case that query failed
	if err != nil {
		fmt.Println(err.Error())
		return device, err
	}
	// in case if device not found
	if _, ok := result.Item["id"]; !ok {
		return device, errors.New("notFound")
	}
	// unmarshall the result in to a Device object
	err = dynamodbattribute.UnmarshalMap(result.Item, &device)

	// in case of mapping error
	if err != nil {
		fmt.Println(err.Error())
		return device, err
	}
	return device, nil

}
