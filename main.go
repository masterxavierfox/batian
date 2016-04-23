package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Batian 0.0.0")
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
