package buildlog

import (
	"database/sql"
	"time"
)

type BuildLog struct {
	db *sql.DB
}

type Build struct {
	Id int32

	Key  string
	Name string

	Status   string
	Started  *time.Time
	Finished *time.Time
}

func NewBuildLog(db *sql.DB) *BuildLog {
	return &BuildLog{
		db: db,
	}
}

func (bl *BuildLog) Create(key string) *Build {
	var id int32
	bl.db.QueryRow(`INSERT INTO buildlog (key) VALUES ($1) RETURNING id`, key).Scan(id)

	return &Build{
		Id: id,
	}
}
