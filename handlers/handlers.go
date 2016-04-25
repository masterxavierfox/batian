package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Batian 0.0.1")
}