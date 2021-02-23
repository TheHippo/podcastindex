package podcastindex

import (
	"fmt"
	"strconv"
	"time"
)

// Time is a crutch to get the timestamp parsed correctly
// there is for obvious reasons no information on the timezone
type Time time.Time

// MarshalJSON is used to convert the timestamp to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *Time) UnmarshalJSON(s []byte) (err error) {
	u, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(u, 0)
	return nil
}

// String returns t as a formatted string
func (t Time) String() string {
	return time.Time(t).UTC().String()
}

// Duration is a crutch to get time.Duration parsed from the API
// results
type Duration time.Duration

// MarshalJSON is used to convert the duration to JSON
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(time.Duration(d).Seconds()), 10)), nil
}

// UnmarshalJSON is used to convert the duration from JSON
func (d *Duration) UnmarshalJSON(s []byte) (err error) {
	p, err := time.ParseDuration(fmt.Sprintf("%ss", s))
	if err != nil {
		return err
	}
	*(*time.Duration)(d) = p
	return nil
}

// String returns d as a formated string
func (d Duration) String() string {
	return time.Duration(d).String()
}
