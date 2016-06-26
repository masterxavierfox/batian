package models

import (
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
	"testing"
	"time"
	"os"
)

func TestNewDbManager(t *testing.T) {
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}

	m, err := NewDbManager(tempDb)
	if err != nil {
		t.Errorf("Error when calling NewDbManager: %v", err)
	}

	m.Close()
	os.Remove(tempDb)

}

func TestNewEvent(t *testing.T) {
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}

	m, err := NewDbManager(tempDb)

	if err != nil {
		t.Errorf("Error when calling NewDbManager: %v", err)
	}

	var event = Event{
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

	err = m.NewEvent(event)

	if err != nil {
		t.Errorf("event not created")
	}

	m.Close()
	os.Remove(tempDb)
}

func TestAllEvents(t *testing.T) {
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}

	m, err := NewDbManager(tempDb)

	event1 := Event{
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

	event2 := Event{
		bson.NewObjectId(),
		"batian.io",
		"queries",
		time.Now(),
		bson.M{
	      "latency": "0.2328",
	      "query": "select * from projects",
	    },
	}

	m.NewEvent(event1)
	m.NewEvent(event2)

	events,err := m.AllEvents()

	if err != nil {
		t.Errorf("%v", err)
	}
	if len(events) != 2 {
		t.Errorf("Expected 2 events. Recieved %v", len(events))
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