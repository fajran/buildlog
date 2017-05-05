package buildlog

import (
	"database/sql"
	"time"
)

type BuildLog struct {
	db *sql.DB
}

type Build struct {
	Id int

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

func (bl *BuildLog) Create(key string) (*Build, error) {
	var id int
	err := bl.db.QueryRow(`INSERT INTO builds (key) VALUES ($1) RETURNING id`, key).Scan(&id)
	if err != nil {
		return nil, err
	}

	return &Build{
		Id: id,
	}, nil
}

func (bl *BuildLog) Get(id int) (*Build, error) {
	rows, err := bl.db.Query(`SELECT id, key, name, status, started, finished FROM builds WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		build := Build{}
		var started, finished string
		rows.Scan(&build.Id, &build.Key, &build.Name, &build.Status, &started, &finished)
		return &build, nil
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return nil, nil
}
