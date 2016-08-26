package rethinkdb

import (
	r "gopkg.in/dancannon/gorethink.v2"

	"zenithar.org/go/common/dao/api"
)

// Default contains the basic implementation of the EntityCRUD interface
type Default struct {
	table   string
	db      string
	session *r.Session
}

// NewCRUDTable sets up a new Default struct
func NewCRUDTable(session *r.Session, db, table string) *Default {
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
	return r.Table(d.table)
}

// GetSession returns the current session
func (d *Default) GetSession() interface{} {
	return d.session
}

// Insert inserts a document into the database
func (d *Default) Insert(data interface{}) error {
	_, err := r.Table(d.table).Insert(data).RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// InsertOrUpdate a document occording to ID presence in database
func (d *Default) InsertOrUpdate(id interface{}, data interface{}) error {
	_, err := r.Table(d.table).Insert(data, r.InsertOpts{Conflict: "update"}).RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Find a document match given id
func (d *Default) Find(id interface{}, value interface{}) error {
	cursor, err := r.Table(d.table).Get(id).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.One(value); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindOneBy a couple (k = v) in the database
func (d *Default) FindOneBy(key string, value interface{}, result interface{}) error {

	cursor, err := r.Table(d.table).GetAllByIndex(key, value).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.One(result); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindBy all couples (k = v) in the database
func (d *Default) FindBy(key string, value interface{}, results interface{}) error {
	cursor, err := r.Table(d.table).Filter(func(row r.Term) r.Term {
		return row.Field(key).Eq(value)
	}).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.All(results); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindByAndCount is used to count object that matchs the (key = value) predicate
func (d *Default) FindByAndCount(key string, value interface{}) (int, error) {
	cursor, err := r.Table(d.table).Filter(func(row r.Term) r.Term {
		return row.Field(key).Eq(value)
	}).Count().Run(d.session)
	if err != nil {
		return 0, err
	}

	var count int
	if err := cursor.One(&count); err != nil {
		if err == r.ErrEmptyResult {
			return 0, api.ErrNoResult
		}
		return 0, api.NewDatabaseError(d, err, "")
	}

	return count, nil
}

// Where is used to fetch documents that match th filter from the database
func (d *Default) Where(filter interface{}, results interface{}) error {
	cursor, err := r.Table(d.table).Filter(filter).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.All(results); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereCount returns the document count that match the filter
func (d *Default) WhereCount(filter interface{}) (int, error) {
	cursor, err := r.Table(d.table).Filter(filter).Count().Run(d.session)
	if err != nil {
		return 0, err
	}

	var count int
	if err := cursor.One(&count); err != nil {
		if err == r.ErrEmptyResult {
			return 0, api.ErrNoResult
		}
		return 0, api.NewDatabaseError(d, err, "")
	}

	return count, nil
}

// WhereAndFetchOne returns one document that match the filter
func (d *Default) WhereAndFetchOne(filter interface{}, result interface{}) error {
	cursor, err := r.Table(d.table).Filter(filter).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.One(result); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereAndFetchLimit returns paginated list of document
func (d *Default) WhereAndFetchLimit(filter interface{}, paginator *api.Pagination, results interface{}) error {
	cursor, err := r.Table(d.table).Filter(filter).Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	if err := cursor.All(results); err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Update a document that match the selector
func (d *Default) Update(selector interface{}, data interface{}) error {
	_, err := r.Table(d.table).Filter(selector).Update(data).RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// UpdateID updates a document using his id
func (d *Default) UpdateID(id interface{}, data interface{}) error {
	_, err := r.Table(d.table).Get(id).Update(data).RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// DeleteAll documents from the database
func (d *Default) DeleteAll(pred interface{}) error {
	_, err := r.Table(d.table).Filter(pred).Delete().RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Delete a document from the database
func (d *Default) Delete(id interface{}) error {
	_, err := r.Table(d.table).Get(id).Delete().RunWrite(d.session)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// List all entities from the database
func (d *Default) List(results interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	return d.Search(results, nil, sortParams, pagination)
}

// Search all entities in the database
func (d *Default) Search(results interface{}, filter interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	term := r.Table(d.table)

	// Filter
	if filter != nil {
		term = term.Filter(filter)
	}

	// Sort
	if sortParams != nil {
		term = term.OrderBy(ConvertSortParameters(*sortParams)...)
	}

	// Slice result
	if pagination != nil {
		term = term.Slice(pagination.Offset(), pagination.Offset()+pagination.PerPage)
	}

	// Run the query
	cursor, err := term.Run(d.session)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	// Fetch cursor
	err = cursor.All(results)
	if err != nil {
		if err == r.ErrEmptyResult {
			return api.ErrNoResult
		}
		return err
	}

	return nil
}
