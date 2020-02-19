package holiday_test

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/marcelblijleven/holiday"
	"github.com/stretchr/testify/assert"
)

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}

func TestGetHolidays(t *testing.T) {
	expected2020 := helperLoadBytes(t, "expected2020.json")
	holidays2020, _ := holiday.GetHolidays(2020)
	assert.Equal(t, expected2020, holidays2020)
}

func TestGetHolidaysWithShiftedKingsday(t *testing.T) {
	expected2025 := helperLoadBytes(t, "expected2025.json")
	holidays2025, _ := holiday.GetHolidays(2025)
	assert.Equal(t, expected2025, holidays2025)
}

func TestGetHolidaysWithUnofficialLiberationDay(t *testing.T) {
	expected2021 := helperLoadBytes(t, "expected2021.json")
	holidays2021, _ := holiday.GetHolidays(2021)
	assert.Equal(t, expected2021, holidays2021)
}
