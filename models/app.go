package models

import (
	"time"
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type App struct {
	ID			bson.ObjectId `storm:"id"`
	Name		string `storm:"index"`
	Framework	string
	Language	string
	CreatedAt	time.Time
}

type AppFields struct {
	Name		string
	Framework	string
	Language	string
}

type Apps []App

func InitApp() App {
	return App{ ID: bson.NewObjectId(), CreatedAt: time.Now() }
}

func (app *App) Update(fields AppFields) error {
	if fields.Name == "" && fields.Framework == "" && fields.Language == "" {
		return errors.New("Error: you are trying to update empty fields")
	}

	if fields.Name != "" {
		app.Name = fields.Name
	}

	if fields.Framework != "" {
		app.Framework = fields.Framework
	}

	if fields.Language != "" {
		app.Language = fields.Language
	}

	return nil

}

func (app *App) Validate() error {
	var message string

	if app.ID == "" {
		return errors.New("Error: uninitialized app")
	}

	if app.Name == "" {
		message += " name field "
	}

	if app.Framework == "" {
		message += " framework field "
	}

	if app.Language == "" {
		message += " language field "
	}

	if message != "" {
		return errors.New("Error: app missing "+message)
	}

	return nil
}