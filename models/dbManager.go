package models

import (
	"github.com/asdine/storm"
	"encoding/binary"
	"gopkg.in/mgo.v2/bson"
)

type DbManager struct {
	db *storm.DB
}

func NewDbManager(path string) (*DbManager, error) {
	db, err := storm.Open(path)
	if err != nil {
		return nil, err
	}
	return &DbManager{db}, nil
}

func (m *DbManager) Close() error {
	return m.db.Close()
}

func (m *DbManager) NewEvent(event Event) error {
	event.ID = bson.NewObjectId()
	err := m.db.Save(&event)
	return err
}

func (m *DbManager) AllEvents() (Events, error) {
	var events Events
	err := m.db.AllByIndex("Timestamp", &events)

	if err != nil {
		return nil, err
	}
	return events, nil
}

func itob(v int) []byte {
    b := make([]byte, 8)
    binary.BigEndian.PutUint64(b, uint64(v))
    return b
}