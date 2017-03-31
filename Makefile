default: build

# Build, test, run in local

dep-init:
	go get github.com/kardianos/govendor
	govendor init

dep-update:
	govendor remove +unused
	govendor add +external

build-local:
	go build -o bin/wx-devdemo
	cp -r conf/ bin/

run-local:
	./bin/wx-devdemo

test-local:
	go test ./...

clean-local:
	rm -rf bin/*

fmt:
	go fmt $(go list ./... | grep -v /vendor/)

# Build, test, run in Docker container

build:
	docker build -t zyfdedh/wx-devdemo .

run:
	docker run --rm -p 80:80 \
		-e TOKEN=${TOKEN} \
		-e APPID=${APPID} \
		-e SECRET=${SECRET} \
		-e POLLING_SEC=7100 \
		zyfdedh/wx-devdemo

push: build
	docker push zyfdedh/wx-devdemo .
