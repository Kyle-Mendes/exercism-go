package clock

import "fmt"
import "math"

const TestVersion = 2

type Clock struct {
	Hours   int
	Minutes int
}

func New(hours, minutes int) Clock {
	h := hours
	m := minutes

	c := Clock{h, m}
	c = c.simplify()

	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.Hours, c.Minutes)
}

func (c Clock) Add(a int) Clock {
	c.Minutes += a

	return c.simplify()
}

func (c Clock) simplify() Clock {

	// If the minutes are over an hour, reduce minutes under 59 and increment hours accordingly.
	if c.Minutes > 59 {
		floor := int(math.Floor(float64(c.Minutes / 60)))

		c.Hours += floor
		c.Minutes -= floor * 60
	}

	if c.Hours > 23 {
		c.Hours -= 24
	}

	if c.Minutes < 0 {
		floor := int(math.Floor(float64(c.Minutes / 60)))

		c.Hours += (floor - 1)
		c.Minutes -= (floor - 1) * 60
	}

	if c.Hours < 0 {
		floor := int(math.Floor(float64(c.Hours / -24)))
		h := c.Hours

		c.Hours = (floor+1)*24 + h
	}

	return c
}
