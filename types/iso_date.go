package types

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type IsoDate time.Time

func ParseIsoDate(value string) (IsoDate, error) {
	parsedDate, err := time.Parse(time.DateOnly, value)
	if err != nil {
		return IsoDate{}, err
	}

	return IsoDate(parsedDate), nil
}

func (t IsoDate) Value() (driver.Value, error) {
	return time.Time(t), nil
}

func (t IsoDate) String() string {
	return time.Time(t).Format(time.DateOnly)
}

func (t IsoDate) GoString() string {
	return time.Time(t).GoString()
}

func (t IsoDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Format(time.DateOnly))
}

func (t *IsoDate) UnmarshalJSON(data []byte) error {
	var ptr *string
	if err := json.Unmarshal(data, &ptr); err != nil {
		return err
	}

	if ptr == nil {
		return nil
	}

	parsedDate, err := time.Parse(time.DateOnly, *ptr)
	if err != nil {
		return err
	}

	*t = IsoDate(parsedDate)
	return nil
}
