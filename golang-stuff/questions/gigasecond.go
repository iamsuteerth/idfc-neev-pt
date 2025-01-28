package questions

import "time"

// What's the time after 1 gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
