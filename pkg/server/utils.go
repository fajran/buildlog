package server

import (
	"fmt"
	"time"
)

type iso8601 time.Time

func (t iso8601) MarshalJSON() ([]byte, error) {
	s := fmt.Sprintf("\"%s\"", time.Time(t).UTC().Format(time.RFC3339))
	return []byte(s), nil
}
