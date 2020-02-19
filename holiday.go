package holiday

import (
	"encoding/json"
	"log"
	"time"

	"github.com/marcelblijleven/holiday/easter"
)

type holiday struct {
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Offical bool      `json:"official"`
}

func newHoliday(name string, date time.Time, official bool) holiday {
	return holiday{name, date, official}
}

// GetHolidays returns the mayor holidays in the given year
func GetHolidays(year int) ([]byte, error) {
	var holidays []holiday
	day := time.Hour * 24
	// First, determine the date of Easter
	easterDate := easter.GetEasterDateForYear(year)
	// Using the date of Easter, we can calculate the dynamic holidays
	goodFriday := newHoliday("Good Friday", easterDate.Add(day*-2), false)
	easter := newHoliday("Easter", easterDate, true)
	easterMonday := newHoliday("Easter Monday", easterDate.Add(day), true)
	ascensionDay := newHoliday("Ascension Day", easterDate.Add(day*39), true)
	pentecost := newHoliday("Pentecost", ascensionDay.Date.Add(day*10), true)
	pentecostMonday := newHoliday("Pentecost Monday", pentecost.Date.Add(day), true)
	holidays = append(
		holidays,
		goodFriday,
		easter,
		easterMonday,
		ascensionDay,
		pentecost,
		pentecostMonday,
	)

	// Now add the fixed holidays
	newYear := newHoliday("New Year", time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC), true)
	kingsDay := newHoliday("Kingsday", time.Date(year, 4, 27, 0, 0, 0, 0, time.UTC), true)

	if kingsDay.Date.Weekday() == 0 {
		// If Kingsday falls on a Sunday, move it to one day earlier
		kingsDay.Date.Add(day * -1)
	}

	liberationDay := newHoliday("Liberation Day", time.Date(year, 5, 5, 0, 0, 0, 0, time.UTC), false)

	if (year-2020)%5 == 0 {
		// Once every five years, you get a day off on Liberation Day
		liberationDay.Offical = true
	}

	christmas := newHoliday("Christmas", time.Date(year, 12, 25, 0, 0, 0, 0, time.UTC), true)
	boxingDay := newHoliday("Boxing Day", time.Date(year, 12, 26, 0, 0, 0, 0, time.UTC), true)

	holidays = append(
		holidays,
		newYear,
		kingsDay,
		liberationDay,
		christmas,
		boxingDay,
	)

	marshalled, err := json.MarshalIndent(holidays, "", "\t")

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return marshalled, nil
}
