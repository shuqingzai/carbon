package carbon

import (
	"testing"
)

func BenchmarkCarbon_DiffInYears(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInYears()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInYears(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInYears(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInYears()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInYears(c2)
		}
	})
}

func BenchmarkCarbon_DiffInMonths(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMonths()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMonths(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInMonths(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMonths()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMonths(c2)
		}
	})
}
func BenchmarkCarbon_DiffInWeeks(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInWeeks()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInWeeks(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInWeeks(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInWeeks()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInWeeks(c2)
		}
	})
}

func BenchmarkCarbon_DiffInDays(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInDays()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInDays(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInDays(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDays()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDays(c2)
		}
	})
}

func BenchmarkCarbon_DiffInHours(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInHours()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInHours(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInHours(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInHours()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInHours(c2)
		}
	})
}

func BenchmarkCarbon_DiffInMinutes(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMinutes()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInMinutes(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInMinutes(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMinutes()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInMinutes(c2)
		}
	})
}

func BenchmarkCarbon_DiffInSeconds(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInSeconds()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInSeconds(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInSeconds(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInSeconds()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInSeconds(c2)
		}
	})
}

func BenchmarkCarbon_DiffInString(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInString()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffInString(c2)
		}
	})
}

func BenchmarkCarbon_DiffAbsInString(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInString()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInString(c2)
		}
	})
}

func BenchmarkCarbon_DiffInDuration(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.DiffInDuration()
	}
}

func BenchmarkCarbon_DiffAbsInDuration(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDuration()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffAbsInDuration(c2)
		}
	})
}

func BenchmarkCarbon_DiffForHumans(b *testing.B) {
	c1 := Parse("2020-08-05")
	c2 := Now()

	b.Run("without parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffForHumans()
		}
	})

	b.Run("with parameter", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c1.DiffForHumans(c2)
		}
	})
}
