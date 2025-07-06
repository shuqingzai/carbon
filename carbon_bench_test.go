package carbon

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkNewCarbon(b *testing.B) {
	loc, _ := time.LoadLocation(PRC)
	t, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			NewCarbon(t)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				NewCarbon(t)
			}()
		}
		wg.Wait()
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
	b.Run("sequential", func(b *testing.B) {
		c := Parse("2020-08-05").SetLocale("zh-CN")
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.Copy()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		c := Parse("2020-08-05").SetLocale("zh-CN")
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.Copy()
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		c := Parse("2020-08-05").SetLocale("zh-CN")
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Copy()
			}
		})
	})
}
