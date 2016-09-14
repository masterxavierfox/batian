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

func AllApps(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apps, err := db.AllApps()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(apps)
	})
}

func NewApp(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		app := models.InitApp()
		err := decoder.Decode(&app)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = db.NewApp(app)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.WriteHeader(200)
	})
}