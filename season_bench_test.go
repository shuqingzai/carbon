package carbon

import (
	"testing"
)

func BenchmarkCarbon_Season(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Season()
	}
}

func BenchmarkCarbon_StartOfSeason(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.StartOfSeason()
	}
}

func BenchmarkCarbon_EndOfSeason(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.EndOfSeason()
	}
}

func BenchmarkCarbon_IsSpring(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSpring()
	}
}

func BenchmarkCarbon_IsSummer(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSummer()
	}
}

func BenchmarkCarbon_IsAutumn(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsAutumn()
	}
}

func BenchmarkCarbon_IsWinter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsWinter()
	}
}
