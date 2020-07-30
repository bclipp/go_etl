package main

import (
	"fmt"
)

func main() {
	// creating the instance of the database "class"
	var cr = CockroachDB{
		IPAddress:  "localhost",
		Database: "test",
		Username: "test",
		Password: "test",
	}
	// using the Connect "method"
	err := cr.Connect();if err != nil {
		fmt.Print(err.Error())
	}
	// at the end of the function close the connection using the "method"
	defer cr.Disconnect()
	// Send SQL using the "method" SendQuery
	cr.SendQuery(ETL1())
	cr.SendQuery(ETL2())

}