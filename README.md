# Rest API integration test example
In this repo you will find: 
- simple [Rest API server](/pkg/router/router.go)
- [integration test](/test/integration/rest_service_test.go) for the API
- [Docker file for integration test build](Int.Dockerfile)
- Script to build [integration test image](/Makefile)
- [Github Action pipeline](/.github/workflows/build.yaml) to build and push integration test image 
