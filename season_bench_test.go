package carbon

import (
	"testing"
)

func BenchmarkCarbon_Season(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Season()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.Season()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_StartOfSeason(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.StartOfSeason()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.StartOfSeason()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_EndOfSeason(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.EndOfSeason()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.EndOfSeason()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_IsSpring(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsSpring()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.IsSpring()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_IsSummer(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsSummer()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.IsSummer()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_IsAutumn(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsAutumn()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.IsAutumn()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}

func BenchmarkCarbon_IsWinter(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsWinter()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.IsWinter()
				done <- true
			}()
		}
		for i := 0; i < b.N; i++ {
			<-done
		}
	})
}
