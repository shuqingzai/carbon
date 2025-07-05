package carbon

import (
	"testing"
	"time"
)

func BenchmarkCreateFromStdTime(b *testing.B) {
	now := time.Now().In(time.Local)

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromStdTime(now)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromStdTime(now)
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
				CreateFromStdTime(now)
			}
		})
	})
}

func BenchmarkCreateFromTimestamp(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimestamp(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimestamp(1649735755)
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
				CreateFromTimestamp(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimestampMilli(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimestampMilli(1649735755)
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
				CreateFromTimestampMilli(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimestampMicro(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimestampMicro(1649735755)
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
				CreateFromTimestampMicro(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromTimestampNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimestampNano(1649735755)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimestampNano(1649735755)
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
				CreateFromTimestampNano(1649735755)
			}
		})
	})
}

func BenchmarkCreateFromDateTime(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateTime(2020, 8, 5, 13, 14, 15)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateTime(2020, 8, 5, 13, 14, 15)
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
				CreateFromDateTime(2020, 8, 5, 13, 14, 15)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
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
				CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
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
				CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999)
			}
		})
	})
}

func BenchmarkCreateFromDateTimeNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
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
				CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999)
			}
		})
	})
}

func BenchmarkCreateFromDate(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDate(2020, 8, 5)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDate(2020, 8, 5)
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
				CreateFromDate(2020, 8, 5)
			}
		})
	})
}

func BenchmarkCreateFromDateMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateMilli(2020, 8, 5, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateMilli(2020, 8, 5, 999)
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
				CreateFromDateMilli(2020, 8, 5, 999)
			}
		})
	})
}

func BenchmarkCreateFromDateMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateMicro(2020, 8, 5, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateMicro(2020, 8, 5, 999999)
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
				CreateFromDateMicro(2020, 8, 5, 999999)
			}
		})
	})
}

func BenchmarkCreateFromDateNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromDateNano(2020, 8, 5, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromDateNano(2020, 8, 5, 999999999)
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
				CreateFromDateNano(2020, 8, 5, 999999999)
			}
		})
	})
}

func BenchmarkCreateFromTime(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTime(13, 14, 15)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTime(13, 14, 15)
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
				CreateFromTime(13, 14, 15)
			}
		})
	})
}

func BenchmarkCreateFromTimeMilli(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimeMilli(13, 14, 15, 999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimeMilli(13, 14, 15, 999)
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
				CreateFromTimeMilli(13, 14, 15, 999)
			}
		})
	})
}

func BenchmarkCreateFromTimeMicro(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimeMicro(13, 14, 15, 999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimeMicro(13, 14, 15, 999999)
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
				CreateFromTimeMicro(13, 14, 15, 999999)
			}
		})
	})
}

func BenchmarkCreateFromTimeNano(b *testing.B) {
	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			CreateFromTimeNano(13, 14, 15, 999999999)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				CreateFromTimeNano(13, 14, 15, 999999999)
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
				CreateFromTimeNano(13, 14, 15, 999999999)
			}
		})
	})
}
