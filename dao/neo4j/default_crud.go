package neo4j

import "esec.sogeti.com/common/dao/api"

// Default contains the basic implementation of the EntityCRUD interface
type Default struct {
	label string
}

// NewCRUDTable sets up a new Default struct
func NewCRUDTable(label string) *Default {
	return &Default{
		label: label,
	}
}

// GetTableName returns table's name
func (d *Default) GetTableName() string {
	return d.label
}

// GetDBName returns database's name
func (d *Default) GetDBName() string {
	return "neo4j"
}

// GetSession returns the current session
func (d *Default) GetSession() interface{} {
	return nil
}

// -----------------------------------------------------------------------------

// -----------------------------------------------------------------------------

// Delete a node by id
func (d *Default) Delete(id uint64) error {
	return api.ErrNotImplemented
}
