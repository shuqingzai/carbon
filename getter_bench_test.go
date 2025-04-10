package carbon

import (
	"testing"
)

func BenchmarkCarbon_StdTime(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.StdTime()
	}
}

func BenchmarkCarbon_DaysInYear(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DaysInYear()
	}
}

func BenchmarkCarbon_DaysInMonth(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DaysInMonth()
	}
}

func BenchmarkCarbon_MonthOfYear(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.MonthOfYear()
	}
}

func BenchmarkCarbon_DayOfYear(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DayOfYear()
	}
}

func BenchmarkCarbon_DayOfMonth(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DayOfMonth()
	}
}

func BenchmarkCarbon_DayOfWeek(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DayOfWeek()
	}
}

func BenchmarkCarbon_WeekOfYear(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.WeekOfYear()
	}
}

func BenchmarkCarbon_WeekOfMonth(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.WeekOfMonth()
	}
}

func BenchmarkCarbon_DateTime(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateTime()
	}
}

func BenchmarkCarbon_DateTimeMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateTimeMilli()
	}
}

func BenchmarkCarbon_DateTimeMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateTimeMicro()
	}
}

func BenchmarkCarbon_DateTimeNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateTimeNano()
	}
}

func BenchmarkCarbon_Date(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Date()
	}
}

func BenchmarkCarbon_DateMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateMilli()
	}
}

func BenchmarkCarbon_DateMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateMicro()
	}
}

func BenchmarkCarbon_DateNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.DateNano()
	}
}

func BenchmarkCarbon_Time(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Time()
	}
}

func BenchmarkCarbon_TimeMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimeMilli()
	}
}

func BenchmarkCarbon_TimeMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimeMicro()
	}
}

func BenchmarkCarbon_TimeNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimeNano()
	}
}

func BenchmarkCarbon_Century(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Century()
	}
}

func BenchmarkCarbon_Decade(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Decade()
	}
}

func BenchmarkCarbon_Year(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Year()
	}
}

func BenchmarkCarbon_Quarter(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Quarter()
	}
}

func BenchmarkCarbon_Month(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Month()
	}
}

func BenchmarkCarbon_Week(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Week()
	}
}

func BenchmarkCarbon_Day(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Day()
	}
}

func BenchmarkCarbon_Hour(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Hour()
	}
}

func BenchmarkCarbon_Minute(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Minute()
	}
}

func BenchmarkCarbon_Second(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Second()
	}
}

func BenchmarkCarbon_Millisecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Millisecond()
	}
}

func BenchmarkCarbon_Microsecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Microsecond()
	}
}

func BenchmarkCarbon_Nanosecond(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Nanosecond()
	}
}

func BenchmarkCarbon_Timestamp(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Timestamp()
	}
}

func BenchmarkCarbon_TimestampMilli(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimestampMilli()
	}
}

func BenchmarkCarbon_TimestampMicro(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimestampMicro()
	}
}

func BenchmarkCarbon_TimestampNano(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.TimestampNano()
	}
}

func BenchmarkCarbon_Timezone(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Timezone()
	}
}

func BenchmarkCarbon_ZoneName(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ZoneName()
	}
}

func BenchmarkCarbon_ZoneOffset(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ZoneOffset()
	}
}

func BenchmarkCarbon_Locale(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Locale()
	}
}

func BenchmarkCarbon_WeekStartsAt(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.WeekStartsAt()
	}
}

func BenchmarkCarbon_CurrentLayout(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.CurrentLayout()
	}
}

func BenchmarkCarbon_Age(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Age()
	}
}
