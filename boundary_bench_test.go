package carbon

import (
	"testing"
)

func BenchmarkCarbon_StartOfCentury(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfCentury()
	}
}

func BenchmarkCarbon_EndOfCentury(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfCentury()
	}
}

func BenchmarkCarbon_StartOfDecade(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfDecade()
	}
}

func BenchmarkCarbon_EndOfDecade(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfDecade()
	}
}

func BenchmarkCarbon_StartOfYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfYear()
	}
}

func BenchmarkCarbon_EndOfYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfYear()
	}
}

func BenchmarkCarbon_StartOfQuarter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfQuarter()
	}
}

func BenchmarkCarbon_EndOfQuarter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfQuarter()
	}
}

func BenchmarkCarbon_StartOfMonth(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfMonth()
	}
}

func BenchmarkCarbon_EndOfMonth(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfMonth()
	}
}

func BenchmarkCarbon_StartOfWeek(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfWeek()
	}
}

func BenchmarkCarbon_EndOfWeek(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfWeek()
	}
}

func BenchmarkCarbon_StartOfDay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfDay()
	}
}

func BenchmarkCarbon_EndOfDay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfDay()
	}
}

func BenchmarkCarbon_StartOfHour(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfHour()
	}
}

func BenchmarkCarbon_EndOfHour(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfHour()
	}
}

func BenchmarkCarbon_StartOfMinute(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfMinute()
	}
}

func BenchmarkCarbon_EndOfMinute(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfMinute()
	}
}

func BenchmarkCarbon_StartOfSecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfSecond()
	}
}

func BenchmarkCarbon_EndOfSecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfSecond()
	}
}
