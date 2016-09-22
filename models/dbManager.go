package models

import (
	"errors"
	"github.com/asdine/storm"
	"gopkg.in/mgo.v2/bson"
)

type DbManager struct {
	db *storm.DB
}

func NewDbManager(path string) (*DbManager, error) {
	db, _ := storm.Open(path)
	db.Init(&Event{})
	db.Init(&App{})
	
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
	var app App
	err = m.db.One("ID", event.AppID, &app)
	if err != nil {
		return errors.New("Error: AppID provided does not exist")
	}
	err = m.db.Save(&event)
	return err
}

func (m *DbManager) NewApp(app App) error {
	err := app.Validate()
	if err != nil {
		return err
	}

	err = m.db.Save(&app)
	return err
}

func (m *DbManager) GetApp(appID string) (App, error) {
	var app App
	err := m.db.One("ID", bson.ObjectIdHex(appID), &app)
	if err != nil {
		return app, err
	}
	return app, nil
}

func (m *DbManager) DeleteApp(appID string) error {
	var app App

	err := m.db.One("ID", bson.ObjectIdHex(appID), &app)
	if err != nil {
		return err
	}

	err = m.db.DeleteStruct(&app)
	if err != nil {
		return err
	}

	return nil
}

func (m *DbManager) UpdateApp(app App) error {
	err := app.Validate()
	if err != nil {
		return err
	}

	err = m.db.Update(&app)
	return err
}

func (m *DbManager) AllApps() (Apps, error) {
	var apps Apps
	err := m.db.All(&apps)
	if err != nil {
		return nil, err
	}
	return apps, nil
}

func (m *DbManager) AllEvents() (Events, error) {
	var events Events
	err := m.db.AllByIndex("Timestamp", &events)

	if err != nil {
		return nil, err
	}
	return events, nil
}
