package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/ishuah/batian/models"
)

func NewEvent(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events := make(models.Events, 0)
		body, _ := ioutil.ReadAll(r.Body)

		json.Unmarshal(body, &events)
		fmt.Printf("%#s", events[0].Timestamp)
		for _, event := range events {

			event.Init()

			err := db.NewEvent(event)

			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
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