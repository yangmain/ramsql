package main

import (
	"database/sql"
	"fmt"

	_ "github.com/proullon/ramsql/driver"
)

func main() {

	db, err := sql.Open("ramsql", "")
	if err != nil {
		fmt.Printf("sql.Open : Error : %s\n", err)
		return
	}
	defer db.Close()

	_, err = db.Exec("CREATE TABLE account (id INT, email TEXT)")
	if err != nil {
		fmt.Printf("sql.Exec: Error: %s\n", err)
		return
	}

	rows, err := db.Query("SELECT * FROM account WHERE email = '?'", "foo@bar.com")
	if err != nil {
		fmt.Printf("sql.Query error : %s\n", err)
		return
	}

	i := 0
	for rows.Next() {
	    i++
	}
	fmt.Printf("%d rows affected\n", i)

	err = db.Close()
	if err != nil {
		fmt.Printf("sql.Close : Error : %s\n", err)
		return
	}
}
