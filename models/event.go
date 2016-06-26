package models

import (
	"time"
	"errors"
	"gopkg.in/mgo.v2/bson"
	)

type Model interface {
	Save() bool
}

type Event struct {
	ID			bson.ObjectId `storm:"id"`
	Source		string `storm:"index"`
	Measurement	string `storm:"index"`
	Timestamp	time.Time `storm:"index"`
	Data	bson.M	`json:"data"`
}

type Events []Event

func (event *Event) Validate() error {
	var message string
	if event.Source == "" {
		message += " source field "
	}

	if event.Measurement == "" {
		message += " measurement field "
	}

	if event.Timestamp.IsZero() {
		message += " timestamp field "
	}

	if event.Data == nil {
		message += " data field "
	}

	if message != "" {
		return errors.New("Error: event missing "+message)
	}
	return nil
}