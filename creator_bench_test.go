package carbon

import (
	"testing"
	"time"
)

func BenchmarkCreateFromStdTime(b *testing.B) {
	now := time.Now().In(time.Local)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromStdTime(now)
	}
}

func BenchmarkCreateFromTimestamp(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimestamp(1649735755, PRC)
	}
}

func BenchmarkCreateFromTimestampMilli(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMilli(1649735755981, PRC)
	}
}

func BenchmarkCreateFromTimestampMicro(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimestampMicro(1649735755981566, PRC)
	}
}

func BenchmarkCreateFromTimestampNano(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimestampNano(1649735755981566888, PRC)
	}
}

func BenchmarkCreateFromDateTime(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateTime(2020, 8, 5, 13, 14, 15, PRC)
	}
}

func BenchmarkCreateFromDateTimeMilli(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999, PRC)
	}
}

func BenchmarkCreateFromDateTimeMicro(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999, PRC)
	}
}

func BenchmarkCreateFromDateTimeNano(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999, PRC)
	}
}

func BenchmarkCreateFromDate(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDate(2020, 8, 5, PRC)
	}
}

func BenchmarkCreateFromDateMilli(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateMilli(2020, 8, 5, 999, PRC)
	}
}

func BenchmarkCreateFromDateMicro(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateMicro(2020, 8, 5, 999999, PRC)
	}
}

func BenchmarkCreateFromDateNano(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromDateNano(2020, 8, 5, 999999999, PRC)
	}
}

func BenchmarkCreateFromTime(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTime(13, 14, 15, PRC)
	}
}

func BenchmarkCreateFromTimeMilli(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimeMilli(13, 14, 15, 999, PRC)
	}
}

func BenchmarkCreateFromTimeMicro(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimeMicro(13, 14, 15, 999999, PRC)
	}
}

func BenchmarkCreateFromTimeNano(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CreateFromTimeNano(13, 14, 15, 999999999, PRC)
	}
}
