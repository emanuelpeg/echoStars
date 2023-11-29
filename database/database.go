package database

import (
	"echoStars/util"
	"strings"

	bolt "go.etcd.io/bbolt"
)

var ConfigFileName string = "database/config.json"

type BoltDB interface {
	Open() (*bolt.DB, error)
}

type BoltDBImpl struct {
	FileName string `json:"file_name"`
	FileMode uint   `json:"file_mode"`
}

func NewBoltDB() (BoltDB, error) {
	var boltDb BoltDBImpl
	localConfigFile := strings.Replace(ConfigFileName, "json", "local.json", 1)
	err := util.ReadJSONFile(localConfigFile, &boltDb)
	if err != nil {
		err := util.ReadJSONFile(ConfigFileName, &boltDb)
		if err != nil {
			return nil, err
		}
	}

	return boltDb, nil
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, 0600, nil)
}
