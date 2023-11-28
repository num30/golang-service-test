.PHONY: servicetest.test test build lint deps 

build: servicetest.build	
	env CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -o bin/api-service cmd/main.go

test:
	go test -v ./...

# Go lint
lint:
	golangci-lint run

# service tests
servicetest.run:
	go test ./test/stest -tags servicetest  -v -count=1


servicetest.build:
	env CGO_ENABLED=0 go test ./test/stest -tags servicetest -v -c -o bin/service-test

servicetest.docker:
	docker build . -f Test.Dockerfile  -t service-test