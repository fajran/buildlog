package buildlog

import (
	"fmt"
	"log"

	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

type migrationLogger struct {
}

func (m migrationLogger) Verbose() bool {
	return false
}

func (m migrationLogger) Printf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	log.Printf("[db migration] %s", s)
}

func (bl *BuildLog) MigrateDb() error {
	driver, err := postgres.WithInstance(bl.db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance("file://migrations", "postgres", driver)
	if err != nil {
		return err
	}

	m.Log = migrationLogger{}

	err = m.Up()
	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}
