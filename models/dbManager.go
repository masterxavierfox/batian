package models

import (
	"github.com/boltdb/bolt"
	"encoding/binary"
	"encoding/json"
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
		_, err := tx.CreateBucketIfNotExists([]byte("events"))
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

func (m *DbManager) NewEvent(event Event) error {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))
		id, _ := b.NextSequence()
		event.ID = int(id)
		buf, err := json.Marshal(event)
		if err != nil {
			return err
		}
		return b.Put(itob(event.ID), buf)
			
	})
	return err
}

func (m *DbManager) AllEvents() (Events, error) {
	var events Events
	var event Event
	err := m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))
		cur := b.Cursor()

		for k, v := cur.First(); k != nil; k, v = cur.Next() {
			err := json.Unmarshal(v, &event)
			if err != nil {
				return err
			}
			events = append(events, event)
		}

		return nil
	})
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