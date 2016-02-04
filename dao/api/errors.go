package api

import (
	"errors"
	"fmt"
)

var (
	// ErrNotImplemented is the error raised when the method is not implemented
	ErrNotImplemented = errors.New("Method not implemented !")
	// ErrNoResult is the error raised when the query result no results
	ErrNoResult = errors.New("No result found !")
)

// DatabaseError is the wrapper for RethinkDB errors that allows passing more data with the message
type DatabaseError struct {
	err     error
	message string
	table   EntityTable
}

func (d *DatabaseError) Error() string {
	return fmt.Sprintf(
		"%s - DB: %s - Table : %s - %s",
		d.message,
		d.table.GetDBName(),
		d.table.GetTableName(),
		d.err,
	)
}

// NewDatabaseError creates a new DatabaseError, wraps err and adds a message
func NewDatabaseError(t EntityTable, err error, message string) *DatabaseError {
	return &DatabaseError{
		err:     err,
		table:   t,
		message: message,
	}
}
