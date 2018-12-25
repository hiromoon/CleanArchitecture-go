package databases

// SQLHandler is ...
type SQLHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

// Result is ...
type Result interface {
	LastInsertID() (int64, error)
	RowsAffected() (int64, error)
}

// Row is ...
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
