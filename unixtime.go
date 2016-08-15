package utils

import (
    "time"
    "strconv"
)

// UnixTime - Time in seconds (UNIX Timestamp)
type UnixTime int64

// Time (UnixTime) - UnixTime to time.Time
func (t UnixTime) Time() time.Time {
    return time.Unix(int64(t), 0)
}

// String (UnixTime) - UnixTime to String
func (t UnixTime) String() string {
    return strconv.FormatInt(int64(t), 10)
}

// AddDate (UnixTime) - Add Date
func (t UnixTime) AddDate(years, months, days int) UnixTime {
    return UnixTime(t.Time().AddDate(years, months, days).Unix())
}

// Add (UnixTime) - Add Date
func (t UnixTime) Add(d time.Duration) UnixTime {
    return UnixTime(t.Time().Add(d).Unix())
}

// NowUnixTime - Create UnixTime from Now time 
func NowUnixTime() UnixTime {
    return UnixTime(time.Now().Unix())
}