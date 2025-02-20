package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
)

var db *sql.DB

func initDB() {
	var err error
	connStr := "postgres://user:password@db:5432/mydb?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to DB: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping DB: %v", err))
	}

	fmt.Println("Connected to PostgreSQL successfully!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := db.Ping()
	if err != nil {
		http.Error(w, fmt.Sprintf("Database ping failed: %v", err), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Database is alive!")
}

func main() {
	initDB()
	http.HandleFunc("/", handler)
	port := "1337"
	fmt.Println("Server running on port:", port)
	http.ListenAndServe(":"+port, nil)
}
