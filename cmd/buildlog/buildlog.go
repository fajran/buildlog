package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/fajran/buildlog/pkg/buildlog"
	"github.com/fajran/buildlog/pkg/server"
)

func openDb() (*sql.DB, error) {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("Please define DB_URI environment variable")
	}
	return sql.Open("postgres", uri)
}

func main() {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	log.Printf("Running server on %s", addr)

	db, err := openDb()
	if err != nil {
		log.Fatal(err)
	}

	bl := buildlog.NewBuildLog(db)
	err = bl.MigrateDb()
	if err != nil {
		log.Fatal(err)
	}

	s := server.NewServer(addr, bl)
	s.Start()
}
