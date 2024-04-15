package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

func generateDiff(db string) string {
	return `
DROP TABLE foo;
`
}

func main() {
	file, err := os.Open("./schema.sql")
	if err != nil {
		log.Fatalln("error opening file", err)
	}

	schema, err := io.ReadAll(file)
	password := os.Getenv("POSTGRES_PASSWORD")

	db, err := sql.Open("postgres", fmt.Sprintf("password=%s", password))
	if err != nil {
		log.Fatalln("error opening db connection", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		log.Fatalln("error creating model database", err)
	}

	query, err := db.Query("SELECT * FROM pg_catalog.pg_tables;")
	if err != nil {
		log.Fatalln("error executing schema query", err)
	}

	var row string
	for query.Next() {
		if err = query.Scan(&row); err != nil {
			log.Fatalln("error scanning row", err)
		}
		fmt.Println("row", row)
	}

	databases := []string{"abc", "foo", "baz"}
	for _, db := range databases {
		fmt.Printf("starting diff for '%s' database", db)
		//diff := generateDiff(db)
	}

	fmt.Println("hello, world!")
}
