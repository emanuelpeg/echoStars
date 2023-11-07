package database

import (
	"echoStars/util"
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
	var boltDb BoltDBImpl
	err := util.ReadJSONFile("database/config.json", &boltDb)
	if err != nil {
		return nil, err
	}

	return boltDb, nil
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, os.FileMode(database.FileMode), nil)
}
