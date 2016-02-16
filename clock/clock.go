package clock

import "fmt"

const TestVersion = 2

type Clock struct {
	minutes int
}

func New(h, m int) Clock {
	minutes := (h*60 + m) % 1440

	if minutes < 0 {
		minutes += 1440
	}

	return Clock{minutes}
}

func (c Clock) Add(m int) Clock {
	c.minutes = (c.minutes + m) % 1440

	if c.minutes < 0 {
		c.minutes += 1440
	}

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}
