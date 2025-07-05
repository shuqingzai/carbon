package carbon

import (
	"testing"
	"time"
)

func BenchmarkNewCarbon(b *testing.B) {
	loc, _ := time.LoadLocation(PRC)
	t, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	b.Run("without time", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			NewCarbon()
		}
	})

	b.Run("with time", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			NewCarbon(t)
		}
	})
}

func BenchmarkCopy(b *testing.B) {
	c := Parse("2020-08-05").SetLocale("zh-CN")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Copy()
	}
}
