package storage

import (
	"io"
)

type Storage interface {
	Info() Info
	Store(id int, content io.Reader) (int64, error)
	Read(id int) (io.Reader, error)
}

type Info struct {
	Name string
}
