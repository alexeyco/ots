package ots

import (
	"strconv"
	"strings"
	"time"
)

type (
	// Status API status.
	Status int
	// State metadata state.
	State string
)

// Duration type.
type Duration time.Duration

// MarshalJSON marshal duration to json value.
func (d *Duration) MarshalJSON() ([]byte, error) {
	s := int64(d.Duration().Truncate(time.Second).Seconds())

	return strconv.AppendInt(nil, s, 10), nil
}

// UnmarshalJSON unmarshal duration from json value.
func (d *Duration) UnmarshalJSON(s []byte) (err error) {
	q, err := strconv.ParseInt(strings.Trim(string(s), `"`), 10, 64)
	if err != nil {
		return err
	}

	*(*time.Duration)(d) = time.Duration(q) * time.Second

	return
}

// Duration returns value as time.Duration.
func (d Duration) Duration() time.Duration {
	return time.Duration(d)
}

// Time type.
type Time time.Time

// MarshalJSON marshal time to json value.
func (t *Time) MarshalJSON() ([]byte, error) {
	return strconv.AppendInt(nil, t.Time().Unix(), 10), nil
}

// UnmarshalJSON unmarshal time from json value.
func (t *Time) UnmarshalJSON(s []byte) error {
	q, err := strconv.ParseInt(strings.Trim(string(s), `"`), 10, 64)
	if err != nil {
		return err
	}

	*(*time.Time)(t) = time.Unix(q, 0)

	return nil
}

// Time returns value as time.Time.
func (t Time) Time() time.Time {
	return time.Time(t)
}
