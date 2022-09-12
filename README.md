# Rest API integration test example

This repository demonstrates one of the approaches to write, run and maintain integration tests for a go service. 

In this repo you will find: 
- simple [Rest API server](/pkg/router/router.go)
- [integration test](/test/integration/rest_service_test.go) for the API
- [Docker file for integration test build](Int.Dockerfile)
- Script to build [integration test image](/Makefile)
- [Github Action pipeline](/.github/workflows/build.yaml) to build and push integration test image 

## Let's break it down
We have our basic api service declared in [router.go](/pkg/router/router.go) file with three methods:
 - `GET /ping` - returns 200 OK and pong in response
 - `GET /boxes/{box_id}` - returns Json with box content by it's ID
 - `PUT /boxes/{box_id}` - updates box content by it's ID

We have a [test](/test/integration/rest_service_test.go) that checks that our service works as expected. Test is written as a go tests the only difference is that we don't access any methods directly but use HTTP calls to call our service.  

## Running locally
Our integration test have a `integration` build tag that prevets our test from running with `go test ./...` command. 

To run our test locally: 
```
go test ./test/integration -tags integration  -v -count=1
```

## Building Test Image 
Running locally is good for one time verification. In order to run test as part of CI/CD pipeline we need to build and push integration test image. To do that we will build a test binary first 
```
env CGO_ENABLED=0 GOOS=linux GARCH=amd64 go test ./test/integration -tags integration -v -a -c -o bin/integration
```

then we will build a docker image using [Int.Dockerfile](Int.Dockerfile) created specifically for integration tests. 

```
docker build -f Int.Dockerfile -t integration-test .
```

You probably want to do it in a pipeline so [here is an example of Github Action pipeline](.github/workflows/build.yaml) that does that for you.
We build application image with two tags `[short-commit-sha]` and `[branch-name]-[build-number]` test image is being built with the same tags but with  `-test` suffix. That way we use same Docker repository for both application and test images. You may consider having separate repository although I would suggest to stick to some naming conventions. This will make it easier to deal with running the test later. 


## Runnig Tests In K8s

### Helm Chart 

If you use helm charts to deploy you application then you can use [helm test](https://helm.sh/docs/topics/chart_tests/) to run your integration tests. 