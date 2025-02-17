package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type DateTime time.Time

func ParseDateTime(value string) (DateTime, error) {
	parsedDate, err := time.Parse(time.RFC3339, value)
	if err != nil {
		return DateTime{}, err
	}

	return DateTime(parsedDate), nil
}

func (t DateTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t DateTime) String() string {
	return time.Time(t).Format(time.RFC3339)
}

func (t DateTime) GoString() string {
	return time.Time(t).GoString()
}

func (t DateTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.RFC3339))
}

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
