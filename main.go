package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

var gloAPP Config
var DB *sql.DB

// se asigna en build.sh
var version string

func main() {

	if len(os.Args) > 1 && os.Args[1] == "-v" {
		fmt.Println(version)
		return
	}
	gloAPP.load()

	log.Printf("START carto-audit %v", version)

	u := &url.URL{
		Scheme: "postgresql",
		User:   url.UserPassword(gloAPP.DbUser, gloAPP.DbPassword),
		// Combine host and port here
		Host:     fmt.Sprintf("%s:%s", gloAPP.DbHost, gloAPP.DbPort),
		Path:     gloAPP.DbName,
		RawQuery: "sslmode=disable", // Recommended to add if not using SSL
	}

	var err error
	DB, err = sql.Open("postgres", u.String())
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer DB.Close()

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to the database successfully.")

	iniciarProceso()
}
