package mongodb

import (
	"zenithar.org/go/common/dao/api"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Default contains the basic implementation of the MongoCRUD interface
type Default struct {
	table   string
	db      string
	session *mgo.Session
}

// NewCRUDTable sets up a new Default struct
func NewCRUDTable(session *mgo.Session, db, table string) api.EntityCRUD {
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

// GetTable returns the table as a mgo.Collection
func (d *Default) GetTable() interface{} {
	return d.session.DB(d.db).C(d.table)
}

// GetSession returns the current session
func (d *Default) GetSession() interface{} {
	return d.session
}

// Insert inserts a document into the database
func (d *Default) Insert(data interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Insert(data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// InsertOrUpdate inserts or update document if exists
func (d *Default) InsertOrUpdate(id interface{}, data interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	_, err := session.DB(d.GetDBName()).C(d.GetTableName()).UpsertId(id, data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Update performs an update on an existing resource according to passed data
func (d *Default) Update(selector interface{}, data interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Update(selector, data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// UpdateID performs an update on an existing resource with ID that equals the id argument
func (d *Default) UpdateID(id interface{}, data interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).UpdateId(id.(string), data)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// DeleteAll deletes resources that match the passed filter
func (d *Default) DeleteAll(pred interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	_, err := session.DB(d.GetDBName()).C(d.GetTableName()).RemoveAll(pred)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Delete deletes a resource with specified ID
func (d *Default) Delete(id interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).RemoveId(id.(string))
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// Find searches for a resource in the database and then returns a cursor
func (d *Default) Find(id interface{}, value interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).FindId(id).One(value)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindFetchOne searches for a resource and then unmarshals the first row into value
func (d *Default) FindFetchOne(id string, value interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(bson.M{"_id": id}).One(value)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindOneBy is an utility for fetching values if they are stored in a key-value manenr.
func (d *Default) FindOneBy(key string, value interface{}, result interface{}) error {
	filterMap := bson.M{
		key: value,
	}
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filterMap).One(result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindBy is an utility for fetching values if they are stored in a key-value manenr.
func (d *Default) FindBy(key string, value interface{}, results interface{}) error {
	filterMap := bson.M{
		key: value,
	}
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filterMap).All(results)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// FindByAndCount returns the number of elements that match the filter
func (d *Default) FindByAndCount(key string, value interface{}) (int, error) {
	filterMap := bson.M{
		key: value,
	}
	session := d.session.Clone()
	defer session.Close()

	n, err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filterMap).Count()
	if err != nil {
		if err == mgo.ErrNotFound {
			return 0, api.ErrNoResult
		}
		return 0, api.NewDatabaseError(d, err, "")
	}

	return n, nil
}

// FindByAndFetch retrieves a value by key and then fills results with the result.
func (d *Default) FindByAndFetch(key string, value interface{}, results interface{}) error {
	filterMap := bson.M{
		key: value,
	}
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filterMap).All(results)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereCount allows counting with multiple fields
func (d *Default) WhereCount(filter interface{}) (int, error) {
	session := d.session.Clone()
	defer session.Close()

	count, err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter).Count()
	if err != nil {
		if err == mgo.ErrNotFound {
			return 0, api.ErrNoResult
		}
		return 0, api.NewDatabaseError(d, err, "")
	}

	return count, nil
}

// Where allows filtering with multiple fields
func (d *Default) Where(filter interface{}, results interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter).All(results)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereAndFetch filters with multiple fields and then fills results with all found resources
func (d *Default) WhereAndFetch(filter interface{}, results interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter).All(results)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereAndFetchLimit filters with multiple fields and then fills results with all found resources
func (d *Default) WhereAndFetchLimit(filter interface{}, paginator *api.Pagination, results interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter).Limit(int(paginator.PerPage)).Skip(int(paginator.Offset())).All(results)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// WhereAndFetchOne filters with multiple fields and then fills result with the first found resource
func (d *Default) WhereAndFetchOne(filter interface{}, result interface{}) error {
	session := d.session.Clone()
	defer session.Close()

	err := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter).One(result)
	if err != nil {
		if err == mgo.ErrNotFound {
			return api.ErrNoResult
		}
		return api.NewDatabaseError(d, err, "")
	}

	return nil
}

// List all entities from the database
func (d *Default) List(results interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	return d.Search(results, bson.M{}, sortParams, pagination)
}

// Search all entities from the database
func (d *Default) Search(results interface{}, filter interface{}, sortParams *api.SortParameters, pagination *api.Pagination) error {
	session := d.session.Clone()
	defer session.Close()

	// Apply Filter
	if filter == nil {
		filter = bson.M{}
	}

	// Get total
	if pagination != nil {
		total, err := d.WhereCount(filter)
		if err != nil {
			return api.NewDatabaseError(d, err, "")
		}
		pagination.SetTotal(uint(total))
	}

	// Prepare the query
	query := session.DB(d.GetDBName()).C(d.GetTableName()).Find(filter)

	// Apply sorts
	if sortParams != nil {
		sort := ConvertSortParameters(*sortParams)
		if len(sort) > 0 {
			query = query.Sort(sort...)
		}
	}

	// Paginate
	if pagination != nil {
		query = query.Limit(int(pagination.PerPage)).Skip(int(pagination.Offset()))
	}

	err := query.All(results)
	if err != nil {
		return api.NewDatabaseError(d, err, "")
	}

	return err
}
