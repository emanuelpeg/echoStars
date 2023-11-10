package server

import (
	"echoStars/database"
	"encoding/json"
	"fmt"
	"github.com/labstack/gommon/log"
	bolt "go.etcd.io/bbolt"
)

const ServersTable = "servers"

type ServerDaoInterface interface {
	GetAll() ([]*Server, error)
	Create(server *Server) (*Server, error)
	Delete(url *string) (bool, error)
}

type ServerDaoBolt struct {
	boltDB database.BoltDB
}

func NewServerDao(configFile string) (ServerDaoInterface, error) {
	bolt, err := database.NewBoltDB(configFile)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	return ServerDaoBolt{boltDB: bolt}, nil
}
func (s ServerDaoBolt) GetAll() ([]*Server, error) {
	db, err := s.boltDB.Open()
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer db.Close()

	servers := make([]*Server, 0)
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(ServersTable))
		if b == nil {
			return fmt.Errorf("bucket servers was not found, becasue 'storage' was not found")
		}

		err := b.ForEach(func(_, value []byte) error {
			var server *Server
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

func (s ServerDaoBolt) Create(server *Server) (*Server, error) {
	db, err := s.boltDB.Open()
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		log.Info(err)
		return nil, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(ServersTable))
	if err != nil {
		log.Info(err)
		return nil, err
	}

	buf, err := json.Marshal(*server)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	err = bucket.Put([]byte(server.UrlHealth), buf)
	if err != nil {
		log.Info(err)
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		log.Info(err)
		return nil, err
	}
	// TODO should return the just created object
	return server, nil
}

func (s ServerDaoBolt) Delete(hostname *string) (bool, error) {
	db, err := s.boltDB.Open()
	defer db.Close()
	// TODO add find by 'id' in order to check if the server exists before delete.
	tx, err := db.Begin(true)
	if err != nil {
		log.Info(err)
		return false, err
	}
	defer tx.Rollback()

	bucket, err := tx.CreateBucketIfNotExists([]byte(ServersTable))
	if err != nil {
		log.Info(err)
		return false, err
	}

	//buf, err := json.Marshal(*server)
	if err != nil {
		log.Info(err)
		return false, err
	}
	err = bucket.Delete([]byte(*hostname))
	if err != nil {
		log.Info(err)
		return false, err
	}

	if err := tx.Commit(); err != nil {
		log.Info(err)
		return false, err
	}

	return true, nil
}
