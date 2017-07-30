package clock

import "fmt"

const testVersion = 4

// Clock represents a time with hours and minutes
type Clock struct {
	hour   int
	minute int
}

// UpdateHour updates a clock by hours
func (c Clock) UpdateHour(i int) Clock {
	c.hour += i

	for {
		if c.hour < 0 {
			c.hour += 24
		} else if c.hour >= 24 {
			c.hour += -24
		} else {
			break
		}
	}

	return c
}

// UpdateMinute updates a clock by minutes
func (c Clock) UpdateMinute(i int) Clock {
	c.minute += i

	for {
		if c.minute < 0 {
			c.minute += 60
			c = c.UpdateHour(-1)
		} else if c.minute >= 60 {
			c.minute += -60
			c = c.UpdateHour(1)
		} else {
			break
		}
	}

	return c
}

// New creates a new clock with the given hour and minute
func New(hour, minute int) Clock {
	c := Clock{
		hour:   0,
		minute: 0,
	}

	c = c.UpdateHour(hour)
	c = c.UpdateMinute(minute)
	return c
}

// Return a string representation of this clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add increments the time by the minutes provided
func (c Clock) Add(minutes int) Clock {
	return c.UpdateMinute(minutes)
}
