package handlers

import (
	"net/http"
	"html/template"
	"encoding/json"
	"github.com/ishuah/batian/models"
)


func Index(w http.ResponseWriter, r *http.Request){
    tmpl, _ := template.ParseFiles("templates/base.html", "templates/index.html")
    tmpl.Execute(w, nil)
}

func Log(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var log models.Log
	err := decoder.Decode(&log)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
	}

	models.Insert(log)
	
	w.WriteHeader(200)
}