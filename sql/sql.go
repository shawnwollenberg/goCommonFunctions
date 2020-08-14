package SQL

import (
	"database/sql"
	"fmt"
	"log"
)

func RunSQLScript(database, execSP, dbPassword, dbServer, dbUser string) {

	fullConnString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", dbServer, dbUser, dbPassword, database)

	conn, err := sql.Open("mssql", fullConnString)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	defer conn.Close()
	res, err := conn.Exec(execSP)
	if err != nil {
		log.Fatal("Open connection failed:", err.Error())
	}
	fmt.Println(res)
}
