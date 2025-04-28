package carbon

import (
	"testing"
)

func BenchmarkCarbon_DiffInYears(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInYears()
	}
}

func BenchmarkCarbon_DiffAbsInYears(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInYears()
	}
}

func BenchmarkCarbon_DiffInMonths(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInMonths()
	}
}

func BenchmarkCarbon_DiffAbsInMonths(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInMonths()
	}
}
func BenchmarkCarbon_DiffInWeeks(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DiffInWeeks()
	}
}

func BenchmarkCarbon_DiffAbsInWeeks(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInWeeks()
	}
}

func BenchmarkCarbon_DiffInDays(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInDays()
	}
}

func BenchmarkCarbon_DiffAbsInDays(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInDays()
	}
}

func BenchmarkCarbon_DiffInHours(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInHours()
	}
}

func BenchmarkCarbon_DiffAbsInHours(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInHours()
	}
}

func BenchmarkCarbon_DiffInMinutes(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInMinutes()
	}
}

func BenchmarkCarbon_DiffAbsInMinutes(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInMinutes()
	}
}

func BenchmarkCarbon_DiffInSeconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInSeconds()
	}
}

func BenchmarkCarbon_DiffAbsInSeconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInSeconds()
	}
}

func BenchmarkCarbon_DiffInString(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInString()
	}
}

func BenchmarkCarbon_DiffAbsInString(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInString()
	}
}

func BenchmarkCarbon_DiffInDuration(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInDuration()
	}
}

func BenchmarkCarbon_DiffAbsInDuration(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffAbsInDuration()
	}
}

func BenchmarkCarbon_DiffForHumans(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffForHumans()
	}
}
