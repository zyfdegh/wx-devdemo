#!/bin/sh

go get gopkg.in/kataras/iris.v6
go get gopkg.in/kataras/iris.v6/adaptors/httprouter
go get github.com/bmizerany/assert

go fmt ./...

go test ./...

go build -o bin/wx-devdemo

sudo docker build -t zyfdedh/wx-devdemo .

# Run
# sudo docker run -p 80:80 zyfdedh/wx-devdemo
