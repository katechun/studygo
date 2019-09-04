package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var (
		id   int
		name string
	)

	rows, err := db.Query("select id,name from users where id=?", 1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
