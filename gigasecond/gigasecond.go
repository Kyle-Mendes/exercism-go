package gigasecond

import (
	"time"
)

const (
	TestVersion = 3
	Giga        = time.Second * 1e9
)

func AddGigasecond(date time.Time) time.Time {
	return date.Add(Giga)
}
