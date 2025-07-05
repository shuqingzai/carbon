package carbon

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Parse("2020-08-05 01:02:03")
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				Parse("2020-08-05 01:02:03")
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
				Parse("2020-08-05 01:02:03")
			}
		})
	})
}

func BenchmarkParseByLayout(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ParseByLayout("2020-08-05 13:14:15", DateTimeLayout)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				ParseByLayout("2020-08-05 13:14:15", DateTimeLayout)
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
				ParseByLayout("2020-08-05 13:14:15", DateTimeLayout)
			}
		})
	})
}

func BenchmarkParseByFormat(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ParseByFormat("2020-08-05 13:14:15", DateTimeFormat)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				ParseByFormat("2020-08-05 13:14:15", DateTimeFormat)
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
				ParseByFormat("2020-08-05 13:14:15", DateTimeFormat)
			}
		})
	})
}

func BenchmarkParseByLayouts(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout, DateTimeLayout})
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout, DateTimeLayout})
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
				ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout, DateTimeLayout})
			}
		})
	})
}

func BenchmarkParseByFormats(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ParseByFormats("2020-08-05 13:14:15", []string{DateFormat, DateTimeFormat})
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				ParseByFormats("2020-08-05 13:14:15", []string{DateFormat, DateTimeFormat})
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
				ParseByFormats("2020-08-05 13:14:15", []string{DateFormat, DateTimeFormat})
			}
		})
	})
}
