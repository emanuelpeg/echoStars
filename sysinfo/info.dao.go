package sysinfo

import (
	"echoStars/database"
	"encoding/json"
	"github.com/labstack/gommon/log"
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
	boltDB, err := database.NewBoltDB()
	if err != nil {
		log.Info(err)
		return false, err
	}

	db, err := boltDB.Open()
	if err != nil {
		log.Info(err)
		return false, err
	}
	defer db.Close()

	return true, nil
}

func (dao InfoDaoBolt) SaveInfo(info *SysInfo) error {
	boltDB, err := database.NewBoltDB()
	if err != nil {
		log.Info(err)
		return err
	}

	db, err := boltDB.Open()
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
	boltDB, err := database.NewBoltDB()
	if err != nil {
		log.Info(err)
		return nil, err
	}

	db, err := boltDB.Open()
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
