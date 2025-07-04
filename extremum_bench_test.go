package carbon

import (
	"testing"
)

func BenchmarkZeroValue(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			ZeroValue()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				ZeroValue()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkEpochValue(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			EpochValue()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				EpochValue()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMaxValue(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			MaxValue()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				MaxValue()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMinValue(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			MinValue()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				MinValue()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMaxDuration(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			MaxDuration()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				MaxDuration()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMinDuration(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			MinDuration()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				MinDuration()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMax(b *testing.B) {
	c1 := Parse("2020-08-06")
	c2 := Parse("2021-08-05")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Max(c1, c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				Max(c1, c2)
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkMin(b *testing.B) {
	c1 := Parse("2020-08-06")
	c2 := Parse("2021-08-05")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			Min(c1, c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				Min(c1, c2)
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_Closest(b *testing.B) {
	c1 := Parse("2020-08-04")
	c2 := Parse("2020-08-05")
	c3 := Parse("2021-08-06")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.Closest(c2, c3)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.Closest(c2, c3)
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_Farthest(b *testing.B) {
	c1 := Parse("2020-08-04")
	c2 := Parse("2020-08-05")
	c3 := Parse("2021-08-06")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.Farthest(c2, c3)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.Farthest(c2, c3)
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}
