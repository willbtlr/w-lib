package marshallable

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Duration struct {
	time.Duration
}

func (d *Duration) SetTaggedDefaults(tag string) (err error) {
	d.Duration, err = time.ParseDuration(tag)
	return
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	duration, err := time.ParseDuration(str)
	if err != nil {
		return fmt.Errorf("failed to parse duration %q: %v", str, err)
	}
	d.Duration = duration
	return nil
}

func (d *Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(d.String())), nil
}

func (d *Duration) String() string {
	return d.Duration.String()
}
