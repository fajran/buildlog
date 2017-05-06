package storage

import (
	"io"
)

type Storage interface {
	Info() Info
	Store(id int, content io.Reader) (StoreInfo, error)
}

type Info struct {
	Name string
}

type StoreInfo struct {
	Id   string
	Size int64
}
