package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dsn := "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Ping failed:", err)
	}

	fmt.Println("Connected to DB")

	runQueryWithTimeout(db)
}

func runQueryWithTimeout(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("⏳ Running query with 2s timeout...")
	row := db.QueryRowContext(ctx, "SELECT pg_sleep(5)")

	var dummy string
	err := row.Scan(&dummy)

	if err != nil {
		fmt.Println("Query failed:", err)
	} else {
		fmt.Println("Query succeeded:", dummy)
	}
}
