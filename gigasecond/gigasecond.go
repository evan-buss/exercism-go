package gigasecond

import (
	"time"
)

// AddGigasecond returns the time t plus a gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
