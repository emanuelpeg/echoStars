package sysinfo

import (
	"echoStars/database"
	"encoding/json"
	"github.com/labstack/gommon/log"
	bolt "go.etcd.io/bbolt"
)

type InfoDao interface {
	CheckDb() (bool, error)
	SaveInfo(info *SysInfo) error
	GetInfo() (*SysInfo, error)
}

type InfoDaoBolt struct {
}

func NewInfoDao() InfoDao {
	return InfoDaoBolt{}
}

func (dao InfoDaoBolt) CheckDb() (bool, error) {
	db, err := bolt.Open(database.FileName, 0600, nil)
	if err != nil {
		log.Info(err)
		return false, err
	}
	defer db.Close()

	return true, nil
}

func (dao InfoDaoBolt) SaveInfo(info *SysInfo) error {
	db, err := bolt.Open(database.FileName, 0600, nil)
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

func (dao InfoDaoBolt) GetInfo() (*SysInfo, error) {
	db, err := bolt.Open(database.FileName, 0600, nil)
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

	var info SysInfo

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
