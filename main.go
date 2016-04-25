package main

import (
	"log"
	"net/http"
	"github.com/ishuah/batian/routes"
)


func main() {
	router := routes.NewRouter()

	log.Fatal(http.ListenAndServe(":5000", router))
}
