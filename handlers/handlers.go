package handlers

import (
	"net/http"
	"path"
	"os"
	"html/template"
	"encoding/json"
	"github.com/ishuah/batian/models"
)


func Index(w http.ResponseWriter, r *http.Request){
	cwd, _ := os.Getwd()
    tmpl := template.Must(
    	template.ParseFiles(path.Join(cwd, "templates/base.html"), path.Join(cwd, "templates/index.html")))
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