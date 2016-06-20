package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/ishuah/batian/models"
)

func NewEvent(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var event models.Event
		err := decoder.Decode(&event)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = db.NewEvent(event)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(200)
	})
}

func AllEvents(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events,err := db.AllEvents()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
		})
}