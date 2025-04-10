package carbon

import (
	"testing"
)

func BenchmarkCarbon_String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.String()
	}
}

func BenchmarkCarbon_GoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.GoString()
	}
}

func BenchmarkCarbon_ToString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToString()
	}
}

func BenchmarkCarbon_ToMonthString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToMonthString()
	}
}

func BenchmarkCarbon_ToShortMonthString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortMonthString()
	}
}

func BenchmarkCarbon_ToWeekString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToWeekString()
	}
}

func BenchmarkCarbon_ToShortWeekString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortWeekString()
	}
}

func BenchmarkCarbon_ToDayDateTimeString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDayDateTimeString()
	}
}

func BenchmarkCarbon_ToDateTimeString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateTimeString()
	}
}

func BenchmarkCarbon_ToDateTimeMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateTimeMilliString()
	}
}

func BenchmarkCarbon_ToDateTimeMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateTimeMicroString()
	}
}

func BenchmarkCarbon_ToDateTimeNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateTimeNanoString()
	}
}

func BenchmarkCarbon_ToShortDateTimeString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateTimeString()
	}
}

func BenchmarkCarbon_ToShortDateTimeMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateTimeMilliString()
	}
}

func BenchmarkCarbon_ToShortDateTimeMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateTimeMicroString()
	}
}

func BenchmarkCarbon_ToShortDateTimeNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateTimeNanoString()
	}
}

func BenchmarkCarbon_ToDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateString()
	}
}

func BenchmarkCarbon_ToDateMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateMilliString()
	}
}

func BenchmarkCarbon_ToDateMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateMicroString()
	}
}

func BenchmarkCarbon_ToDateNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToDateNanoString()
	}
}

func BenchmarkCarbon_ToShortDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateString()
	}
}

func BenchmarkCarbon_ToShortDateMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateMilliString()
	}
}

func BenchmarkCarbon_ToShortDateMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateMicroString()
	}
}

func BenchmarkCarbon_ToShortDateNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortDateNanoString()
	}
}

func BenchmarkCarbon_ToTimeString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToTimeString()
	}
}

func BenchmarkCarbon_ToTimeMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToTimeMilliString()
	}
}

func BenchmarkCarbon_ToTimeMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToTimeMicroString()
	}
}

func BenchmarkCarbon_ToTimeNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToTimeNanoString()
	}
}

func BenchmarkCarbon_ToShortTimeString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortTimeString()
	}
}

func BenchmarkCarbon_ToShortTimeMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortTimeMilliString()
	}
}

func BenchmarkCarbon_ToShortTimeMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortTimeMicroString()
	}
}

func BenchmarkCarbon_ToShortTimeNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToShortTimeNanoString()
	}
}

func BenchmarkCarbon_ToAtomString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToAtomString()
	}
}

func BenchmarkCarbon_ToAnsicString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToAnsicString()
	}
}

func BenchmarkCarbon_ToCookieString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToCookieString()
	}
}

func BenchmarkCarbon_ToRssString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRssString()
	}
}

func BenchmarkCarbon_ToW3cString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToW3cString()
	}
}

func BenchmarkCarbon_ToUnixDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToUnixDateString()
	}
}

func BenchmarkCarbon_ToRubyDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRubyDateString()
	}
}

func BenchmarkCarbon_ToKitchenString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToKitchenString()
	}
}

func BenchmarkCarbon_ToIso8601String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601String()
	}
}

func BenchmarkCarbon_ToIso8601MilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601MilliString()
	}
}

func BenchmarkCarbon_ToIso8601NanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601NanoString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601ZuluString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluMilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601ZuluMilliString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluMicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601ZuluMicroString()
	}
}

func BenchmarkCarbon_ToIso8601ZuluNanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToIso8601ZuluNanoString()
	}
}

func BenchmarkCarbon_ToRfc822String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc822String()
	}
}

func BenchmarkCarbon_ToRfc822zString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc822zString()
	}
}

func BenchmarkCarbon_ToRfc850String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc850String()
	}
}

func BenchmarkCarbon_ToRfc1036String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc1036String()
	}
}

func BenchmarkCarbon_ToRfc1123String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc1123String()
	}
}

func BenchmarkCarbon_ToRfc1123zString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc1123zString()
	}
}

func BenchmarkCarbon_ToRfc2822String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc2822String()
	}
}

func BenchmarkCarbon_ToRfc3339String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc3339String()
	}
}

func BenchmarkCarbon_ToRfc3339MilliString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc3339MilliString()
	}
}

func BenchmarkCarbon_ToRfc3339MicroString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc3339MicroString()
	}
}

func BenchmarkCarbon_ToRfc3339NanoString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc3339NanoString()
	}
}

func BenchmarkCarbon_ToRfc7231String(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToRfc7231String()
	}
}

func BenchmarkCarbon_ToFormattedDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToFormattedDateString()
	}
}

func BenchmarkCarbon_ToFormattedDayDateString(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.ToFormattedDayDateString()
	}
}

func BenchmarkCarbon_Layout(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Layout(DateTimeLayout)
	}
}

func BenchmarkCarbon_Format(b *testing.B) {
	c := Now()
	for n := 0; n < b.N; n++ {
		c.Format(DateTimeFormat)
	}
}
