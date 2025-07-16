package hebrew

import (
	"testing"
	"time"
)

func BenchmarkFromStdTime(b *testing.B) {
	testDates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 6, 21, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 9, 22, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 12, 21, 0, 0, 0, 0, time.UTC),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		FromStdTime(date)
	}
}

func BenchmarkToGregorian(b *testing.B) {
	testHebrewDates := []*Hebrew{
		NewHebrew(5784, 1, 1),
		NewHebrew(5784, 6, 15),
		NewHebrew(5784, 12, 29),
		NewHebrew(5785, 1, 1),
		NewHebrew(5785, 12, 30),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := testHebrewDates[i%len(testHebrewDates)]
		h.ToGregorian()
	}
}

func BenchmarkToGregorianWithTimezone(b *testing.B) {
	testHebrewDates := []*Hebrew{
		NewHebrew(5784, 1, 1),
		NewHebrew(5784, 6, 15),
		NewHebrew(5784, 12, 29),
		NewHebrew(5785, 1, 1),
		NewHebrew(5785, 12, 30),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h := testHebrewDates[i%len(testHebrewDates)]
		h.ToGregorian("PRC")
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	testYears := []int{5780, 5781, 5784, 5785, 5787, 5788, 5790, 5791, 5793, 5794}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		h := NewHebrew(year, 1, 1)
		h.IsLeapYear()
	}
}

func BenchmarkYear(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Year()
	}
}

func BenchmarkMonth(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Month()
	}
}

func BenchmarkDay(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Day()
	}
}

func BenchmarkString(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.String()
	}
}

func BenchmarkToMonthString(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToMonthString()
	}
}

func BenchmarkToMonthStringEnLocale(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToMonthString(EnLocale)
	}
}

func BenchmarkToMonthStringHeLocale(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToMonthString(HeLocale)
	}
}

func BenchmarkToWeekString(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToWeekString()
	}
}

func BenchmarkToWeekStringEnLocale(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToWeekString(EnLocale)
	}
}

func BenchmarkToWeekStringHeLocale(b *testing.B) {
	h := NewHebrew(5784, 6, 15)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.ToWeekString(HeLocale)
	}
}

func BenchmarkGregorian2Jdn(b *testing.B) {
	testDates := []struct {
		year, month, day int
	}{
		{2024, 1, 1},
		{2024, 3, 20},
		{2024, 6, 21},
		{2024, 9, 22},
		{2024, 12, 21},
	}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		h.gregorian2jdn(date.year, date.month, date.day)
	}
}

func BenchmarkHebrew2Jdn(b *testing.B) {
	testDates := []struct {
		year, month, day int
	}{
		{5784, 1, 1},
		{5784, 6, 15},
		{5784, 12, 29},
		{5785, 1, 1},
		{5785, 12, 30},
	}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		h.hebrew2jdn(date.year, date.month, date.day)
	}
}

func BenchmarkJdn2Hebrew(b *testing.B) {
	testJdns := []float64{
		2460100.5, // 2024-01-01
		2460200.5, // 2024-04-10
		2460300.5, // 2024-07-19
		2460400.5, // 2024-10-27
		2460500.5, // 2025-02-04
	}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jdn := testJdns[i%len(testJdns)]
		h.jdn2hebrew(jdn)
	}
}

func BenchmarkJdn2Gregorian(b *testing.B) {
	testJdns := []int{
		2460100, // 2024-01-01
		2460200, // 2024-04-10
		2460300, // 2024-07-19
		2460400, // 2024-10-27
		2460500, // 2025-02-04
	}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		jdn := testJdns[i%len(testJdns)]
		h.jdn2gregorian(jdn)
	}
}

func BenchmarkGetFirstDelay(b *testing.B) {
	testYears := []int{5780, 5781, 5784, 5785, 5787, 5788, 5790, 5791, 5793, 5794}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		h.getFirstDelay(year)
	}
}

func BenchmarkGetSecondDelay(b *testing.B) {
	testYears := []int{5780, 5781, 5784, 5785, 5787, 5788, 5790, 5791, 5793, 5794}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		h.getSecondDelay(year)
	}
}

func BenchmarkGetMonthsInYear(b *testing.B) {
	testYears := []int{5780, 5781, 5784, 5785, 5787, 5788, 5790, 5791, 5793, 5794}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		h.getMonthsInYear(year)
	}
}

func BenchmarkGetDaysInYear(b *testing.B) {
	testYears := []int{5780, 5781, 5784, 5785, 5787, 5788, 5790, 5791, 5793, 5794}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		h.getDaysInYear(year)
	}
}

func BenchmarkGetDaysInMonth(b *testing.B) {
	testCases := []struct {
		year, month int
	}{
		{5784, 1},  // Nisan
		{5784, 6},  // Elul
		{5784, 7},  // Tishri
		{5784, 8},  // Heshvan
		{5784, 9},  // Kislev
		{5784, 12}, // Adar
		{5784, 13}, // Adar Bet (leap year)
		{5785, 12}, // Adar (non-leap year)
	}

	h := &Hebrew{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testCase := testCases[i%len(testCases)]
		h.getDaysInMonth(testCase.year, testCase.month)
	}
}

func BenchmarkNewHebrew(b *testing.B) {
	testDates := []struct {
		year, month, day int
	}{
		{5784, 1, 1},
		{5784, 6, 15},
		{5784, 12, 29},
		{5785, 1, 1},
		{5785, 12, 30},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		NewHebrew(date.year, date.month, date.day)
	}
}
