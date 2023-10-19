package computer

import (
	"encoding/json"

	"github.com/labstack/gommon/log"
	bolt "go.etcd.io/bbolt"
)

type ServerImplBoltDb struct {
}

// GetAll implements ServerDao.
func (ServerImplBoltDb) GetAll() ([]Server, error) {
	panic("unimplemented")
}

// Create implements ServerDao.
func (ServerImplBoltDb) Create(server *Server) error {
	db, err := bolt.Open("data.db", 0600, nil)
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

	bucket, err := tx.CreateBucketIfNotExists([]byte("servers"))
	if err != nil {
		log.Info(err)
		return err
	}

	buf, err := json.Marshal(*server)
	if err != nil {
		log.Info(err)
		return err
	}

	err = bucket.Put([]byte("hostname3"), buf)
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
