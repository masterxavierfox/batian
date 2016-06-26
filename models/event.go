package models

import (
	"time"
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