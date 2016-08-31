package main

import (
	"github.com/ishuah/batian/models"
	"github.com/ishuah/batian/routes"
	"log"
	"net/http"
)

func main() {
	db, _ := models.NewDbManager("batian.db")
	allroutes := routes.BuildRoutes(db)
	router := routes.NewRouter(allroutes)

	log.Fatal(http.ListenAndServe(":5000", router))
}