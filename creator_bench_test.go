package carbon

import (
	"sync"
	"testing"
	"time"
)

func BenchmarkCreateFromStdTime(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		now := time.Now().In(time.Local)
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromStdTime(now)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		now := time.Now().In(time.Local)
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromStdTime(now)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		now := time.Now().In(time.Local)
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromStdTime(now)
			}
		})
	})
}

func BenchmarkCreateFromTimestamp(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimestamp(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimestamp(1649735755)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimestamp(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimestampMilli(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimestampMilli(1649735755)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimestampMilli(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimestampMicro(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimestampMicro(1649735755)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimestampMicro(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimestampNano(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimestampNano(1649735755)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimestampNano(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromDateTime(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateTime(2020, 8, 5, 13, 14, 15)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateTime(2020, 8, 5, 13, 14, 15)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateTime(2020, 8, 5, 13, 14, 15)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
			}
		})
	})
}

func BenchmarkCreateFromDate(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDate(2020, 8, 5)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDate(2020, 8, 5)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDate(2020, 8, 5)
			}
		})
	})
}

func BenchmarkCreateFromDateMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateMilli(2020, 8, 5, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateMilli(2020, 8, 5, 999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateMilli(2020, 8, 5, 999)
			}
		})
	})
}

func BenchmarkCreateFromDateMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateMicro(2020, 8, 5, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateMicro(2020, 8, 5, 999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateMicro(2020, 8, 5, 999999)
			}
		})
	})
}

func BenchmarkCreateFromDateNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromDateNano(2020, 8, 5, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromDateNano(2020, 8, 5, 999999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromDateNano(2020, 8, 5, 999999999)
			}
		})
	})
}

func BenchmarkCreateFromTime(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTime(13, 14, 15)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTime(13, 14, 15)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTime(13, 14, 15)
			}
		})
	})
}

func BenchmarkCreateFromTimeMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimeMilli(13, 14, 15, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimeMilli(13, 14, 15, 999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimeMilli(13, 14, 15, 999)
			}
		})
	})
}

func BenchmarkCreateFromTimeMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimeMicro(13, 14, 15, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimeMicro(13, 14, 15, 999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimeMicro(13, 14, 15, 999999)
			}
		})
	})
}

func BenchmarkCreateFromTimeNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < 10; n++ {
			CreateFromTimeNano(13, 14, 15, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		var wg sync.WaitGroup
		b.ResetTimer()
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				CreateFromTimeNano(13, 14, 15, 999999999)
			}()
		}
		wg.Wait()
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				CreateFromTimeNano(13, 14, 15, 999999999)
			}
		})
	})
}
