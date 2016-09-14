package models

import (
	"testing"
	"encoding/json"
	"strings"
)

func TestApp(t *testing.T) {
	appJson := `{ "name": "batian.io", "framework": "JumpinJacks", "language": "Golang" }`
	malformedAppJson := `{ }`

	app, err := bundleApp(malformedAppJson)

	if err != nil {
		t.Errorf("Non expected error when bundling app: %v ", err.Error())
	}

	err = app.Validate()

	if err == nil {
		t.Errorf("Malformed app passed validation")
	}

	app, err = bundleApp(appJson)

	if err != nil {
		t.Errorf("Non expected error when bundling app: %v ", err.Error())
	}

	err = app.Validate()

	if err != nil {
		t.Errorf("Non expected error when validating event: %v ", err.Error())
	}
}

func bundleApp(appJson string) (*App, error) {
	app := InitApp()
	decoder := json.NewDecoder(strings.NewReader(appJson))
	err := decoder.Decode(&app)

	if err != nil {
		return nil, err
	}

	return &app, nil
}