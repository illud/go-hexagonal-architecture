package helpers

import (
	"time"
)

func AddDayToDate(days int) time.Time {
	curr_date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	oneDayLater := curr_date.AddDate(0, 0, 1)

	return oneDayLater
}
