package SQL

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

func RunSQLScript(database, execSP, dbPassword, dbServer, dbUser string) {
	fullConnString := SQLdBString(dbServer, dbUser, dbPassword, database)
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

func SQLdBString(srvr, usr, pwd, db string) string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s", srvr, usr, pwd, db)
}
