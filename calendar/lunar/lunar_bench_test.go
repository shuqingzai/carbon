package lunar

import (
	"testing"
	"time"
)

func BenchmarkFromStdTime(b *testing.B) {
	testDates := []time.Time{
		time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 2, 10, 0, 0, 0, 0, time.UTC), // 春节
		time.Date(2024, 6, 10, 0, 0, 0, 0, time.UTC), // 端午节
		time.Date(2024, 9, 17, 0, 0, 0, 0, time.UTC), // 中秋节
		time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		date := testDates[i%len(testDates)]
		FromStdTime(date)
	}
}

func BenchmarkToGregorian(b *testing.B) {
	testLunarDates := []*Lunar{
		NewLunar(2024, 1, 1, false),  // 春节
		NewLunar(2024, 5, 5, false),  // 端午节
		NewLunar(2024, 8, 15, false), // 中秋节
		NewLunar(2024, 12, 8, false), // 腊八节
		NewLunar(2024, 6, 1, true),   // 闰月
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testLunarDates[i%len(testLunarDates)]
		l.ToGregorian()
	}
}

func BenchmarkIsLeapYear(b *testing.B) {
	testYears := []int{2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		l := NewLunar(year, 1, 1, false)
		l.IsLeapYear()
	}
}

func BenchmarkIsValid(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 1, 1, false),
		NewLunar(2024, 12, 30, false),
		NewLunar(2024, 6, 1, true),    // 闰月
		NewLunar(1900, 1, 1, false),   // 边界值
		NewLunar(2100, 12, 31, false), // 边界值
		NewLunar(0, 1, 1, false),      // 无效
		NewLunar(2024, 13, 1, false),  // 无效
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.IsValid()
	}
}

func BenchmarkString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.String()
	}
}

func BenchmarkToYearString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ToYearString()
	}
}

func BenchmarkToMonthString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ToMonthString()
	}
}

func BenchmarkToDayString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ToDayString()
	}
}

func BenchmarkToDateString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ToDateString()
	}
}

func BenchmarkToWeekString(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.ToWeekString()
	}
}

func BenchmarkAnimal(b *testing.B) {
	l := NewLunar(2024, 6, 15, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Animal()
	}
}

func BenchmarkFestival(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 1, 1, false),  // 春节
		NewLunar(2024, 1, 15, false), // 元宵节
		NewLunar(2024, 5, 5, false),  // 端午节
		NewLunar(2024, 8, 15, false), // 中秋节
		NewLunar(2024, 12, 8, false), // 腊八节
		NewLunar(2024, 6, 15, false), // 无节日
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.Festival()
	}
}

func BenchmarkLeapMonth(b *testing.B) {
	testYears := []int{2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		l := NewLunar(year, 1, 1, false)
		l.LeapMonth()
	}
}

func BenchmarkIsLeapMonth(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 1, 1, false),
		NewLunar(2024, 6, 1, true),  // 闰月
		NewLunar(2024, 6, 1, false), // 非闰月
		NewLunar(2024, 12, 1, false),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.IsLeapMonth()
	}
}

func BenchmarkAnimalYearChecks(b *testing.B) {
	testYears := []int{2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		l := NewLunar(year, 1, 1, false)
		l.IsRatYear()
		l.IsOxYear()
		l.IsTigerYear()
		l.IsRabbitYear()
		l.IsDragonYear()
		l.IsSnakeYear()
		l.IsHorseYear()
		l.IsGoatYear()
		l.IsMonkeyYear()
		l.IsRoosterYear()
		l.IsDogYear()
		l.IsPigYear()
	}
}

func BenchmarkGetDaysInYear(b *testing.B) {
	testYears := []int{2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		l := NewLunar(year, 1, 1, false)
		l.getDaysInYear()
	}
}

func BenchmarkGetDaysInMonth(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 1, 1, false),
		NewLunar(2024, 2, 1, false),
		NewLunar(2024, 6, 1, false),
		NewLunar(2024, 12, 1, false),
		NewLunar(2024, 6, 1, true), // 闰月
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.getDaysInMonth()
	}
}

func BenchmarkGetDaysInLeapMonth(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 6, 1, true),  // 闰月
		NewLunar(2024, 1, 1, false), // 非闰月
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.getDaysInLeapMonth()
	}
}

func BenchmarkGetOffsetInYear(b *testing.B) {
	testYears := []int{2020, 2021, 2022, 2023, 2024, 2025, 2026, 2027, 2028, 2029}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		year := testYears[i%len(testYears)]
		l := NewLunar(year, 1, 1, false)
		l.getOffsetInYear()
	}
}

func BenchmarkGetOffsetInMonth(b *testing.B) {
	testDates := []*Lunar{
		NewLunar(2024, 1, 1, false),
		NewLunar(2024, 6, 15, false),
		NewLunar(2024, 12, 30, false),
		NewLunar(2024, 6, 1, true), // 闰月
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l := testDates[i%len(testDates)]
		l.getOffsetInMonth()
	}
}
