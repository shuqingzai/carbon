package carbon

import (
	"testing"
)

func BenchmarkLanguage_SetLocale(b *testing.B) {
	lang := NewLanguage()
	for n := 0; n < b.N; n++ {
		lang.SetLocale("en")
	}
}

func BenchmarkLanguage_SetResources(b *testing.B) {
	resources := map[string]string{
		"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
		"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
	}
	lang := NewLanguage()
	for n := 0; n < b.N; n++ {
		lang.SetResources(resources)
	}
}

func BenchmarkLanguage_translate(b *testing.B) {
	lang := NewLanguage()
	lang.SetLocale("en")
	for n := 0; n < b.N; n++ {
		lang.translate("month", 1)
	}
}
