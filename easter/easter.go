package easter

import (
	"time"
)

func getGoldenNumber(year int) int {
	// To determine the golden number, divide the year by 19 and use
	// the remainder + 1. We use 19 because that is the length of a metonic cycle
	return (year % 19) + 1
}

// GetEasterDateForYear returns the Easter date using the Gauss algorithm
func GetEasterDateForYear(year int) time.Time {
	a := year % 19
	b := int(year / 100)
	c := year % 100
	d := (19*a + b - int(b/4) - int(((b - int((b+8)/25) + 1) / 3)) + 15) % 30
	e := (32 + 2*(b%4) + 2*(int(c/4)) - d - (c % 4)) % 7
	f := d + e - 7*(int((a+11*d+22*e)/451)) + 114
	month := int(f / 31)
	day := f%31 + 1
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
