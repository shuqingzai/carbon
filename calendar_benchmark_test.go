package carbon

import (
	"testing"
)

func BenchmarkCarbon_Julian(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Julian()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.Julian()
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
				c.Julian()
			}
		})
	})
}

func BenchmarkCreateFromJulian(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromJulian(2460333.051563)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromJulian(2460333.051563)
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
				CreateFromJulian(2460333.051563)
			}
		})
	})
}

func BenchmarkCarbon_Lunar(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Lunar()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.Lunar()
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
				c.Lunar()
			}
		})
	})
}

func BenchmarkCreateFromLunar(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromLunar(2023, 12, 11, false)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromLunar(2023, 12, 11, false)
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
				CreateFromLunar(2023, 12, 11, false)
			}
		})
	})
}

func BenchmarkCarbon_Persian(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Persian()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.Persian()
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
				c.Persian()
			}
		})
	})
}

func BenchmarkCreateFromPersian(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromPersian(1178, 10, 11)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromPersian(1178, 10, 11)
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
				CreateFromPersian(1178, 10, 11)
			}
		})
	})
}
