package calendar

import (
	"time"
)

// Calendar represents a month-based calendar calculator.
type Calendar struct {
	target    time.Time
	startDay  time.Weekday
	workdays  map[time.Weekday]struct{}
	threshold int
}

// New creates a new Calendar instance targeting the given time's month and year.
func New(t time.Time) *Calendar {
	// Default to standard industry practices
	c := &Calendar{
		target:    t,
		startDay:  time.Sunday,
		threshold: 3,
		workdays: map[time.Weekday]struct{}{
			time.Monday:    {},
			time.Tuesday:   {},
			time.Wednesday: {},
			time.Thursday:  {},
			time.Friday:    {},
		},
	}
	return c
}

// WithStartDay sets the first day of the week (e.g., Sunday or Monday).
func (c *Calendar) WithStartDay(d time.Weekday) *Calendar {
	c.startDay = d
	return c
}

// WithWorkdays sets which days are considered workdays.
func (c *Calendar) WithWorkdays(days ...time.Weekday) *Calendar {
	c.workdays = make(map[time.Weekday]struct{})
	for _, d := range days {
		c.workdays[d] = struct{}{}
	}
	return c
}

// WithThreshold sets the minimum number of workdays required for a week
// to belong to the target month (the "N workdays" rule).
func (c *Calendar) WithThreshold(n int) *Calendar {
	c.threshold = n
	return c
}

func (c *Calendar) isWorkday(d time.Weekday) bool {
	_, ok := c.workdays[d]
	return ok
}

// firstOfMonth returns the 1st day of the targeted month.
func (c *Calendar) firstOfMonth() time.Time {
	return time.Date(c.target.Year(), c.target.Month(), 1, 0, 0, 0, 0, c.target.Location())
}
