package boltdb

import (
	"go.zenithar.org/common/dao/api"

	"github.com/boltdb/bolt"
)

// Default contains the basic implementation of the MongoCRUD interface
type Default struct {
	table   string
	db      string
	session *bolt.DB
}

// NewCRUDTable sets up a new Default struct
func NewCRUDTable(session *bolt.DB, db, table string) api.EntityCRUD {
	return &Default{
		db:      db,
		table:   table,
		session: session,
	}
}
