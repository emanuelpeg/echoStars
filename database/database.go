package database

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
	"os"
)

type BoltDB interface {
	Open() (*bolt.DB, error)
}

type BoltDBImpl struct {
	FileName string `json:"file_name"`
	FileMode uint   `json:"file_mode"`
}

func NewBoltDB() (BoltDB, error) {
	configFile, err := os.Open("database/config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var boltDb BoltDBImpl
	decoder := json.NewDecoder(configFile)
	if err := decoder.Decode(&boltDb); err != nil {
		return nil, err
	}

	return boltDb, nil
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, os.FileMode(database.FileMode), nil)
}
