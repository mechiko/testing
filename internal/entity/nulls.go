package entity

import (
	// "alcomerge/core"
	"database/sql/driver"
	"errors"
)

// var app = core.GetInstance()

type NullString string
type NullInt64 int64

func (s *NullString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	strVal, ok := value.(string)
	if !ok {
		return errors.New("column is not a string")
	}
	*s = NullString(strVal)
	return nil
}
func (s NullString) Value() (driver.Value, error) {
	// if len(s) == 0 { // if nil or empty string
	// 	return nil, nil
	// }
	return string(s), nil
}

func (i *NullInt64) Scan(value interface{}) error {
	// s := fmt.Sprintf("%T %v", value, value)
	// app.ZeroLogger().Debug().Str("Scan value", s).Send()
	if value == nil {
		*i = 0
		return nil
	}
	i64Val, ok := value.(int64)
	if !ok {
		return errors.New("column is not a int64")
	}
	*i = NullInt64(i64Val)
	return nil
}

func (i NullInt64) Value() (driver.Value, error) {
	// s := fmt.Sprintf("%T %v", i, i)
	// app.ZeroLogger().Debug().Str("Value()", s).Send()
	return int64(i), nil
}
