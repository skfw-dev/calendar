package calendar

import (
	"testing"
	"time"
)

func TestManagedWeeks(t *testing.T) {
	// April 2026: April 1st is Wednesday.
	// Week starts Sunday March 29.
	// Workdays: Mar 30(M), Mar 31(T), Apr 1(W), Apr 2(Th), Apr 3(F).
	// CurrentMonth workdays: 3 (Apr 1, 2, 3)
	// If threshold is 3, this week belongs to April.

	target := time.Date(2026, time.April, 1, 0, 0, 0, 0, time.UTC)

	t.Run("Default Threshold 3", func(t *testing.T) {
		cal := New(target)
		weeks := cal.ManagedWeeks()

		// Expected first week should start on March 29
		if !weeks[0][0].Equal(time.Date(2026, time.March, 29, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("Expected first week to start on March 29, got %s", weeks[0][0])
		}
	})

	t.Run("Custom Threshold 4", func(t *testing.T) {
		// With threshold 4, March 29 week has only 3 April workdays, so it should NOT belong to April.
		cal := New(target).WithThreshold(4)
		weeks := cal.ManagedWeeks()

		// First week should now start on April 5
		if !weeks[0][0].Equal(time.Date(2026, time.April, 5, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("Expected first week to start on April 5 with threshold 4, got %s", weeks[0][0])
		}
	})

	t.Run("Custom Workdays (6-day week)", func(t *testing.T) {
		// If Saturday is also a workday, workdays are: Mar 30, 31, Apr 1, 2, 3, 4.
		// April workdays = 4.
		// Even with threshold 4, it should now belong to April.
		cal := New(target).
			WithThreshold(4).
			WithWorkdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday, time.Saturday)

		weeks := cal.ManagedWeeks()
		if !weeks[0][0].Equal(time.Date(2026, time.March, 29, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("Expected first week to start on March 29 with 6 workdays and threshold 4, got %s", weeks[0][0])
		}
	})

	t.Run("StartDay Shift Impact", func(t *testing.T) {
		// Target: December 2025 (Dec 1 is Monday)
		// Assume threshold 3, Mon-Fri workdays.
		// Week 1 (Sun Start): Nov 30 (S), Dec 1 (M), Dec 2 (T), Dec 3 (W), Dec 4 (Th), Dec 5 (F), Dec 6 (S)
		//   Dec workdays: 5 (M, T, W, Th, F). Result: Belongs to Dec.
		// Week 1 (Mon Start): Dec 1 (M), Dec 2 (T), Dec 3 (W), Dec 4 (Th), Dec 5 (F), Dec 6 (S), Dec 7 (S)
		//   Dec workdays: 5 (M, T, W, Th, F). Result: Belongs to Dec.

		targetDec := time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC)

		sunCal := New(targetDec).WithStartDay(time.Sunday)
		monCal := New(targetDec).WithStartDay(time.Monday)

		sunWeeks := sunCal.ManagedWeeks()
		monWeeks := monCal.ManagedWeeks()

		if !sunWeeks[0][0].Equal(time.Date(2025, time.November, 30, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("Sun-start: Expected Nov 30, got %s", sunWeeks[0][0])
		}
		if !monWeeks[0][0].Equal(time.Date(2025, time.December, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("Mon-start: Expected Dec 1, got %s", monWeeks[0][0])
		}
	})
}
