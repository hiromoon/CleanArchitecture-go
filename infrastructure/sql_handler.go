package infrastructure

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/hiromoon/CleanArchitecture/interfaces/databases"
)

// SQLHandler is ...
type SQLHandler struct {
	Conn *sql.DB
}

// NewSQLHandler is ...
func NewSQLHandler() databases.SQLHandler {
	conn, err := sql.Open("mysql", "root:@tcp(db:3306)/sample")
	if err != nil {
		panic(err.Error)
	}

	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

// Execute is ...
func (handler *SQLHandler) Execute(statement string, args ...interface{}) (databases.Result, error) {
	res := SQLResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, err
	}

	res.Result = result
	return res, nil
}

// Query is ...
func (handler *SQLHandler) Query(statement string, args ...interface{}) (databases.Row, error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SQLRow), err
	}
	row := new(SQLRow)
	row.Rows = rows
	return row, nil
}

// SQLResult is ...
type SQLResult struct {
	Result sql.Result
}

// LastInsertID is ...
func (r SQLResult) LastInsertID() (int64, error) {
	return r.Result.LastInsertId()
}

// RowsAffected is ...
func (r SQLResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

// SQLRow is ...
type SQLRow struct {
	Rows *sql.Rows
}

// Scan is ...
func (r SQLRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

// Next is ...
func (r SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close is ...
func (r SQLRow) Close() error {
	return r.Rows.Close()
}
