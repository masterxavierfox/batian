package models

import (
	"testing"
	"encoding/json"
	"strings"
)

func TestEvent(t *testing.T) {
	eventJson := `{ "source":"brandsight.com", 
					"measurement": "exceptions", 
					"timestamp": "2016-06-14T13:55:01.000Z", 
					"data": { 
						"status_code": 500, 
						"message": "Does not compute", 
						"path": "/ap1/v1/projects", 
						"method": "GET" 
						}
					}`

	malformedEventJson := `{ }`

	event, err := bundleEvent(malformedEventJson)

	if err != nil {
		t.Errorf("Non expected error when bundling event: %v ", err.Error())
	}

	err = event.Validate()

	if err == nil {
		t.Errorf("Malformed event passed validation")
	}
	
	event, err = bundleEvent(eventJson)

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