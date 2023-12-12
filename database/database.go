package database

import (
	"echoStars/util"
	bolt "go.etcd.io/bbolt"
)

type BoltDB interface {
	Open() (*bolt.DB, error)
}

type BoltDBImpl struct {
	FileName string `json:"file_name"`
	FileMode uint   `json:"file_mode"`
}

func NewBoltDB() (BoltDB, error) {
	boltDb := BoltDBImpl{
		FileName: util.ApplicationConfig.GetString("database.file.name"),
	}

	return boltDb, nil
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, 0600, nil)
}
