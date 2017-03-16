package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/ishuah/batian/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"os"
	"time"
	"strings"
	"io/ioutil"
)

func TestApp(t *testing.T){
	tempDb, db := initDatabase(t)

	app := models.App{
		bson.NewObjectId(),
		"batian.io",
		"JumpingJacks",
		"Golang",
		time.Now(),
	}

	err := db.NewApp(app)

	if err != nil {
		t.Fatalf("Non-expected error while creating app %s", err.Error())
	}

	//TestAllApps
	t.Run("TestAllApps", func(t *testing.T){
		allApps := AllApps(db)
		request, response := generateRequest("GET", "/api/v1/app", "")

		allApps(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code%v:\n\tbody: %v ",response.Code, response.Body)
		}
		})

	//TestNewApp
	t.Run("TestNewApp", func(t *testing.T){
		newApp := NewApp(db)

		goodParams := `{ "name": "batian.io", "framework": "JumpinJacks", "language": "Golang" }`

		malformedParams := `[{ "invalid":"fields" }]`

		request, response := generateRequest("POST", "/api/v1/app", malformedParams)

		newApp(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "500", response.Code)
		}

		request, response = generateRequest("POST", "/api/v1/app", goodParams)

		newApp(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "200", response.Code)
		}
	})

	//TestShowApp
	t.Run("TestShowApp", func(t *testing.T){
		showApp := ShowApp(db)

		request, response := generateRequest("GET", "/api/v1/app/"+app.ID.Hex(), "")
	
		m := mux.NewRouter()
		m.HandleFunc("/api/v1/app/{appID:[a-z0-9]+}", showApp)

		m.ServeHTTP(response, request)

		
		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "200", response.Code)
		}
	})

	//TestUpdateApp
	t.Run("TestUpdateApp", func(t *testing.T){
		updateApp := UpdateApp(db)

		request, response := generateRequest("PUT", "/api/v1/app/"+app.ID.Hex(), `{ "framework": "Rails", "language": "Ruby" }`)
	
		m := mux.NewRouter()
		m.HandleFunc("/api/v1/app/{appID:[a-z0-9]+}", updateApp)

		m.ServeHTTP(response, request)
		
		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "200", response.Code)
		}
	})

	//TestNewEvent
	t.Run("TestNewEvent", func(t *testing.T) {
		newEvent := NewEvent(db)

		goodParams, err := json.Marshal(models.Event{
			bson.NewObjectId(),
			app.ID,
			"batian.io",
			"requests",
			time.Now().UTC().Add(-1 * time.Hour),
			bson.M{
				"host": "batian.io",
				"method": "GET",
				"path": "/",
				"response_time": 0.03460812568664551,
				"status_code": 200,
				"view": "web_app.views.overview",
			},
		})

		if err != nil {
			t.Fatalf("Non-expected error while marshalling json %s", err.Error())
		}

		var malformedParams = `[{ "invalid":"fields" }]`

		request, response := generateRequest("POST", "/log", malformedParams)

		newEvent(response, request)

		if response.Code != http.StatusInternalServerError {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", response.Code, response.Body)
		}

		request, response = generateRequest("POST", "/log", "["+string(goodParams[:])+"]")

		newEvent(response, request)

		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", response.Code, response.Body)
		}

	})

	//TestAppAnalysis
	t.Run("TestAppAnalysis", func(t *testing.T) {
		appAnalysis := AppAnalysis(db)

		request, response := generateRequest("GET", "/api/v1/app/"+app.ID.Hex()+"/analysis/24", "")
	
		m := mux.NewRouter()
		m.HandleFunc("/api/v1/app/{appID:[a-z0-9]+}/analysis/{duration:[0-9]+}", appAnalysis)

		m.ServeHTTP(response, request)
		
		if response.Code != http.StatusOK {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "200", response.Code)
		}

	})

	//TestDeleteApp
	t.Run("TestDeleteApp", func(t *testing.T){
		deleteApp := DeleteApp(db)

		request, response := generateRequest("DELETE", "/api/v1/app/"+app.ID.Hex(), "")
	
		m := mux.NewRouter()
		m.HandleFunc("/api/v1/app/{appID:[a-z0-9]+}", deleteApp)

		m.ServeHTTP(response, request)
		
		if response.Code != http.StatusNoContent {
			t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "204", response.Code)
		}
	})

	db.Close()
	os.Remove(tempDb)
}



func generateRequest(method string, url string, params string) (*http.Request,*httptest.ResponseRecorder){
	request, _ := http.NewRequest(method, url, strings.NewReader(params))
	response := httptest.NewRecorder()
	return request, response
}

func createTempFile() string {
	tmpDirPath := os.TempDir()
	f, err := ioutil.TempFile(tmpDirPath, "batian_dbTest")
	if err != nil {
		return ""
	}
	f.Close()
	return f.Name()
}

func initDatabase(t *testing.T) (string, *models.DbManager){
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}
	db, err := models.NewDbManager(tempDb)
	if err != nil {
		t.Errorf("Error when calling NewDbManager: %v", err)
	}

	return tempDb, db
}