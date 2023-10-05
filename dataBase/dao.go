package dataBase

import (
	"github.com/labstack/gommon/log"
	bolt "go.etcd.io/bbolt"
)

func CheckDb(fileName string) (bool, error) {
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Info(err)
		return false, err
	}
	defer db.Close()

	return true, nil
}
