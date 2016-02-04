package api

// EntityTable contains the most basic table functions
type EntityTable interface {
	GetTableName() string
	GetDBName() string
	GetTable() interface{}
	GetSession() interface{}
}

// EntityCreator contains a function to create new instances in the table
type EntityCreator interface {
	Insert(data interface{}) error
	InsertOrUpdate(id interface{}, data interface{}) error
}

// EntityReader allows fetching resources from the database
type EntityReader interface {
	Find(id interface{}, result interface{}) error
	FindOneBy(key string, value interface{}, result interface{}) error
	FindBy(key string, value interface{}, results interface{}) error
	FindByAndCount(key string, value interface{}) (int, error)

	Where(filter map[string]interface{}, results interface{}) error
	WhereCount(filter map[string]interface{}) (int, error)
	WhereAndFetchOne(filter map[string]interface{}, result interface{}) error
	WhereAndFetchLimit(filter map[string]interface{}, paginator *Pagination, results interface{}) error

	List(results interface{}, sortParams *SortParameters, pagination *Pagination) error
	Search(results interface{}, filter map[string]interface{}, sortParams *SortParameters, pagination *Pagination) error
}

// EntityUpdater allows updating existing resources in the database
type EntityUpdater interface {
	Update(selector interface{}, data interface{}) error
	UpdateID(id interface{}, data interface{}) error
}

// EntityDeleter allows deleting resources from the database
type EntityDeleter interface {
	Delete(id interface{}) error
	DeleteAll(pred interface{}) error
}

// EntityCRUD is the interface that every table should implement
type EntityCRUD interface {
	EntityCreator
	EntityReader
	EntityUpdater
	EntityDeleter
	EntityTable
}
