package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	parsedDate, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return time.Time{}
	}
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	parsedDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return parsedDate.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	parsedDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	hour := parsedDate.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	parsedDate, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil {
		return "error"
	}
	return "You have an appointment on " + parsedDate.Format("Monday, January 2, 2006, at 15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentDate := time.Now()
	anniversaryDate := time.Date(currentDate.Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversaryDate
}
