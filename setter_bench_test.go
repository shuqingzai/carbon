package carbon

import (
	"testing"
	"time"
)

func BenchmarkSetLayout(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetLayout(DateTimeLayout)
	}
}

func BenchmarkSetFormat(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetFormat(DateTimeFormat)
	}
}

func BenchmarkSetWeekStartsAt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetWeekStartsAt(Monday)
	}
}

func BenchmarkSetTimezone(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetTimezone(UTC)
	}
}

func BenchmarkSetLocation(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetLocation(time.UTC)
	}
}

func BenchmarkSetLocale(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SetLocale("en")
	}
}

func BenchmarkCarbon_SetLayout(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetLayout(DateTimeLayout)
	}
}

func BenchmarkCarbon_SetFormat(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetFormat(DateTimeFormat)
	}
}

func BenchmarkCarbon_SetWeekStartsAt(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetWeekStartsAt(Monday)
	}
}

func BenchmarkCarbon_SetLocale(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetLocale("en")
	}
}

func BenchmarkCarbon_SetTimezone(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetTimezone(UTC)
	}
}

func BenchmarkCarbon_SetLocation(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetLocation(time.UTC)
	}
}

func BenchmarkCarbon_SetLanguage(b *testing.B) {
	c := Now()
	lang := NewLanguage()
	lang.SetLocale("en")
	for n := 0; n < b.N; n++ {
		c.SetLanguage(lang)
	}
}

func BenchmarkCarbon_SetDateTime(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateTime(2020, 8, 5, 13, 14, 15)
	}
}

func BenchmarkCarbon_SetDateTimeMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
	}
}

func BenchmarkCarbon_SetDateTimeMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
	}
}

func BenchmarkCarbon_SetDateTimeNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999999)
	}
}

func BenchmarkCarbon_SetDate(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDate(2020, 8, 5)
	}
}

func BenchmarkCarbon_SetDateMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateMilli(2020, 8, 5, 999)
	}
}

func BenchmarkCarbon_SetDateMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateMicro(2020, 8, 5, 999999)
	}
}

func BenchmarkCarbon_SetDateNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDateNano(2020, 8, 5, 999999999)
	}
}

func BenchmarkCarbon_SetTime(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetTime(13, 14, 15)
	}
}

func BenchmarkCarbon_SetTimeMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetTimeMilli(13, 14, 15, 999)
	}
}

func BenchmarkCarbon_SetTimeMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetTimeMicro(13, 14, 15, 999999)
	}
}

func BenchmarkCarbon_SetTimeNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetTimeNano(13, 14, 15, 999999999)
	}
}

func BenchmarkCarbon_SetYear(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetYear(2020)
	}
}

func BenchmarkCarbon_SetYearNoOverflow(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetYearNoOverflow(2020)
	}
}

func BenchmarkCarbon_SetMonth(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetMonth(2)
	}
}

func BenchmarkCarbon_SetMonthNoOverflow(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetMonthNoOverflow(2)
	}
}

func BenchmarkCarbon_SetDay(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetDay(31)
	}
}

func BenchmarkCarbon_SetHour(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetHour(10)
	}
}

func BenchmarkCarbon_SetMinute(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetMinute(10)
	}
}

func BenchmarkCarbon_SetSecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetSecond(10)
	}
}

func BenchmarkCarbon_SetMillisecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetMillisecond(999)
	}
}

func BenchmarkCarbon_SetMicrosecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetMicrosecond(999999)
	}
}

func BenchmarkCarbon_SetNanosecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.SetNanosecond(999999999)
	}
}
