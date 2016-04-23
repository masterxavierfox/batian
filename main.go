package main

import (
	"log"
	"net/http"
	"github.com/ishuah/batian/conf"
)


func main() {
	conf.Route()
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
