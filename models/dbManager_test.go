package models

import (
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"testing"
	"time"
	"os"
)

func TestNewDbManager(t *testing.T) {
	tempDb, m := initializeDatabase(t)

	m.Close()
	os.Remove(tempDb)

}

func TestNewApp(t *testing.T) {
	tempDb, m := initializeDatabase(t)

	var app = App{
		bson.NewObjectId(),
		"",
		"",
		"",
		time.Time{},
	}

	err := m.NewApp(app)

	if err == nil {
		t.Errorf("Persisted app with invalid parameters")
	}

	_, err = createApp(m)

	if err != nil {
		t.Errorf("app not created")
	}

	m.Close()
	os.Remove(tempDb)

}

func TestAllApps(t *testing.T){
	tempDb, m := initializeDatabase(t)
	_, _ = createApp(m)
	_, _ = createApp(m)

	apps,err := m.AllApps()

	if err != nil {
		t.Errorf("%v", err)
	}
	if len(apps) != 2 {
		t.Errorf("Expected 2 apps. Recieved %v", len(apps))
	}

	m.Close()
	os.Remove(tempDb)
}

func TestNewEvent(t *testing.T) {
	tempDb, m := initializeDatabase(t)
	app, err := createApp(m)

	badEvent1 := Event{
		bson.NewObjectId(),
		bson.NewObjectId(),
		"batian.io",
		"requests",
		time.Now(),
		bson.M{
	      "message": "Does not compute",
	      "method": "GET",
	      "path": "/ap1/v1/events",
	      "status_code": 500,
	    },
	}

	err = m.NewEvent(badEvent1)

	if err == nil {
		t.Errorf("Persisted event with invalid AppID")
	}

	badEvent2 := Event{
		bson.NewObjectId(),
		"",
		"",
		"",
		time.Time{},
		nil,
	}

	err = m.NewEvent(badEvent2)

	if err == nil {
		t.Errorf("Persisted event with invalid parameters")
	}

	goodEvent := Event{
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
	}

	err = m.NewEvent(goodEvent)

	if err != nil {
		t.Errorf(err.Error())
	}



	m.Close()
	os.Remove(tempDb)
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

func initializeDatabase(t *testing.T) (string, *DbManager) {
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}

	m, err := NewDbManager(tempDb)

	if err != nil {
		t.Errorf("Error when calling NewDbManager: %v", err)
	}

	return tempDb, m
}

func createApp(m *DbManager) (App, error) {
	var app = App{
		bson.NewObjectId(),
		"batian.io",
		"JumpingJacks",
		"Golang",
		time.Now(),
	}

	err := m.NewApp(app)
	return app, err
}