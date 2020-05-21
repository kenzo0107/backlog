package backlog

import (
	"bytes"
	"fmt"
	"time"
)

// JSONTime exists so that we can have a String method converting the date
type JSONTime string

// String converts the unix timestamp into a string
func (t JSONTime) String() string {
	tm := t.Time()
	return fmt.Sprintf("\"%s\"", tm.Format(time.RFC3339))
}

// Time returns a `time.Time` representation of this value.
func (t JSONTime) Time() time.Time {
	tt, _ := time.Parse(time.RFC3339, string(t))
	return tt
}

// UnmarshalJSON will unmarshal both string and int JSON values
func (t *JSONTime) UnmarshalJSON(buf []byte) error {
	s := bytes.Trim(buf, `"`)
	*t = JSONTime(string(s))
	return nil
}
