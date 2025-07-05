package carbon

import (
	"testing"
	"time"
)

func BenchmarkNewCarbon(b *testing.B) {
	loc, _ := time.LoadLocation(PRC)
	t, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			NewCarbon(t)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				NewCarbon(t)
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				NewCarbon(t)
			}
		})
	})
}

func BenchmarkCopy(b *testing.B) {
	c := Parse("2020-08-05").SetLocale("zh-CN")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Copy()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.Copy()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Copy()
			}
		})
	})
}
