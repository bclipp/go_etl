package main

import (
	"fmt"
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

	cr.SendQuery(ETL1())
	cr.SendQuery(ETL2())

}