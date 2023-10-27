package computer

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	bolt "go.etcd.io/bbolt"
)

type ServerDaoBoltDb struct {
}

// GetAll implements ServerDao.
func (ServerDaoBoltDb) GetAll() ([]Server, error) {
	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer db.Close()

	servers := make([]Server, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("servers"))
		if b == nil {
			return fmt.Errorf("bucket servers was not found, becasue 'storage' was not found")
		}

		err := b.ForEach(func(_, value []byte) error {
			var server Server
			err := json.Unmarshal(value, &server)
			if err != nil {
				log.Error(err)
			}
			servers = append(servers, server)
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return servers, nil
}

// Create implements ServerDao.
func (ServerDaoBoltDb) Create(server *Server) error {
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

	err = bucket.Put([]byte(server.Hostname), buf)
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

func (ServerDaoBoltDb) Delete(hostname *string) error {
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

	//buf, err := json.Marshal(*server)
	if err != nil {
		log.Info(err)
		return err
	}
	err = bucket.Delete([]byte(*hostname))
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
