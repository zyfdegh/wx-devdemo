FROM golang:1.8
ENV PROJECT $GOPATH/src/github.com/zyfdegh/wx-devdemo
WORKDIR $PROJECT

COPY . $PROJECT

RUN go test $(go list ./... | grep -v /vendor/) && \
	go build -o bin/wx-devdemo

EXPOSE 80
CMD ["./bin/wx-devdemo"]
