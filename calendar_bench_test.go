package carbon

import (
	"testing"
)

func BenchmarkCarbon_Julian(b *testing.B) {
	c := Parse("2020-08-05")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Julian()
	}
}

func BenchmarkCreateFromJulian(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromJulian(2460333.051563)
	}
}

func BenchmarkCarbon_Lunar(b *testing.B) {
	c := Parse("2020-08-05")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Lunar()
	}
}

func BenchmarkCreateFromLunar(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromLunar(2023, 12, 11, false)
	}
}

func BenchmarkCarbon_Persian(b *testing.B) {
	c := Parse("2020-08-05")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Persian()
	}
}

func BenchmarkCreateFromPersian(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromPersian(1178, 10, 11)
	}
}
