package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	)

type Model interface {
	Save() bool
}

type Event struct {
	ID			int
	Source		string
	Measurement	string
	Timestamp	time.Time
	Data	bson.M	`json:"data"`
}

type Events []Event