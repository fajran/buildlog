package disk

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/fajran/buildlog/pkg/storage"
)

type DiskStorage struct {
	DataPath string
}

func ensureDirectory(p string) error {
	info, err := os.Stat(p)
	if os.IsNotExist(err) {
		return os.MkdirAll(p, 0755)
	}
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	return fmt.Errorf("Path is not a directory: %s", p)
}

func NewDiskStorage(dataPath string) (*DiskStorage, error) {
	err := ensureDirectory(dataPath)
	if err != nil {
		return nil, err
	}

	return &DiskStorage{
		DataPath: dataPath,
	}, nil
}

func (s *DiskStorage) Info() storage.Info {
	return storage.Info{
		Name: "disk",
	}
}

func (s *DiskStorage) Store(id int, content io.Reader) (storage.StoreInfo, error) {
	sid := fmt.Sprintf("%d", id)
	p := path.Join(s.DataPath, sid)
	f, err := os.Create(p)
	if err != nil {
		return storage.StoreInfo{}, err
	}

	size, err := io.Copy(f, content)
	if err != nil {
		return storage.StoreInfo{}, err
	}

	return storage.StoreInfo{
		Id:   sid,
		Size: size,
	}, nil
}

func (s *DiskStorage) Read(id int) (io.Reader, error) {
	sid := fmt.Sprintf("%d", id)
	p := path.Join(s.DataPath, sid)

	return os.Open(p)
}
