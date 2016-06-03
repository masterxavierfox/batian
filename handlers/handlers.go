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

func Event(w http.ResponseWriter, r *http.Request){
	decoder := json.NewDecoder(r.Body)
	var event models.Event
	err := decoder.Decode(&event)

	if err != nil {
		panic(err)
		w.WriteHeader(500)
	}

	event.Save()
	
	w.WriteHeader(200)
}