package easter_test

import (
	"fmt"
	"github.com/marcelblijleven/holiday/easter"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var easterTests = []struct {
	year int
	date time.Time
}{
	{1989, time.Date(1989, 3, 26, 0, 0, 0, 0, time.UTC)},
	{1990, time.Date(1990, 4, 15, 0, 0, 0, 0, time.UTC)},
	{1991, time.Date(1991, 3, 31, 0, 0, 0, 0, time.UTC)},
	{1992, time.Date(1992, 4, 19, 0, 0, 0, 0, time.UTC)},
	{1993, time.Date(1993, 4, 11, 0, 0, 0, 0, time.UTC)},
	{1994, time.Date(1994, 4, 3, 0, 0, 0, 0, time.UTC)},
	{1995, time.Date(1995, 4, 16, 0, 0, 0, 0, time.UTC)},
	{1996, time.Date(1996, 4, 7, 0, 0, 0, 0, time.UTC)},
	{1997, time.Date(1997, 3, 30, 0, 0, 0, 0, time.UTC)},
	{1998, time.Date(1998, 4, 12, 0, 0, 0, 0, time.UTC)},
	{1999, time.Date(1999, 4, 4, 0, 0, 0, 0, time.UTC)},
	{2000, time.Date(2000, 4, 23, 0, 0, 0, 0, time.UTC)},
	{2001, time.Date(2001, 4, 15, 0, 0, 0, 0, time.UTC)},
	{2002, time.Date(2002, 3, 31, 0, 0, 0, 0, time.UTC)},
	{2003, time.Date(2003, 4, 20, 0, 0, 0, 0, time.UTC)},
	{2004, time.Date(2004, 4, 11, 0, 0, 0, 0, time.UTC)},
	{2005, time.Date(2005, 3, 27, 0, 0, 0, 0, time.UTC)},
	{2006, time.Date(2006, 4, 16, 0, 0, 0, 0, time.UTC)},
	{2007, time.Date(2007, 4, 8, 0, 0, 0, 0, time.UTC)},
	{2008, time.Date(2008, 3, 23, 0, 0, 0, 0, time.UTC)},
	{2009, time.Date(2009, 4, 12, 0, 0, 0, 0, time.UTC)},
	{2010, time.Date(2010, 4, 4, 0, 0, 0, 0, time.UTC)},
	{2011, time.Date(2011, 4, 24, 0, 0, 0, 0, time.UTC)},
	{2012, time.Date(2012, 4, 8, 0, 0, 0, 0, time.UTC)},
	{2013, time.Date(2013, 3, 31, 0, 0, 0, 0, time.UTC)},
	{2014, time.Date(2014, 4, 20, 0, 0, 0, 0, time.UTC)},
	{2015, time.Date(2015, 4, 5, 0, 0, 0, 0, time.UTC)},
	{2016, time.Date(2016, 3, 27, 0, 0, 0, 0, time.UTC)},
	{2017, time.Date(2017, 4, 16, 0, 0, 0, 0, time.UTC)},
	{2018, time.Date(2018, 4, 1, 0, 0, 0, 0, time.UTC)},
	{2019, time.Date(2019, 4, 21, 0, 0, 0, 0, time.UTC)},
	{2020, time.Date(2020, 4, 12, 0, 0, 0, 0, time.UTC)},
	{2021, time.Date(2021, 4, 4, 0, 0, 0, 0, time.UTC)},
	{2022, time.Date(2022, 4, 17, 0, 0, 0, 0, time.UTC)},
	{2023, time.Date(2023, 4, 9, 0, 0, 0, 0, time.UTC)},
	{2024, time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)},
	{2025, time.Date(2025, 4, 20, 0, 0, 0, 0, time.UTC)},
	{2026, time.Date(2026, 4, 5, 0, 0, 0, 0, time.UTC)},
	{2027, time.Date(2027, 3, 28, 0, 0, 0, 0, time.UTC)},
	{2028, time.Date(2028, 4, 16, 0, 0, 0, 0, time.UTC)},
	{2029, time.Date(2029, 4, 1, 0, 0, 0, 0, time.UTC)},
	{2030, time.Date(2030, 4, 21, 0, 0, 0, 0, time.UTC)},
	{2031, time.Date(2031, 4, 13, 0, 0, 0, 0, time.UTC)},
	{2032, time.Date(2032, 3, 28, 0, 0, 0, 0, time.UTC)},
	{2033, time.Date(2033, 4, 17, 0, 0, 0, 0, time.UTC)},
	{2034, time.Date(2034, 4, 9, 0, 0, 0, 0, time.UTC)},
	{2035, time.Date(2035, 3, 25, 0, 0, 0, 0, time.UTC)},
	{2036, time.Date(2036, 4, 13, 0, 0, 0, 0, time.UTC)},
	{2037, time.Date(2037, 4, 5, 0, 0, 0, 0, time.UTC)},
	{2038, time.Date(2038, 4, 25, 0, 0, 0, 0, time.UTC)},
	{2039, time.Date(2039, 4, 10, 0, 0, 0, 0, time.UTC)},
	{2040, time.Date(2040, 4, 1, 0, 0, 0, 0, time.UTC)},
	{2041, time.Date(2041, 4, 21, 0, 0, 0, 0, time.UTC)},
	{2042, time.Date(2042, 4, 6, 0, 0, 0, 0, time.UTC)},
	{2043, time.Date(2043, 3, 29, 0, 0, 0, 0, time.UTC)},
	{2044, time.Date(2044, 4, 17, 0, 0, 0, 0, time.UTC)},
	{2045, time.Date(2045, 4, 9, 0, 0, 0, 0, time.UTC)},
	{2046, time.Date(2046, 3, 25, 0, 0, 0, 0, time.UTC)},
	{2047, time.Date(2047, 4, 14, 0, 0, 0, 0, time.UTC)},
	{2048, time.Date(2048, 4, 5, 0, 0, 0, 0, time.UTC)},
	{2049, time.Date(2049, 4, 18, 0, 0, 0, 0, time.UTC)},
	{2050, time.Date(2050, 4, 10, 0, 0, 0, 0, time.UTC)},
}

func TestGetEasterDateForYear(t *testing.T) {
	for _, tt := range easterTests {
		s := fmt.Sprintf("Easter test %v", tt.year)
		t.Run(s, func(t *testing.T) {
			assert.Equal(t, tt.date, easter.GetEasterDateForYear(tt.year))
		})
	}
}
