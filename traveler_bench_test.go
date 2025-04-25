package carbon

import (
	"testing"
)

func BenchmarkNow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Now()
	}
}

func BenchmarkTomorrow(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Tomorrow()
	}
}

func BenchmarkYesterday(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Yesterday()
	}
}

func BenchmarkCarbon_AddDuration(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDuration("10.5m")
	}
}

func BenchmarkCarbon_SubDuration(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDuration("10.5m")
	}
}

func BenchmarkCarbon_AddCenturies(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddCenturies(2)
	}
}

func BenchmarkCarbon_AddCenturiesNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddCenturiesNoOverflow(2)
	}
}

func BenchmarkCarbon_AddCentury(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddCentury()
	}
}

func BenchmarkCarbon_AddCenturyNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddCenturyNoOverflow()
	}
}

func BenchmarkCarbon_SubCenturies(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubCenturies(2)
	}
}

func BenchmarkCarbon_SubCenturiesNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubCenturiesNoOverflow(2)
	}
}

func BenchmarkCarbon_SubCentury(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubCentury()
	}
}

func BenchmarkCarbon_SubCenturyNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubCenturyNoOverflow()
	}
}

func BenchmarkCarbon_AddDecades(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDecades(2)
	}
}

func BenchmarkCarbon_AddDecadesNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDecadesNoOverflow(2)
	}
}

func BenchmarkCarbon_AddDecade(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDecade()
	}
}

func BenchmarkCarbon_AddDecadeNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDecadeNoOverflow()
	}
}

func BenchmarkCarbon_SubDecades(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDecades(2)
	}
}

func BenchmarkCarbon_SubDecadesNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDecadesNoOverflow(2)
	}
}

func BenchmarkCarbon_SubDecade(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDecade()
	}
}

func BenchmarkCarbon_SubDecadeNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDecadeNoOverflow()
	}
}

func BenchmarkCarbon_AddYears(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddYears(2)
	}
}

func BenchmarkCarbon_AddYearsNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddYearsNoOverflow(2)
	}
}

func BenchmarkCarbon_AddYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddYear()
	}
}

func BenchmarkCarbon_AddYearNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddYearNoOverflow()
	}
}

func BenchmarkCarbon_SubYears(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubYears(2)
	}
}

func BenchmarkCarbon_SubYearsNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubYearsNoOverflow(2)
	}
}

func BenchmarkCarbon_SubYear(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubYear()
	}
}

func BenchmarkCarbon_SubYearNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubYearNoOverflow()
	}
}

func BenchmarkCarbon_AddQuarters(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddQuarters(2)
	}
}

func BenchmarkCarbon_AddQuartersNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddQuartersNoOverflow(2)
	}
}

func BenchmarkCarbon_AddQuarter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddQuarter()
	}
}

func BenchmarkCarbon_AddQuarterNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddQuarterNoOverflow()
	}
}

func BenchmarkCarbon_SubQuarters(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubQuarters(2)
	}
}

func BenchmarkCarbon_SubQuartersNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubQuartersNoOverflow(2)
	}
}

func BenchmarkCarbon_SubQuarter(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubQuarter()
	}
}

func BenchmarkCarbon_SubQuarterNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubQuarterNoOverflow()
	}
}

func BenchmarkCarbon_AddMonths(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMonths(2)
	}
}

func BenchmarkCarbon_AddMonthsNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMonthsNoOverflow(2)
	}
}

func BenchmarkCarbon_AddMonth(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMonth()
	}
}

func BenchmarkCarbon_AddMonthNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMonthNoOverflow()
	}
}

func BenchmarkCarbon_SubMonths(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMonths(2)
	}
}

func BenchmarkCarbon_SubMonthsNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMonthsNoOverflow(2)
	}
}

func BenchmarkCarbon_SubMonth(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMonth()
	}
}

func BenchmarkCarbon_SubMonthNoOverflow(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMonthNoOverflow()
	}
}

func BenchmarkCarbon_AddWeeks(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddWeeks(2)
	}
}

func BenchmarkCarbon_AddWeek(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddWeek()
	}
}

func BenchmarkCarbon_SubWeeks(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubWeeks(2)
	}
}

func BenchmarkCarbon_SubWeek(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubWeek()
	}
}

func BenchmarkCarbon_AddDays(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDays(2)
	}
}

func BenchmarkCarbon_AddDay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddDay()
	}
}

func BenchmarkCarbon_SubDays(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDays(2)
	}
}

func BenchmarkCarbon_SubDay(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubDay()
	}
}

func BenchmarkCarbon_AddHours(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddHours(2)
	}
}

func BenchmarkCarbon_AddHour(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddHour()
	}
}

func BenchmarkCarbon_SubHours(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubHours(2)
	}
}

func BenchmarkCarbon_SubHour(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubHour()
	}
}

func BenchmarkCarbon_AddMinutes(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMinutes(2)
	}
}

func BenchmarkCarbon_AddMinute(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMinute()
	}
}

func BenchmarkCarbon_SubMinutes(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMinutes(2)
	}
}

func BenchmarkCarbon_SubMinute(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMinute()
	}
}

func BenchmarkCarbon_AddSeconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddSeconds(2)
	}
}

func BenchmarkCarbon_AddSecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddSecond()
	}
}

func BenchmarkCarbon_SubSeconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubSeconds(2)
	}
}

func BenchmarkCarbon_SubSecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubSecond()
	}
}

func BenchmarkCarbon_AddMilliseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMilliseconds(2)
	}
}

func BenchmarkCarbon_AddMillisecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMillisecond()
	}
}

func BenchmarkCarbon_SubMilliseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMilliseconds(2)
	}
}

func BenchmarkCarbon_SubMillisecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMillisecond()
	}
}

func BenchmarkCarbon_AddMicroseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMicroseconds(2)
	}
}

func BenchmarkCarbon_AddMicrosecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddMicrosecond()
	}
}

func BenchmarkCarbon_SubMicroseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMicroseconds(2)
	}
}

func BenchmarkCarbon_SubMicrosecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubMicrosecond()
	}
}

func BenchmarkCarbon_AddNanoseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddNanoseconds(2)
	}
}

func BenchmarkCarbon_AddNanosecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.AddNanosecond()
	}
}

func BenchmarkCarbon_SubNanoseconds(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubNanoseconds(2)
	}
}

func BenchmarkCarbon_SubNanosecond(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.SubNanosecond()
	}
}
