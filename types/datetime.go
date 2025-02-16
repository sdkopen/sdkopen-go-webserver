package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// DateTime for empty date/time field
type DateTime time.Time

// ParseDateTime converts string to date time
func ParseDateTime(value string) (DateTime, error) {
	parsedDate, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return DateTime{}, err
	}

	return DateTime(parsedDate), nil
}

// Value converts iso date time to sql driver value
func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

// String returns the iso date time formatted using the format string
func (t DateTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}

// GoString returns the iso date time in Go source code format string
func (t DateTime) GoString() string {
	return time.Time(t).GoString()
}

// MarshalJSON converts iso date time to json string format
func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

// UnmarshalJSON converts json string to iso date time
func (t *DateTime) UnmarshalJSON(data []byte) error {
	var ptr *string
	if err := json.Unmarshal(data, &ptr); err != nil {
		return err
	}

	if ptr == nil {
		return nil
	}

	parsedDateTime, err := time.Parse(time.RFC3339, *ptr)
	if err != nil {
		return err
	}

	*t = DateTime(parsedDateTime)
	return nil
}
