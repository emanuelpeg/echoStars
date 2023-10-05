package dataBase

import (
	"echoStars/computer"
	"encoding/json"
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

func SaveInfo(fileName string, info *computer.SysInfo) error {
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Info(err)
		return err
	}
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		log.Info(err)
		return err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte("localhost"))
	if err != nil {
		log.Info(err)
		return err
	}

	buf, err := json.Marshal(*info)
	if err != nil {
		log.Info(err)
		return err
	}

	err = bucket.Put([]byte("info"), buf)
	if err != nil {
		log.Info(err)
		return err
	}

	if err := tx.Commit(); err != nil {
		log.Info(err)
		return err
	}

	return nil
}

func GetInfo(fileName string) (*computer.SysInfo, error) {
	db, err := bolt.Open(fileName, 0600, nil)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer tx.Rollback()

	bucket := tx.Bucket([]byte("localhost"))
	buf := bucket.Get([]byte("info"))

	if buf == nil {
		return nil, nil
	}

	var info computer.SysInfo

	err = json.Unmarshal(buf, &info)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Info(err)
		return nil, err
	}

	return &info, nil
}
