package carbon

import (
	"testing"
)

func BenchmarkCarbon_DiffInYears(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInYears(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInYears(c2)
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
				c1.DiffInYears(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInYears(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInYears(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInYears(c2)
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
				c1.DiffAbsInYears(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInMonths(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMonths(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInMonths(c2)
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
				c1.DiffInMonths(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInMonths(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMonths(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInMonths(c2)
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
				c1.DiffAbsInMonths(c2)
			}
		})
	})
}
func BenchmarkCarbon_DiffInWeeks(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInWeeks(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInWeeks(c2)
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
				c1.DiffInWeeks(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInWeeks(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInWeeks(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInWeeks(c2)
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
				c1.DiffAbsInWeeks(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInDays(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInDays(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInDays(c2)
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
				c1.DiffInDays(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInDays(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDays(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInDays(c2)
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
				c1.DiffAbsInDays(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInHours(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInHours(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInHours(c2)
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
				c1.DiffInHours(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInHours(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInHours(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInHours(c2)
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
				c1.DiffAbsInHours(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInMinutes(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMinutes(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInMinutes(c2)
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
				c1.DiffInMinutes(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInMinutes(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMinutes(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInMinutes(c2)
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
				c1.DiffAbsInMinutes(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInSeconds(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInSeconds(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInSeconds(c2)
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
				c1.DiffInSeconds(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInSeconds(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInSeconds(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInSeconds(c2)
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
				c1.DiffAbsInSeconds(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInString(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInString(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffInString(c2)
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
				c1.DiffInString(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInString(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInString(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInString(c2)
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
				c1.DiffAbsInString(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffInDuration(b *testing.B) {
	c := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.DiffInDuration()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c.DiffInDuration()
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
				c.DiffInDuration()
			}
		})
	})
}

func BenchmarkCarbon_DiffAbsInDuration(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDuration(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffAbsInDuration(c2)
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
				c1.DiffAbsInDuration(c2)
			}
		})
	})
}

func BenchmarkCarbon_DiffForHumans(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffForHumans(c2)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				c1.DiffForHumans(c2)
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
				c1.DiffForHumans(c2)
			}
		})
	})
}
