.PHONY: integration.test test build lint deps 

build: integration.build	
	env CGO_ENABLED=0 GOOS=linux GARCH=amd64 go build -a -o bin/api-service cmd/main.go

test:
	go test -v ./...

# Go lint
lint:
	golangci-lint run

deps:
	go install

clean:
	rm pb/gen/*.go


# integration tests

# service tests
servicetest.run:
	go test ./test/stest -tags servicetest  -v -count=1


servicetest.build:
	env CGO_ENABLED=0 go test ./test/stest -tags servicetest -v -c -o bin/service-test
