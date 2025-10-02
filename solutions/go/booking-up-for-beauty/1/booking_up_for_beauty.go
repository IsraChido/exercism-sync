package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/02/2006 15:04:05"
    
	t, _ := time.Parse(layout, date)

    return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"

    t, _ := time.Parse(layout, date)
    
	return !t.After(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)

    formattedHour := t.Format("15")
    
	if formattedHour >= "12" && formattedHour < "18" {
        return true
    } else {
        return false
    }
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
    t, _ := time.Parse(layout, date)

    formattedHour := t.Format("Monday, January 2, 2006, at 15:04.")

    return "You have an appointment on " + formattedHour
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().Year()

    t := time.Date(currentYear,time.September,15,0,0,0,0,time.UTC)

    return t
}
