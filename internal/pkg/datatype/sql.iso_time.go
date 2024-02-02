package datatype

import (
	"database/sql/driver"
	"reflect"
	"strconv"
	"time"
)

type SqlIsoTime string

func (d *SqlIsoTime) Value() (driver.Value, error) {
	return d, nil
}

func (d *SqlIsoTime) Scan(value interface{}) error {
	*d = ""

	if reflect.ValueOf(value).IsZero() {
		return nil
	}

	if reflect.ValueOf(value).Kind() == reflect.String {
		t, err := strconv.Atoi(reflect.ValueOf(value).String())
		if err != nil {
			return err
		}
		now := time.UnixMilli(int64(t))
		*d = SqlIsoTime(now.Format(time.RFC3339))
	}
	return nil
}
