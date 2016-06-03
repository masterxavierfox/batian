package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	)

type Model interface {
	Save() bool
}

type Event struct {
	Source		string
	Measurement	string
	Timestamp	time.Time
	Data	bson.M	`json:"data"`
}

func (e *Event) Save() bool {
	Insert(e)
	return true
}