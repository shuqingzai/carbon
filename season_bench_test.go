package carbon

import (
	"sync"
	"testing"
)

func BenchmarkCarbon_Season(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.Season()
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
					c.Season()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.Season()
			}
		})
	})
}

func BenchmarkCarbon_StartOfSeason(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.StartOfSeason()
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
					c.StartOfSeason()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.StartOfSeason()
			}
		})
	})
}

func BenchmarkCarbon_EndOfSeason(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.EndOfSeason()
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
					c.EndOfSeason()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.EndOfSeason()
			}
		})
	})
}

func BenchmarkCarbon_IsSpring(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.IsSpring()
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
					c.IsSpring()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.IsSpring()
			}
		})
	})
}

func BenchmarkCarbon_IsSummer(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.IsSummer()
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
					c.IsSummer()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.IsSummer()
			}
		})
	})
}

func BenchmarkCarbon_IsAutumn(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.IsAutumn()
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
					c.IsAutumn()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.IsAutumn()
			}
		})
	})
}

func BenchmarkCarbon_IsWinter(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			c.IsWinter()
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
					c.IsWinter()
				}
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				c.IsWinter()
			}
		})
	})
}
