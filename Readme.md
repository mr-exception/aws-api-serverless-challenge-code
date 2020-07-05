# Simple AWS api-gateway with AWS lamnda and dynamodb in golang

This project is a simple rest api service in go language that uses aws lambda function and aws api gateway to create two endpoints to store and fetch devices. This service uses dynamodb for its data storage and uses serverless framework for deploy.

## how to build

1. clone the project
2. run `install.sh` to install all required packages.
3. exec `make build` to build files into `/bin` directory of project

## deploy to aws

1. check serverless configuration for aws credentials
2. go to project root directory and exec `sls deploy`


## how to test