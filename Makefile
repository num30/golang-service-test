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

integration.test:
	go test ./test/integration -tags integration  -v -count=1


integration.build:
	env CGO_ENABLED=0 GOOS=linux GARCH=amd64 go test ./test/integration -tags integration -v -a -c -o bin/integration

integration.docker: integration.build
	docker build .
