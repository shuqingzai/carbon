package carbon

import (
	"testing"
)

func BenchmarkSetDefault(b *testing.B) {
	defer ResetDefault()

	d := Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "zh-CN",
		WeekStartsAt: Monday,
		WeekendDays: []string{
			Saturday, Sunday,
		},
	}

	for n := 0; n < b.N; n++ {
		SetDefault(d)
	}
}

func BenchmarkResetDefault(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ResetDefault()
	}
}
