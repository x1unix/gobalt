package dateutil

import (
	"encoding"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var (
	_ json.Marshaler           = (*DurationString)(nil)
	_ json.Unmarshaler         = (*DurationString)(nil)
	_ encoding.TextMarshaler   = (*DurationString)(nil)
	_ encoding.TextUnmarshaler = (*DurationString)(nil)
)

// DurationString type is wrapper around time.Duration that
// implements marshaling and unmarshal of duration from string.
//
// See: time.Duration.String() method.
type DurationString struct {
	time.Duration
}

// NewDurationString wraps time.Duration with DurationString type.
func NewDurationString(d time.Duration) DurationString {
	return DurationString{
		Duration: d,
	}
}

func (d *DurationString) UnmarshalText(src []byte) error {
	duration, err := time.ParseDuration(string(src))
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}

func (d DurationString) MarshalText() ([]byte, error) {
	return []byte(d.Duration.String()), nil
}

func (d DurationString) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Duration.String())
}

func (d *DurationString) UnmarshalJSON(src []byte) error {
	unquoted, err := strconv.Unquote(string(src))
	if err != nil {
		return fmt.Errorf("value '%s' a valid JSON string (%w)", src, err)
	}

	duration, err := time.ParseDuration(unquoted)
	if err != nil {
		return err
	}

	d.Duration = duration
	return nil
}
