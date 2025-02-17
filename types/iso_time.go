package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type IsoTime time.Time

func ParseIsoTime(value string) (IsoTime, error) {
	parsedTime, err := time.Parse(time.TimeOnly, value)
	if err != nil {
		return IsoTime{}, err
	}

	return IsoTime(parsedTime), nil
}

func (t IsoTime) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t IsoTime) String() string {
	return time.Time(t).Format(time.TimeOnly)
}

func (t IsoTime) GoString() string {
	return time.Time(t).GoString()
}

func (t IsoTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.TimeOnly))
}

func (t *IsoTime) UnmarshalJSON(data []byte) error {
	var ptr *string
	if err := json.Unmarshal(data, &ptr); err != nil {
		return err
	}

	if ptr == nil {
		return nil
	}

	parsedTime, err := time.Parse(time.TimeOnly, *ptr)
	if err != nil {
		return err
	}

	*t = IsoTime(parsedTime)
	return nil
}
