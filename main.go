package main

import (
	"log"
	"net/http"
	"github.com/ishuah/batian/routes"
	"github.com/ishuah/batian/models"
)


func main() {
	router := routes.NewRouter()
	models.Init()
	
	log.Fatal(http.ListenAndServe(":5000", router))
}
