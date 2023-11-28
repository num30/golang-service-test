package main

import (
	"log"
	"net/http"

	"github.com/num30/golang-service-test/pkg/router"
)

func main() {
	log.Println("listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", router.NewRouter()))
}
