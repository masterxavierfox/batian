package models

import (
	"github.com/asdine/storm"
)

type DbManager struct {
	db *storm.DB
}

func NewDbManager(path string) (*DbManager, error) {
	db, err := storm.Open(path)
	db.Init(&Event{})
	if err != nil {
		return nil, err
	}
	return &DbManager{db}, nil
}

func (m *DbManager) Close() error {
	return m.db.Close()
}

func (m *DbManager) NewEvent(event Event) error {
	err := event.Validate()
	if err != nil {
		return err
	}
	err = m.db.Save(&event)
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
