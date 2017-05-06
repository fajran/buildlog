package buildlog

import (
	"database/sql"
)

type BuildLog struct {
	db *sql.DB
}

type Build struct {
	Id  int
	Key string

	buildlog *BuildLog
}

type Log struct {
	Id int

	build *Build
}

func NewBuildLog(db *sql.DB) *BuildLog {
	return &BuildLog{
		db: db,
	}
}

func (bl *BuildLog) Create(key string) (int, error) {
	var id int
	err := bl.db.QueryRow(`INSERT INTO builds (key) VALUES ($1) RETURNING id`, key).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (bl *BuildLog) Get(id int) (*Build, error) {
	rows, err := bl.db.Query(`SELECT id, key FROM builds WHERE id=$1`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		build := Build{
			buildlog: bl,
		}
		rows.Scan(&build.Id, &build.Key)
		return &build, nil
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return nil, nil
}

func (b *Build) Log(logType, contentType string) (int, error) {
	var id int
	err := b.buildlog.db.QueryRow(
		`INSERT INTO logs (build_id, type, content_type) VALUES ($1, $2, $3) RETURNING id`,
		b.Id, logType, contentType).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
