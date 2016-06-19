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

func NewEvent(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var event models.Event
		err := decoder.Decode(&event)

		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		err = db.NewEvent(event)

		if err != nil {
			http.Error(w, err.Error(), 500)
		}

		w.WriteHeader(200)
	})
}

func AllEvents(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events,err := db.AllEvents()
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
		})
}