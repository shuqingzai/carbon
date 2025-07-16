package hebrew

import (
	"fmt"
	"math"
	"time"

	"github.com/dromara/carbon/v2/calendar"
	"github.com/dromara/carbon/v2/calendar/julian"
)

type Locale string

const (
	EnLocale      Locale = "en"
	HeLocale      Locale = "he"
	defaultLocale        = EnLocale
	hebrewEpoch          = 347995.5
)

var (
	EnMonths = []string{"Nisan", "Iyyar", "Sivan", "Tammuz", "Av", "Elul", "Tishri", "Heshvan", "Kislev", "Teveth", "Shevat", "Adar", "Adar Bet"}
	HeMonths = []string{"ניסן", "אייר", "סיוון", "תמוז", "אב", "אלול", "תשרי", "חשוון", "כסלו", "טבת", "שבט", "אדר", "אדר ב"}
	EnWeeks  = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	HeWeeks  = []string{"ראשון", "שני", "שלישי", "רביעי", "חמישי", "שישי", "שבת"}
)

type Hebrew struct {
	year, month, day int
	Error            error
}

// NewHebrew creates a new Hebrew calendar instance with the specified year, month, and day.
func NewHebrew(year, month, day int) *Hebrew {
	return &Hebrew{year: year, month: month, day: day}
}

// FromStdTime converts a standard time.Time to Hebrew calendar date.
func FromStdTime(t time.Time) *Hebrew {
	if t.IsZero() {
		return nil
	}
	h := &Hebrew{}
	jdn := h.gregorian2jdn(t.Year(), int(t.Month()), t.Day())
	y, m, d := h.jdn2hebrew(jdn)
	return &Hebrew{year: y, month: m, day: d}
}

// ToGregorian converts the Hebrew date to Gregorian calendar date.
func (h *Hebrew) ToGregorian(timezone ...string) *calendar.Gregorian {
	g := new(calendar.Gregorian)
	if h == nil {
		return g
	}
	loc := time.UTC
	if len(timezone) > 0 {
		loc, g.Error = time.LoadLocation(timezone[0])
	}
	if g.Error != nil {
		return g
	}
	jd := h.hebrew2jdn(h.year, h.month, h.day)
	date := julian.NewJulian(jd).ToGregorian()
	g.Time = time.Date(date.Time.Year(), date.Time.Month(), date.Time.Day(), 12, 0, 0, 0, loc)
	return g
}

// IsLeapYear determines if the Hebrew year is a leap year.
func (h *Hebrew) IsLeapYear() bool {
	if h == nil {
		return false
	}
	return ((h.year*7 + 1) % 19) < 7
}

// Year returns the Hebrew year.
func (h *Hebrew) Year() int {
	if h == nil {
		return 0
	}
	return h.year
}

// Month returns the Hebrew month (1-13, where 13 is Adar Bet in leap years).
func (h *Hebrew) Month() int {
	if h == nil {
		return 0
	}
	return h.month
}

// Day returns the Hebrew day of the month.
func (h *Hebrew) Day() int {
	if h == nil {
		return 0
	}
	return h.day
}

// String returns the Hebrew date in "YYYY-MM-DD" format.
func (h *Hebrew) String() string {
	if h == nil {
		return ""
	}
	return fmt.Sprintf("%04d-%02d-%02d", h.year, h.month, h.day)
}

