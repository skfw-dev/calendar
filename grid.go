package calendar

import "time"

// Grid represents a 2D slice of time.Time, typically weeks and days.
type Grid [][]time.Time

// Grid generates the full calendar grid for the month, including padding days
// from the previous and next months to complete the rows.
func (c *Calendar) Grid() Grid {
	first := c.firstOfMonth()

	// Determine start offset based on the configured start day
	// (c.startDay - first.Weekday()) mod 7 logic:
	startOffset := int(first.Weekday() - c.startDay)
	if startOffset < 0 {
		startOffset += 7
	}

	curr := first.AddDate(0, 0, -startOffset)
	var grid Grid

	for {
		week := make([]time.Time, 7)
		for i := range 7 {
			week[i] = curr
			curr = curr.AddDate(0, 0, 1)
		}
		grid = append(grid, week)

		// Stop if we've passed the month and finished a full week
		if curr.Month() != c.target.Month() && curr.After(first) {
			break
		}
	}

	return grid
}
