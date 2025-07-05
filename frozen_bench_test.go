package carbon

import (
	"testing"
)

func BenchmarkSetTestNow(b *testing.B) {
	defer ClearTestNow()

	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		SetTestNow(c)
	}
}

func BenchmarkClearTestNow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ClearTestNow()
	}
}

func BenchmarkIsTestNow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		IsTestNow()
	}
}
