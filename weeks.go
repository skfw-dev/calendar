package calendar

import "time"

// ManagedWeeks returns only the weeks that belong to the target month
// based on the configured workday threshold.
func (c *Calendar) ManagedWeeks() Grid {
	first := c.firstOfMonth()

	// Start from the first week that could possibly belong to this month
	startOffset := int(first.Weekday() - c.startDay)
	if startOffset < 0 {
		startOffset += 7
	}
	curr := first.AddDate(0, 0, -startOffset)

	var managed Grid

	for {
		week := make([]time.Time, 7)
		workdaysInTargetMonth := 0
		workdaysInNextMonth := 0

		for i := range 7 {
			week[i] = curr
			if c.isWorkday(curr.Weekday()) {
				if curr.Month() == c.target.Month() {
					workdaysInTargetMonth++
				} else if curr.After(first) {
					workdaysInNextMonth++
				}
			}
			curr = curr.AddDate(0, 0, 1)
		}

		// RULE: If workdaysInNextMonth >= threshold, this month's managed weeks are over.
		if workdaysInNextMonth >= c.threshold {
			break
		}

		// RULE: If workdaysInTargetMonth >= threshold, this week belongs to the target month.
		// (This automatically handles the previous month's boundary as well)
		if workdaysInTargetMonth >= c.threshold {
			managed = append(managed, week)
		}

		// Safety break: if we are far past the target month
		if curr.Month() != c.target.Month() && curr.After(first.AddDate(0, 1, 7)) {
			break
		}
	}

	return managed
}
