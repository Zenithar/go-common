package mock

import "go.zenithar.org/common/dao/api"

// Default contains the basic implementation of the EntityCRUD interface
type Default struct {
	table   string
	db      string
	session interface{}
}

// NewCRUDTable sets up a new Default struct
func NewCRUDTable(session interface{}, db, table string) *Default {
	return &Default{
		db:      db,
		table:   table,
		session: session,
	}
}

// GetTableName returns table's name
func (d *Default) GetTableName() string {
	return d.table
}

// GetDBName returns database's name
func (d *Default) GetDBName() string {
	return d.db
}

// GetTable returns no table
func (d *Default) GetTable() interface{} {
	return nil
}

// GetSession returns the current session
func (d *Default) GetSession() interface{} {
	return nil
}

// Insert inserts a document into the database
func (d *Default) Insert(data interface{}) error {
	return api.ErrNotImplemented
}

// InsertOrUpdate a document occording to ID presence in database
func (d *Default) InsertOrUpdate(id interface{}, data interface{}) error {
	return api.ErrNotImplemented
}

// Find a document match given id
func (d *Default) Find(id interface{}, value interface{}) error {
	return api.ErrNotImplemented
}

// FindOneBy a couple (k = v) in the database
func (d *Default) FindOneBy(key string, value interface{}, result interface{}) error {
	return api.ErrNotImplemented
}

// FindBy a couple (k = v) in the database
func (d *Default) FindBy(key string, value interface{}, results interface{}) error {
	return api.ErrNotImplemented
}

// FindByAndCount is used to count object that matchs the (key = value) predicate
func (d *Default) FindByAndCount(key string, value interface{}) (int, error) {
	return -1, api.ErrNotImplemented
}

// Where is used to fetch documents that match th filter from the database
func (d *Default) Where(filter map[string]interface{}, results interface{}) error {
	return api.ErrNotImplemented
}

// WhereCount returns the document count that match the filter
func (d *Default) WhereCount(filter map[string]interface{}) (int, error) {
	return -1, api.ErrNotImplemented
}

// WhereAndFetchOne returns one document that match the filter
func (d *Default) WhereAndFetchOne(filter map[string]interface{}, result interface{}) error {
	return api.ErrNotImplemented
}

// WhereAndFetchLimit returns paginated list of document
func (d *Default) WhereAndFetchLimit(filter map[string]interface{}, paginator *api.Pagination, results interface{}) error {
	return api.ErrNotImplemented
}

// Update a document that match the selector
func (d *Default) Update(selector interface{}, data interface{}) error {
	return api.ErrNotImplemented
}

// UpdateID updates a document using his id
func (d *Default) UpdateID(id interface{}, data interface{}) error {
	return api.ErrNotImplemented
}

// DeleteAll documents from the database
func (d *Default) DeleteAll(pred interface{}) error {
	return api.ErrNotImplemented
}

// Delete a document from the database
func (d *Default) Delete(id interface{}) error {
	return api.ErrNotImplemented
}

// List all entities from the database
func (d *Default) List(results interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	return api.ErrNotImplemented
}

// Search all entities from the database
func (d *Default) Search(results interface{}, filter map[string]interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	return api.ErrNotImplemented
}
