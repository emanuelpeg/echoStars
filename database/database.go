package database

import (
	"echoStars/util"
	"fmt"
	"path/filepath"

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
	dbDirectory := util.ApplicationConfig.GetString("database.file.directory")
	dbFileName := util.ApplicationConfig.GetString("database.file.name")
	dbFilePath := filepath.Join(dbDirectory, string(filepath.Separator), dbFileName)
	fmt.Println("Connected to db: ", dbFilePath)

	boltDb := BoltDBImpl{
		FileName: dbFilePath,
	}

	return boltDb, nil
}

func (database BoltDBImpl) Open() (*bolt.DB, error) {
	return bolt.Open(database.FileName, 0600, nil)
}
