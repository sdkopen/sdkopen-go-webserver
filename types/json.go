package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JsonB map[string]interface{}

func (j *JsonB) Scan(value interface{}) error {
	result, valid := value.([]byte)
	if !valid {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(result, &j)
}

func (j JsonB) Value() (driver.Value, error) {
	return json.Marshal(j)
}
