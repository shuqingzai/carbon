package carbon

import (
	"testing"
)

func BenchmarkSetTestNow(b *testing.B) {
	defer CleanTestNow()

	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SetTestNow(c)
	}
}

func BenchmarkCleanTestNow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CleanTestNow()
	}
}

func BenchmarkIsTestNow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		IsTestNow()
	}
}
