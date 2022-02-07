package timer

import (
	"github.com/curatorc/cngf/config"
	"time"
)

type Time time.Time

const (
	timeFormat = `2006-01-02 15:04:05`
)

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	chinaTimezone, _ := time.LoadLocation(config.GetString("app.timezone"))
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), chinaTimezone)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

// Add returns the time t+d.
func (t Time) Add(duration time.Duration) Time {
	return Time(time.Time(t).Add(duration))
}

// After reports whether the time instant t is after u.
func (t Time) After(u Time) bool {
	return time.Time(t).After(time.Time(u))
}

// Before reports whether the time instant t is before u.
func (t Time) Before(u Time) bool {
	return time.Time(t).Before(time.Time(u))
}

func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (t Time) IsZero() bool {
	return time.Time(t).IsZero()
}
