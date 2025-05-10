package carbon

import (
	"testing"
)

func BenchmarkCarbon_GoString(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.GoString()
	}
}

func BenchmarkCarbon_ToString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToString(PRC)
		}
	})
}

func BenchmarkCarbon_ToMonthString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToMonthString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToMonthString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortMonthString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortMonthString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortMonthString(PRC)
		}
	})
}

func BenchmarkCarbon_ToWeekString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToWeekString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToWeekString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortWeekString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortWeekString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortWeekString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDayDateTimeString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDayDateTimeString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDayDateTimeString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateTimeString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateTimeMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateTimeMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateTimeNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateTimeNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateTimeString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateTimeMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateTimeMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateTimeNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateTimeNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToDateNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToDateNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortDateNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortDateNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToTimeString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeString(PRC)
		}
	})
}

func BenchmarkCarbon_ToTimeMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToTimeMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToTimeNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToTimeNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortTimeString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortTimeMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortTimeMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToShortTimeNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToShortTimeNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToAtomString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToAtomString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToAtomString(PRC)
		}
	})
}

func BenchmarkCarbon_ToAnsicString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToAnsicString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToAnsicString(PRC)
		}
	})
}

func BenchmarkCarbon_ToCookieString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToCookieString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToCookieString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRssString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRssString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRssString(PRC)
		}
	})
}

func BenchmarkCarbon_ToW3cString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToW3cString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToW3cString(PRC)
		}
	})
}

func BenchmarkCarbon_ToUnixDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToUnixDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToUnixDateString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRubyDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRubyDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRubyDateString(PRC)
		}
	})
}

func BenchmarkCarbon_ToKitchenString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToKitchenString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToKitchenString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601String(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601MilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601MilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601MilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601NanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601NanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601NanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601ZuluString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601ZuluMilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluMilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluMilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601ZuluMicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluMicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluMicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToIso8601ZuluNanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluNanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToIso8601ZuluNanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc822String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc822String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc822String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc822zString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc822zString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc822zString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc850String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc850String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc850String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc1036String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1036String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1036String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc1123String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1123String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1123String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc1123zString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1123zString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc1123zString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc2822String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc2822String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc2822String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc3339String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339String(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc3339MilliString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339MilliString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339MilliString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc3339MicroString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339MicroString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339MicroString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc3339NanoString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339NanoString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc3339NanoString(PRC)
		}
	})
}

func BenchmarkCarbon_ToRfc7231String(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc7231String()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToRfc7231String(PRC)
		}
	})
}

func BenchmarkCarbon_ToFormattedDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToFormattedDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToFormattedDateString(PRC)
		}
	})
}

func BenchmarkCarbon_ToFormattedDayDateString(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToFormattedDayDateString()
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.ToFormattedDayDateString(PRC)
		}
	})
}

func BenchmarkCarbon_Layout(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Layout(DateTimeLayout)
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Layout(DateTimeLayout, PRC)
		}
	})
}

func BenchmarkCarbon_Format(b *testing.B) {
	c := Now()

	b.Run("without timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Format(DateTimeLayout)
		}
	})

	b.Run("with timezone", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.Format(DateTimeLayout, PRC)
		}
	})
}
