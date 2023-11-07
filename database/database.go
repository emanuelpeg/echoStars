package database

import (
	"echoStars/util"
	bolt "go.etcd.io/bbolt"
	"os"
)

var ConfigFileName string = "database/config.json"

type BoltDB interface {
	Open() (*bolt.DB, error)
}

type BoltDBImpl struct {
	FileName string `json:"file_name"`
	FileMode uint   `json:"file_mode"`
}

func newBoltDB(fileName string) (BoltDB, error) {
	var boltDb BoltDBImpl
	err := util.ReadJSONFile(fileName, &boltDb)
	if err != nil {
		return nil, err
	}

	return boltDb, nil
}

func NewBoltDB(configFile string) (BoltDB, error) {
	return newBoltDB(configFile)
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, os.FileMode(database.FileMode), nil)
}
