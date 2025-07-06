package carbon

import (
	"sync"
	"testing"
)

func BenchmarkParse(b *testing.B) {
	dateStr := "2020-08-05 01:02:03"

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			Parse(dateStr)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := 0; n < b.N/10; n++ {
					Parse(dateStr)
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				Parse(dateStr)
			}
		})
	})
}

func BenchmarkParseByLayout(b *testing.B) {
	dateStr := "2020-08-05 13:14:15"
	layout := DateTimeLayout

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			ParseByLayout(dateStr, layout)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := 0; n < b.N/10; n++ {
					ParseByLayout(dateStr, layout)
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ParseByLayout(dateStr, layout)
			}
		})
	})
}

func BenchmarkParseByFormat(b *testing.B) {
	dateStr := "2020-08-05 13:14:15"
	format := DateTimeFormat

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			ParseByFormat(dateStr, format)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := 0; n < b.N/10; n++ {
					ParseByFormat(dateStr, format)
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ParseByFormat(dateStr, format)
			}
		})
	})
}

func BenchmarkParseByLayouts(b *testing.B) {
	dateStr := "2020-08-05 13:14:15"
	layouts := []string{DateLayout, DateTimeLayout}

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			ParseByLayouts(dateStr, layouts)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := 0; n < b.N/10; n++ {
					ParseByLayouts(dateStr, layouts)
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ParseByLayouts(dateStr, layouts)
			}
		})
	})
}

func BenchmarkParseByFormats(b *testing.B) {
	dateStr := "2020-08-05 13:14:15"
	formats := []string{DateFormat, DateTimeFormat}

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			ParseByFormats(dateStr, formats)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		var wg sync.WaitGroup
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for n := 0; n < b.N/10; n++ {
					ParseByFormats(dateStr, formats)
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ParseByFormats(dateStr, formats)
			}
		})
	})
}
