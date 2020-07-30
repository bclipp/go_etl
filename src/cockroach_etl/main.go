package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	var cr = CockroachDB{
		IPAddress:  "localhost",
		PostgresDB: "test",
	}
	err := cr.Connect();if err != nil {
		fmt.Print(err.Error())
	}
	defer cr.Disconnect()

}