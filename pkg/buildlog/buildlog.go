package buildlog

import (
	"database/sql"
	"time"
)

type BuildLog struct {
	db *sql.DB

	counter int64
}

type Build struct {
	Id int64

	Key  string
	Name string

	Status   string
	Started  *time.Time
	Finished *time.Time
}

func NewBuildLog(db *sql.DB) *BuildLog {
	return &BuildLog{
		db:      db,
		counter: 0,
	}
}

func (bl *BuildLog) Create() *Build {
	bl.counter += 1
	id := bl.counter
	return &Build{
		Id: id,
	}
}
