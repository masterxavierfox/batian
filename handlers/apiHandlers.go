package handlers

import (
	"strconv"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/ishuah/batian/models"
	"github.com/ishuah/batian/engines"
)

func NewEvent(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events := make(models.Events, 0)
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.Unmarshal(body, &events)

		if len(events) == 0 {
			http.Error(w, "Error: no events received", 500)
			return
		}

		for _, event := range events {

			event.Init()

			err = db.NewEvent(event)
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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	})
}

func AppDetails(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		
		app, err := db.GetApp(params["appID"])
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	})
}

func UpdateApp(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var fields models.AppFields
		var app models.App
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&fields)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		params := mux.Vars(r)
		app, err = db.GetApp(params["appID"])

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = app.Update(fields)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		err = db.UpdateApp(app)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(app)
	})
}

func DeleteApp(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := db.DeleteApp(params["appID"])
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.WriteHeader(204)
	})
}

func AppAnalysis(db *models.DbManager) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		duration, err := strconv.Atoi("-"+params["duration"])

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		var events models.Events
		events, err = db.GetAppEvents(params["appID"], duration)
		if err != nil {
			if err.Error() == "not found" {
				http.Error(w, "No events in the given time window.", 404)
				return
			}

			http.Error(w, err.Error(), 500)
			return
		}

		reports, err := engines.AppAnalysis(events)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(reports)
	})
}