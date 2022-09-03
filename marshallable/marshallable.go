package marshallable

import (
	"encoding/json"
	"errors"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) UnmarshalJSON(bs []byte) error {
	var i interface{}
	err := json.Unmarshal(bs, &i)
	if err != nil {
		return err
	}

	if s, ok := i.(string); ok {
		dur, err := time.ParseDuration(s)
		if err != nil {
			return err
		}

		d.Duration = dur
	}

	return errors.New("time.Duration field must be string")
}
