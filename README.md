# Go Custom Calendar Module

[![Go Reference](https://pkg.go.dev/badge/github.com/skfw-dev/calendar.svg)](https://pkg.go.dev/github.com/skfw-dev/calendar)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

A robust, enterprise-grade Go module for sophisticated date calculations, focusing on month-based grids and managed weeks with industry-standard logic.

## Key Features

- **Fluent API Design**: Chained methods for intuitive configuration.
- **Industry-Standard Robustness**: Precise handling of month boundaries and workday thresholds.
- **Customizable Logic**: Configurable workdays, start of week, and "N-workday" inclusion rules.
- **Grid Generation**: Easily generate 2D slices for calendar UI components.
- **Managed Weeks**: Calculate weeks belonging to a specific month based on business rules.

## Installation

```bash
go get github.com/skfw-dev/calendar
```

## Quick Start

The module provides a `Calendar` struct with a fluent interface to set up your calculation rules.

```go
package main

import (
    "fmt"
    "time"
    "github.com/skfw-dev/calendar"
)

func main() {
    // Create a calendar for Feb 2026
    t := time.Date(2026, 2, 1, 0, 0, 0, 0, time.UTC)
    
    cal := calendar.New(t).
        WithStartDay(time.Monday).
        WithWorkdays(time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday).
        WithThreshold(3)

    // Get managed weeks for this month
    weeks := cal.ManagedWeeks()
    fmt.Printf("February 2026 has %d managed weeks\n", len(weeks))
}
```

## API Reference

### `New(t time.Time) *Calendar`

Initializes a new calendar targeting the month and year of the provided time. Defaults to Sunday start and Mon-Fri workdays.

### `WithStartDay(d time.Weekday) *Calendar`

Sets the first day of the week.

### `WithWorkdays(days ...time.Weekday) *Calendar`

Defines which days are considered workdays for threshold calculations.

### `WithThreshold(n int) *Calendar`

Sets the minimum number of workdays required in a week for it to be associated with the target month (the "Nrd workday rule").

### `Grid() [][]time.Time`

Returns a full calendar grid including padding days from adjacent months.

### `ManagedWeeks() [][]time.Time`

Returns only those weeks that belong to the target month based on the configured threshold logic.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.
