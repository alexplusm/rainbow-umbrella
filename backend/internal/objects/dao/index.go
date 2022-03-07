package dao

import (
	"fmt"
	"time"
)

func timeToDAO(t time.Time) string {
	return t.Format(time.RFC3339)
}

func timeFromDAO(t string) (time.Time, error) {
	value, err := time.Parse("2006-01-02 15:04:05", t)
	if err != nil {
		return time.Time{}, fmt.Errorf("[timeFromDAO][1]: %+v", err)
	}

	return value, nil
}
