#!/bin/bash
cd devices
if [ $1 == "simple" ];
then
  go test
fi
if [ $1 == "coverage" ];
then
  go test -v -coverprofile=../cover.out
fi
if [ $1 == "coverage-html" ];
then
  go test -v -coverprofile=../cover.out
  go tool cover -html=../cover.out -o ../cover.html
fi