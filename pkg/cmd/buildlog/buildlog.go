package buildlog

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/fajran/buildlog/pkg/buildlog"
	"github.com/fajran/buildlog/pkg/server"
	"github.com/fajran/buildlog/pkg/storage/disk"
)

func openDb() (*sql.DB, error) {
	uri := os.Getenv("DB_URI")
	if uri == "" {
		log.Fatal("Please define DB_URI environment variable")
	}
	return sql.Open("postgres", uri)
}

func Run() {
	addr := os.Getenv("SERVER_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	dataPath := os.Getenv("DATA_PATH")
	if dataPath == "" {
		log.Fatal("Please define DATA_PATH environment variable")
	}

	db, err := openDb()
	if err != nil {
		log.Fatal(err)
	}

	storage, err := disk.NewDiskStorage(dataPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Using data path at %s", dataPath)

	bl := buildlog.NewBuildLog(db, storage)

	log.Printf("Migrating database")
	err = bl.MigrateDb()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Running server on %s", addr)
	s := server.NewServer(addr, bl)
	s.Start()
}
