package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

const db_url = "postgres://postgres:postgres@localhost:5432/postgres"

func main() {
	conn, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// err = conn.QueryRow(context.Background(), "SELECT name FROM users WHERE name IS NULL").Scan(&name)
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	// 	os.Exit(1)
	// }

	rows, err := conn.Query(context.Background(), "SELECT name FROM users")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}
	for rows.Next() {
		var name *string
		err := rows.Scan(&name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("name: %v", name)
	}

}
