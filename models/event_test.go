package models

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
	"encoding/json"
	"strings"
	"time"
)

func TestEvent(t *testing.T) {
	app, err := bundleApp(`{ "name": "batian.io", "framework": "JumpinJacks", "language": "Golang" }`)

	eventJson, err := json.Marshal(Event{
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


	malformedEventJson1, err := json.Marshal(Event{
		"",
		"",
		"",
		"",
		time.Time{},
		nil,
	})

	malformedEventJson2, err := json.Marshal(Event{
		bson.NewObjectId(),
		app.ID,
		"",
		"",
		time.Time{},
		nil,
	})

	event, err := bundleEvent(string(malformedEventJson1[:]))

	if err != nil {
		t.Errorf("Non expected error when bundling event: %v ", err.Error())
	}

	err = event.Validate()

	if err == nil {
		t.Errorf("Malformed event passed validation")
	}

	event, err = bundleEvent(string(malformedEventJson2[:]))

	if err != nil {
		t.Errorf("Non expected error when bundling event: %v ", err.Error())
	}

	err = event.Validate()

	if err == nil {
		t.Errorf("Malformed event passed validation")
	}

	event, err = bundleEvent(string(eventJson[:]))

	if err != nil {
		t.Errorf("Non expected error when bundling event: %v ", err.Error())
	}

	err = event.Validate()

	if err != nil {
		t.Errorf("Non expected error when validating event: %v ", err.Error())
	}
}

func bundleEvent(eventJson string) (*Event, error) {
	event := InitEvent()
	decoder := json.NewDecoder(strings.NewReader(eventJson))
	err := decoder.Decode(&event)

	if err != nil {
		return nil, err
	}
	return &event, nil
}