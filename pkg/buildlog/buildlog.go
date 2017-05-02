package buildlog

import (
	"time"
)

type BuildLog struct {
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

func NewBuildLog() *BuildLog {
	return &BuildLog{
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
