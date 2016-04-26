package models

import (
	"time"
	"gopkg.in/mgo.v2/bson"
	)

type Log struct {
	Measurement	string
	Timestamp	time.Time
	Data	bson.M	`bson:",inline"`
}