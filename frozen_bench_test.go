package carbon

import (
	"sync"
	"testing"
)

func BenchmarkSetTestNow(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		defer ClearTestNow()
		c := Now()
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			SetTestNow(c)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		defer ClearTestNow()
		var wg sync.WaitGroup
		c := Now()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				SetTestNow(c)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		defer ClearTestNow()
		c := Now()
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				SetTestNow(c)
			}
		})
	})
}

func BenchmarkClearTestNow(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			ClearTestNow()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ClearTestNow()
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ClearTestNow()
			}
		})
	})
}

func BenchmarkIsTestNow(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			IsTestNow()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				IsTestNow()
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				IsTestNow()
			}
		})
	})
}
