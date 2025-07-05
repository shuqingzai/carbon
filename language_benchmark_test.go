package carbon

import (
	"testing"
)

func BenchmarkLanguage_Copy(b *testing.B) {
	lang := NewLanguage()
	lang.SetLocale("en")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			lang.Copy()
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				_ = lang.Copy()
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
				lang.Copy()
			}
		})
	})
}

func BenchmarkLanguage_SetLocale(b *testing.B) {
	lang := NewLanguage()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			lang.SetLocale("en")
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				lang.SetLocale("en")
			}()
		}
	})

	b.Run("parallel", func(b *testing.B) {
		b.ResetTimer()
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				lang.SetLocale("en")
			}
		})
	})
}

func BenchmarkLanguage_SetResources(b *testing.B) {
	resources := map[string]string{
		"months":       "Ⅰ月|Ⅱ月|Ⅲ月|Ⅳ月|Ⅴ月|Ⅵ月|Ⅶ月|Ⅷ月|Ⅸ月|Ⅹ月|Ⅺ月|Ⅻ月",
		"short_months": "Ⅰ|Ⅱ|Ⅲ|Ⅳ|Ⅴ|Ⅵ|Ⅶ|Ⅷ|Ⅸ|Ⅹ|Ⅺ|Ⅻ",
		"weeks":        "星期日|星期一|星期二|星期三|星期四|星期五|星期六",
		"short_weeks":  "周日|周一|周二|周三|周四|周五|周六",
		"seasons":      "春季|夏季|秋季|冬季",
		"year":         "1年|%d年",
		"month":        "1个月|%d个月",
		"week":         "1周|%d周",
		"day":          "1天|%d天",
		"hour":         "1小时|%d小时",
		"minute":       "1分钟|%d分钟",
		"second":       "1秒|%d秒",
		"now":          "刚刚",
		"ago":          "%s前",
		"from_now":     "%s后",
		"before":       "%s前",
		"after":        "%s后",
	}
	lang := NewLanguage()

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			lang.SetResources(resources)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				lang.SetResources(resources)
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
				lang.SetResources(resources)
			}
		})
	})
}

func BenchmarkLanguage_translate(b *testing.B) {
	lang := NewLanguage()
	lang.SetLocale("en")

	b.Run("sequential", func(b *testing.B) {
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			lang.translate("month", 1)
		}
	})

	b.Run("concurrent", func(b *testing.B) {
		done := make(chan bool, b.N)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			go func() {
				lang.translate("month", 1)
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
				lang.translate("month", 1)
			}
		})
	})
}
