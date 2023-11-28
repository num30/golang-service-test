# API Service Test Example

## Objective

Service tests are automated tests executed against an instance of running service, usually after deployment, to ensure that the service works properly with other parts of the system. For example, if a service uses DB then we want to test its interface and ensure that the data from DB is returned.  

In [testing pyramid](https://martinfowler.com/articles/practical-test-pyramid.html) they are located between unit-test and end-to-end tests. The main target of those tests is service API so it's fairly easy to write them which makes them a crucial part of software quality assurance

## Scope 

This repository demonstrates one of the approaches to write, run and maintain service tests for a Golang service. 

## What is in the repository?

In this repo you will find: 
- simple [Rest API server](/pkg/router/router.go)
- [service test](/test/stest/rest_service_test.go) for the API
- [Docker file for service test build](Test.Dockerfile)
- Script to build [service test image](/Makefile)
- [Github Action pipeline](/.github/workflows/build.yaml) to build and push service test image 

## Let's break it down
We have our basic api service declared in [router.go](/pkg/router/router.go) file with three methods:
 - `GET /ping` - returns 200 OK and pong in response
 - `GET /boxes/{box_id}` - returns Json with box content by it's ID
 - `PUT /boxes/{box_id}` - updates box content by it's ID

We have a [test](/test/stest/rest_service_test.go) that checks that our service works as expected. Test is written as a go tests the only difference is that we don't access any methods directly but use HTTP calls to call our service.  



## Building Test Image 
Running locally is good for one time verification. In order to run test as part of CI/CD pipeline we need to build and push service test image. To do that we will build a test binary first 
```
env CGO_ENABLED=0 go test ./test/stest -tags servicetest -v -c -o bin/service-test
```

then we will build a docker image using [Test.Dockerfile](Test.Dockerfile) created specifically for service tests. 

```
docker build -f Test.Dockerfile -t service-test .
```

You probably want to do it in a pipeline so [here is an example of Github Action pipeline](.github/workflows/build.yaml) that does that for you.
We build application image with two tags `[short-commit-sha]` and `[branch-name]-[build-number]` test image is being built with the same tags but with  `-test` suffix. That way we use same Docker repository for both application and test images. You may consider having separate repository although I would suggest to stick to some naming conventions. This will make it easier to deal with running the test later. 


## Runnig Tests

### Locally
Our service test have a `servicetest` build tag that prevents our test from running with `go test ./...` command. 

To run our test locally: 
```
go test ./test/stest -tags servicetest  -v -count=1
```

### Docker

Test are packaged into a docker image so we can run them in a container. This docker container is easy to run in a K8s as part of helm test, k8s job, or as a step in a CD pipeline.


### Helm Chart 

If you deploy your service to Kubernetes then deploying service test as a cron job will be a good idea. You can trigger them CD pipeline or manually. This repo contains and exemplar helm chart that deploys [service](helm/boxes-api/templates/deployment.yaml) and [service test](helm/boxes-api/templates/tests/service-test-job.yaml) to Kubernetes.


## Tests Results 
If you run test as a container then it will finish executions with `FAILED` status in case when any of test cases failed. Otherwise it will finish without any errors.
