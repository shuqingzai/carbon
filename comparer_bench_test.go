package carbon

import (
	"testing"
)

func BenchmarkCarbon_HasError(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.HasError()
	}
}

func BenchmarkCarbon_IsNil(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsNil()
	}
}

func BenchmarkCarbon_IsZero(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsZero()
	}
}

func BenchmarkCarbon_IsEpoch(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsEpoch()
	}
}

func BenchmarkCarbon_IsValid(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsValid()
	}
}

func BenchmarkCarbon_IsInvalid(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsValid()
	}
}

func BenchmarkCarbon_IsDST(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsDST()
	}
}

func BenchmarkCarbon_IsAM(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsAM()
	}
}

func BenchmarkCarbon_IsPM(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsPM()
	}
}

func BenchmarkCarbon_IsLeapYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsLeapYear()
	}
}

func BenchmarkCarbon_IsLongYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsLongYear()
	}
}

func BenchmarkCarbon_IsJanuary(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsJanuary()
	}
}

func BenchmarkCarbon_IsFebruary(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsFebruary()
	}
}

func BenchmarkCarbon_IsMarch(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsMarch()
	}
}

func BenchmarkCarbon_IsApril(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsApril()
	}
}

func BenchmarkCarbon_IsMay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsMay()
	}
}

func BenchmarkCarbon_IsJune(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsJune()
	}
}

func BenchmarkCarbon_IsJuly(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsJuly()
	}
}

func BenchmarkCarbon_IsAugust(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsAugust()
	}
}

func BenchmarkCarbon_IsSeptember(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSeptember()
	}
}

func BenchmarkCarbon_IsOctober(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsOctober()
	}
}

func BenchmarkCarbon_IsNovember(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsNovember()
	}
}

func BenchmarkCarbon_IsDecember(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsDecember()
	}
}

func BenchmarkCarbon_IsMonday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsMonday()
	}
}

func BenchmarkCarbon_IsTuesday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsTuesday()
	}
}

func BenchmarkCarbon_IsWednesday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsWednesday()
	}
}

func BenchmarkCarbon_IsThursday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsThursday()
	}
}

func BenchmarkCarbon_IsFriday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsFriday()
	}
}

func BenchmarkCarbon_IsSaturday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSaturday()
	}
}

func BenchmarkCarbon_IsSunday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSunday()
	}
}

func BenchmarkCarbon_IsWeekday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsWeekday()
	}
}

func BenchmarkCarbon_IsWeekend(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsWeekend()
	}
}

func BenchmarkCarbon_IsNow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsNow()
	}
}

func BenchmarkCarbon_IsFuture(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsFuture()
	}
}

func BenchmarkCarbon_IsPast(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsPast()
	}
}

func BenchmarkCarbon_IsYesterday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsYesterday()
	}
}

func BenchmarkCarbon_IsToday(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsToday()
	}
}

func BenchmarkCarbon_IsTomorrow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsTomorrow()
	}
}

func BenchmarkCarbon_IsSameCentury(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameCentury(c)
	}
}

func BenchmarkCarbon_IsSameDecade(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameDecade(c)
	}
}

func BenchmarkCarbon_IsSameYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameYear(c)
	}
}

func BenchmarkCarbon_IsSameQuarter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameQuarter(c)
	}
}

func BenchmarkCarbon_IsSameMonth(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameMonth(c)
	}
}

func BenchmarkCarbon_IsSameDay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameDay(c)
	}
}

func BenchmarkCarbon_IsSameHour(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameHour(c)
	}
}

func BenchmarkCarbon_IsSameMinute(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameMinute(c)
	}
}

func BenchmarkCarbon_IsSameSecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.IsSameSecond(c)
	}
}

func BenchmarkCarbon_Compare(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Compare("=", c)
	}
}

func BenchmarkCarbon_Gt(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Gt(c)
	}
}

func BenchmarkCarbon_Lt(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Lt(c)
	}
}

func BenchmarkCarbon_Eq(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Eq(c)
	}
}

func BenchmarkCarbon_Ne(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Ne(c)
	}
}

func BenchmarkCarbon_Gte(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Gte(c)
	}
}

func BenchmarkCarbon_Lte(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Lte(c)
	}
}

func BenchmarkCarbon_Between(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("true", func(b *testing.B) {
		start, end := Parse("2020-08-04"), Parse("2020-08-06")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Between(start, end)
		}
	})

	b.Run("false", func(b *testing.B) {
		start, end := Parse("2020-08-06"), Parse("2020-08-07")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Between(start, end)
		}
	})
}

func BenchmarkCarbon_BetweenIncludedStart(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("true", func(b *testing.B) {
		start, end := Parse("2020-08-05"), Parse("2020-08-06")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedStart(start, end)
		}
	})

	b.Run("false", func(b *testing.B) {
		start, end := Parse("2020-08-06"), Parse("2020-08-07")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedStart(start, end)
		}
	})
}

func BenchmarkCarbon_BetweenIncludedEnd(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("true", func(b *testing.B) {
		start, end := Parse("2020-08-03"), Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedEnd(start, end)
		}
	})

	b.Run("false", func(b *testing.B) {
		start, end := Parse("2020-08-06"), Parse("2020-08-07")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedEnd(start, end)
		}
	})
}

func BenchmarkCarbon_BetweenIncludedBoth(b *testing.B) {
	c := Parse("2020-08-05")

	b.Run("true", func(b *testing.B) {
		start, end := Parse("2020-08-03"), Parse("2020-08-06")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedBoth(start, end)
		}
	})

	b.Run("false", func(b *testing.B) {
		start, end := Parse("2020-08-06"), Parse("2020-08-07")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.BetweenIncludedBoth(start, end)
		}
	})
}
