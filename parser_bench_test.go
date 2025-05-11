package carbon

import (
	"testing"
)

func BenchmarkParse(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Parse("2020-08-05 01:02:03")
	}
}

func BenchmarkParseByLayout(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParseByLayout("2020-08-05 13:14:15", DateTimeLayout)
	}
}

func BenchmarkParseByFormat(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParseByFormat("2020-08-05 13:14:15", DateTimeFormat)
	}
}

func BenchmarkParseByLayouts(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParseByLayouts("2020-08-05 13:14:15", []string{DateLayout, DateTimeLayout})
	}
}

func BenchmarkParseByFormats(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		ParseByFormats("2020-08-05 13:14:15", []string{DateFormat, DateTimeFormat})
	}
}
