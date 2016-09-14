package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/ishuah/batian/models"
	"gopkg.in/mgo.v2/bson"
	"encoding/json"
	"os"
	"time"
	"strings"
	"io/ioutil"
)

func TestNewApp(t *testing.T){
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}
	db, _ := models.NewDbManager(tempDb)

	newApp := NewApp(db)

	goodParams := `{ "name": "batian.io", "framework": "JumpinJacks", "language": "Golang" }`

	var malformedParams = `[{ "invalid":"fields" }]`

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

	db.Close()
	os.Remove(tempDb)
}

func TestAllApps(t *testing.T){
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

	allApps := AllApps(db)
	request, response := generateRequest("GET", "/api/v1/app", "")

	allApps(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v ",response.Code, response.Body)
	}

	db.Close()
	os.Remove(tempDb)
}

func TestNewEvent(t *testing.T){
	tempDb, db := initDatabase(t)
	newEvent := NewEvent(db)

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

	goodParams, err := json.Marshal(models.Event{
		bson.NewObjectId(),
		app.ID,
		"batian.io",
		"requests",
		time.Now(),
		bson.M{
			"message": "Does not compute",
			"method": "GET",
			"path": "/ap1/v1/events",
			"status_code": 500,
		},
	})

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