package storage

import (
	"io"
)

type Storage interface {
	Info() Info
	Store(id int, content io.Reader) (StoreInfo, error)
	Read(id int) (io.Reader, error)
}

type Info struct {
	Name string
}

type StoreInfo struct {
	Id   string
	Size int64
}
