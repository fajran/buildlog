package buildlog

import (
	"database/sql"
	"io"

	logstorage "github.com/fajran/buildlog/pkg/storage"
)

type BuildLog struct {
	db      *sql.DB
	storage logstorage.Storage
}

type Build struct {
	Id  int
	Key string

	buildlog *BuildLog
}

type Log struct {
	Id int

	Type        string
	ContentType string

	ContentId string
	Size      int64

	Build *Build
}

func NewBuildLog(db *sql.DB, storage logstorage.Storage) *BuildLog {
	return &BuildLog{
		db:      db,
		storage: storage,
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
		err = rows.Scan(&build.Id, &build.Key)
		if err != nil {
			return nil, err
		}
		return &build, nil
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return nil, nil
}

func (b *Build) Log(logType, contentType string, content io.Reader) (int, error) {
	var id int
	err := b.buildlog.db.QueryRow(
		`INSERT INTO logs (build_id, type, content_type) VALUES ($1, $2, $3) RETURNING id`,
		b.Id, logType, contentType).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	info, err := b.buildlog.storage.Store(id, content)
	if err != nil {
		return id, err
	}

	_, err = b.buildlog.db.Exec(`UPDATE logs SET identifier=$1, size=$2 WHERE id=$3`, info.Id, info.Size, id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (b *Build) GetLog(id int) (*Log, error) {
	rows, err := b.buildlog.db.Query(
		`SELECT id, type, content_type, identifier, size FROM logs
		 WHERE build_id=$1 AND id=$2`,
		b.Id, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	if rows.Next() {
		data := Log{
			Build: b,
		}
		err = rows.Scan(&data.Id, &data.Type, &data.ContentType, &data.ContentId, &data.Size)
		if err != nil {
			return nil, err
		}
		return &data, nil
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return nil, nil

}
