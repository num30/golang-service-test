package main

import (
	"api-integration-test/pkg/router"
	"log"
	"net/http"
)

func main() {
	log.Println("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
