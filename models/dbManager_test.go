package models

import (
	"github.com/boltdb/bolt"
	"io/ioutil"
	"testing"
	"os"
)

func TestNewDbManager(t *testing.T) {
	tempDb := createTempFile()
	if tempDb == "" {
		t.Skip("Cannot create temp file")
	}

	m, err := NewDbManager(tempDb)
	if err != nil {
		t.Errorf("Error when calling NewDbManager: %v", err)
	}

	createdEventsBucket := false
	m.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))
		if b != nil {
			createdEventsBucket = true
		}
		return nil
	})

	if !createdEventsBucket {
		t.Errorf("'events' bucket does not exist")
	}

	m.Close()
	os.Remove(tempDb)

}

func createTempFile() string {
	tmpDirPath := os.TempDir()
	f, err := ioutil.TempFile(tmpDirPath, "batian_dbTest")
	if err != nil {
		return ""
	}
	f.Close()
	return f.Name()
}