// ToMonthString returns the Hebrew month name in the specified locale.
func (h *Hebrew) ToMonthString(locale ...Locale) string {
	if h == nil || h.month < 1 || h.month > 13 {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	idx := h.month - 1
	switch loc {
	case EnLocale:
		if idx >= 0 && idx < len(EnMonths) {
			return EnMonths[idx]
		}
	case HeLocale:
		if idx >= 0 && idx < len(HeMonths) {
			return HeMonths[idx]
		}
	}
	return ""
}

// ToWeekString returns the weekday name in the specified locale.
func (h *Hebrew) ToWeekString(locale ...Locale) string {
	if h == nil {
		return ""
	}
	loc := defaultLocale
	if len(locale) > 0 {
		loc = locale[0]
	}
	weekday := h.ToGregorian().Time.In(time.UTC).Weekday()
	switch loc {
	case EnLocale:
		return EnWeeks[weekday]
	case HeLocale:
		return HeWeeks[weekday]
	}
	return ""
}

// gregorian2jdn converts Gregorian date to Julian Day Number.
func (h *Hebrew) gregorian2jdn(year, month, day int) float64 {
	if month <= 2 {
		month += 12
		year--
	}
	jd := math.Floor(365.25*float64(year+4716)) +
		math.Floor(30.6001*float64(month+1)) +
		float64(day) - 1524.5
	if year*372+month*31+day >= 588829 {
		century := year / 100
		jd += float64(2 - century + century/4)
	}
	return jd
}

// hebrew2jdn converts Hebrew date to Julian Day Number.
func (h *Hebrew) hebrew2jdn(year, month, day int) float64 {
	firstDelay := float64(h.getFirstDelay(year))
	secondDelay := float64(h.getSecondDelay(year))
	newYearJDN := firstDelay + secondDelay + hebrewEpoch

	days := 0.0
	if month < 7 {
		for m := 7; m <= h.getMonthsInYear(year); m++ {
			days += float64(h.getDaysInMonth(year, m))
		}
		for m := 1; m < month; m++ {
			days += float64(h.getDaysInMonth(year, m))
		}
	} else {
		for m := 7; m < month; m++ {
			days += float64(h.getDaysInMonth(year, m))
		}
	}
	jd := newYearJDN + days + float64(day+1)
	return float64(int64(jd)) + 0.5
}

// jdn2hebrew converts Julian Day Number to Hebrew date.
func (h *Hebrew) jdn2hebrew(jdn float64) (year, month, day int) {
	jd := math.Floor(jdn) + 0.5

	count := int(math.Floor(((jd - hebrewEpoch) * 98496.0) / 35975351.0))
	year = count - 1
	i := count
	for jd >= h.hebrew2jdn(i, 7, 1) {
		i++
		year++
	}

	first := 1
	if jd < h.hebrew2jdn(year, 1, 1) {
		first = 7
	}

	month = first
	i = first
	for jd > h.hebrew2jdn(year, i, h.getDaysInMonth(year, i)) {
		i++
		month++
	}
	day = int(jd-h.hebrew2jdn(year, month, 1)) + 1
	return
}

// jdn2gregorian converts Julian Day Number to Gregorian date.
func (h *Hebrew) jdn2gregorian(jdn int) (year, month, day int) {
	jd := float64(jdn) + 0.5
	a := int(jd)
	b := a + 1524
	c := int((float64(b) - 122.1) / 365.25)
	d := int(365.25 * float64(c))
	e := int(float64(b-d) / 30.6001)
	if e < 14 {
		month = e - 1
	} else {
		month = e - 13
	}
	if month > 2 {
		year = c - 4716
	} else {
		year = c - 4715
	}
	day = b - d - int(30.6001*float64(e))
	return
}

// getFirstDelay calculates the first delay (Molad Tishri) for the Hebrew year.
func (h *Hebrew) getFirstDelay(year int) int {
	months := (235*year - 234) / 19
	parts := 12084 + 13753*months
	day := months*29 + parts/25920
	if (3*(day+1))%7 < 3 {
		day++
	}
	return day
}

// getSecondDelay calculates the second delay for the Hebrew year.
func (h *Hebrew) getSecondDelay(year int) int {
	last := h.getFirstDelay(year - 1)
	present := h.getFirstDelay(year)
	next := h.getFirstDelay(year + 1)

	if next-present == 356 {
		return 2
	}
	if present-last == 382 {
		return 1
	}
	return 0
}

// getMonthsInYear returns the number of months in the Hebrew year.
func (h *Hebrew) getMonthsInYear(year int) int {
	h.year = year
	if h.IsLeapYear() {
		return 13
	}
	return 12
}

// getDaysInYear returns the number of days in the Hebrew year.
func (h *Hebrew) getDaysInYear(year int) int {
	return int(h.hebrew2jdn(year+1, 7, 1) - h.hebrew2jdn(year, 7, 1))
}

// getDaysInMonth returns the number of days in the specified Hebrew month.
func (h *Hebrew) getDaysInMonth(year, month int) int {
	if month == 2 || month == 4 || month == 6 || month == 10 || month == 13 {
		return 29
	}
	if month == 12 && !h.IsLeapYear() {
		return 29
	}
	daysInYear := h.getDaysInYear(year)
	if month == 8 { // Heshvan
		if daysInYear == 355 || daysInYear == 385 {
			return 30
		}
		return 29
	}
	if month == 9 { // Kislev
		if daysInYear == 354 || daysInYear == 383 {
			return 29
		}
		return 30
	}
	return 30
}
