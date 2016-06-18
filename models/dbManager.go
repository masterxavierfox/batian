package models

import (
	"github.com/boltdb/bolt"
)

type DbManager struct {
	db *bolt.DB
}

func NewDbManager(path string) (*DbManager, error) {
	db, err := bolt.Open(path, 0644, nil)
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("logs"))
		if err != nil {
			return err
		}
		return err
	})

	if err != nil {
		return nil, err
	}
	return &DbManager{db}, nil
}

func (m *DbManager) Close() error {
	return m.db.Close()
}

// func Insert(event *Event){
// 	sessionCopy := session.Copy()
// 	defer sessionCopy.Close()

// 	collection := session.DB("batian").C("events")

// 	err := collection.Insert(&event)

// 	if err != nil {
// 		panic(err)
// 	}
// }