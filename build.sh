#!/bin/bash

go get -d ./...
go build -o bin/application application.go
