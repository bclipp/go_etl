package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Database is used for Dep injection
type Database interface {
	connect()  error
	Disconnect()
	sendQueryReturnData(sqlQuery string)(*sql.Rows, error)
	sendQuery(query string)  (sql.Result, error)
}

// CockroachDB is used to store the variables for the database connection
type CockroachDB struct {
	DB               *sql.DB
	IPAddress        string
	Database       string
	Username        string
	Password       string
}

// Connect is used to handle connecting to the database
// Params:
// return:
//       error from the connection setup
func (cr *CockroachDB) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cr.IPAddress, cr.Username, cr.Password, cr.Database)

	db, err := sql.Open("postgres", psqlInfo);if err != nil {
		return err
	}
	cr.DB = db

	err = cr.DB.Ping();if err != nil {
		return err
	}

	return nil
}

// Close is used to handle closing the connection to the database
// Params:
// return:
//       error from the connection setup
func (cr *CockroachDB) Disconnect() {
	_ = cr.DB.Close()
}

// ReadTable is used for reading data from the database and storing it in the
// table field
// Params:
//       tableName: the table to query
//return:
//       Jason return document
//       rest http response code
//       the error
func (cr *CockroachDB) sendQueryReturnData(sqlQuery string)(*sql.Rows, error) {
	rows, err := cr.DB.Query(sqlQuery)
	return rows, err
}

// SendQuery is used for sending query to a database
// Params:
//       SendQuery: SQL to send
//return:
//		 result variable , see result interface doc in sql
//       the error
func (cr *CockroachDB) SendQuery(query string) (sql.Result, error) {
	result, err := cr.DB.Exec(query)
	if err != nil { return result, err}

	return result, nil
}
