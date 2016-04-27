package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	)

type Model interface {
	Save() bool
}

type Log struct {
	Measurement	string
	Timestamp	time.Time
	Data	bson.M	`bson:",inline"`
}

func (l *Log) Save() bool {
	Insert(l)
	return true
}