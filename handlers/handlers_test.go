package handlers

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"github.com/ishuah/batian/models"
	"os"
	"strings"
	"io/ioutil"
)

func TestNewEvent(t *testing.T){
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}
	db, _ := models.NewDbManager(tempDb)
	newEvent := NewEvent(db)

	var goodParams = `{ "source":"brandsight.com", 
						"measurement": "exceptions", 
						"timestamp": "2016-06-14T13:55:01.000Z", 
						"data": { 
							"status_code": 500, 
							"message": "Does not compute", 
							"path": "/ap1/v1/projects", 
							"method": "GET" 
							} 
						}`

	var malformedParams = `{ "invalid":"fields" }`

	request, response := generateRequest("POST", "/api/v1/event", malformedParams)

	newEvent(response, request)

	if response.Code != http.StatusInternalServerError {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "500", response.Code)
	}

	request, response = generateRequest("POST", "/api/v1/event", goodParams)
	
	newEvent(response, request)
	
	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code %v:\n\tbody: %v ", "200", response.Code)
	}

	db.Close()
	os.Remove(tempDb)
}

func TestAllEvents(t *testing.T){
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}
	db, _ := models.NewDbManager(tempDb)

	allEvents := AllEvents(db)

	request, response := generateRequest("GET", "/api/v1/event", "")

	allEvents(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("Non-expected status code%v:\n\tbody: %v ", "200", response.Code)
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