// Copyright (c) 2021 Alexandre Trottier (bonedaddy). All rights reserved.
// Use of this source code is governed by the Apache 2.0 License that can be found in
// the LICENSE file.

package database

import (
	"database/sql"
	"database/sql/driver"
	"time"
)

// Date provides a time.Time wrapper for databases which may not support time formats
type Date time.Time

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	*date = Date(nullTime.Time)
	return
}

func (date Date) Value() (driver.Value, error) {
	y, m, d := time.Time(date).Date()
	h := time.Time(date).Hour()
	min := time.Time(date).Minute()
	s := time.Time(date).Second()
	return time.Date(y, m, d, h, min, s, 0, time.Time(date).Location()), nil
}

// GormDataType gorm common data type
func (date Date) GormDataType() string {
	return "date"
}
