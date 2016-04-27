package models

import (
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

func Init(){
	var err error
	session, err = mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
}

func Insert(log *Log){
	sessionCopy := session.Copy()
	defer sessionCopy.Close()

	collection := session.DB("test").C("logs")

	err := collection.Insert(&log)

	if err != nil {
		panic(err)
	}
}