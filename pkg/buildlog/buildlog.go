package buildlog

import (
	"database/sql"
	"io"
	"strings"

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

	Type string

	ContentType          string
	ContentTypeParameter string

	Size int64

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

func splitContentTypeParameter(contentType string) (string, string) {
	p := strings.SplitN(contentType, ";", 2)

	ct := strings.TrimSpace(p[0])
	ctp := ""
	if len(p) > 1 {
		ctp = strings.TrimSpace(p[1])
	}

	return ct, ctp
}

func (b *Build) Log(logType, contentType string, content io.Reader) (int, error) {
	ct, ctp := splitContentTypeParameter(contentType)

	var id int
	err := b.buildlog.db.QueryRow(
		`INSERT INTO logs (build_id, type, content_type, content_type_parameter) VALUES ($1, $2, $3, $4) RETURNING id`,
		b.Id, logType, ct, ctp).
		Scan(&id)
	if err != nil {
		return 0, err
	}

	size, err := b.buildlog.storage.Store(id, content)
	if err != nil {
		return id, err
	}

	_, err = b.buildlog.db.Exec(`UPDATE logs SET size=$1 WHERE id=$2`, size, id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (b *Build) GetLog(id int) (*Log, error) {
	rows, err := b.buildlog.db.Query(
		`SELECT id, type, content_type, content_type_parameter, size FROM logs
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
		err = rows.Scan(&data.Id, &data.Type, &data.ContentType, &data.ContentTypeParameter, &data.Size)
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

func (l *Log) Read() (io.Reader, error) {
	return l.Build.buildlog.storage.Read(l.Id)
}
