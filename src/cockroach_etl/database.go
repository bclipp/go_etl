package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Database interface {
	connect()  error
	Disconnect()
	sendQueryReturnData(sqlQuery string)(*sql.Rows, error)
	sendQuery(query string)  (sql.Result, error)
}

type CockroachDB struct {
	DB               *sql.DB
	IPAddress        string
	PostgresDB       string
}

func (cr *CockroachDB) Connect() error {
	psqlInfo := fmt.Sprintf("host=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cr.IPAddress, cr.PostgresUser, cr.PostgresPassword, cr.PostgresDB)

	db, err := sql.Open("postgres", psqlInfo);if err != nil {
		return err
	}
	cr.DB = db

	err = cr.DB.Ping();if err != nil {
		return err
	}

	return nil
}

func (cr *CockroachDB) Disconnect() {
	_ = cr.DB.Close()
}

func (cr *CockroachDB) sendQueryReturnData(sqlQuery string)(*sql.Rows, error) {
	rows, err := cr.DB.Query(sqlQuery)
	return rows, err
}

func (cr *CockroachDB) SendQuery(query string) (sql.Result, error) {
	result, err := cr.DB.Exec(query)
	if err != nil { return result, err}

	return result, nil
}
