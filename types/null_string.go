package types

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"reflect"
)

// NullString for empty string field
type NullString sql.NullString

// Scan converts sql driver value to null string
func (t *NullString) Scan(value interface{}) error {
	var i sql.NullString
	if err := i.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*t = NullString{i.String, false}
	} else {
		*t = NullString{i.String, true}
	}

	return nil
}

// Value converts null string to sql driver value
func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}

	return ns.String, nil
}

// MarshalJSON converts null string to json string format
func (t NullString) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(t.String)
}

// UnmarshalJSON converts json string to null string
func (t *NullString) UnmarshalJSON(data []byte) error {
	var ptr *string
	if err := json.Unmarshal(data, &ptr); err != nil {
		return err
	}

	if ptr != nil {
		t.Valid = true
		t.String = *ptr
	} else {
		t.Valid = false
	}

	return nil
}
