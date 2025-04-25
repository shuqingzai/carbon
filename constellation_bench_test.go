package carbon

import (
	"testing"
)

func BenchmarkCarbon_Constellation(b *testing.B) {
	c := Now()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Constellation()
	}
}

func BenchmarkCarbon_IsAries(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-03-21")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsAries()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsAries()
		}
	})
}

func BenchmarkCarbon_IsTaurus(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-04-20")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsTaurus()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsTaurus()
		}
	})
}

func BenchmarkCarbon_IsGemini(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-05-21")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsGemini()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsGemini()
		}
	})
}

func BenchmarkCarbon_IsCancer(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-06-22")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsCancer()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsCancer()
		}
	})
}

func BenchmarkCarbon_IsLeo(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-07-23")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsLeo()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsLeo()
		}
	})
}

func BenchmarkCarbon_IsVirgo(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-08-23")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsVirgo()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsVirgo()
		}
	})
}

func BenchmarkCarbon_IsLibra(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-09-23")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsLibra()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsLibra()
		}
	})
}

func BenchmarkCarbon_IsScorpio(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-10-24")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsScorpio()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsScorpio()
		}
	})
}

func BenchmarkCarbon_IsSagittarius(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-11-23")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsSagittarius()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsSagittarius()
		}
	})
}

func BenchmarkCarbon_IsCapricorn(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-12-22")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsCapricorn()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsCapricorn()
		}
	})
}

func BenchmarkCarbon_IsAquarius(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-01-20")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsAquarius()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsAquarius()
		}
	})
}

func BenchmarkCarbon_IsPisces(b *testing.B) {
	b.Run("true", func(b *testing.B) {
		c := Parse("2020-02-19")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsPisces()
		}
	})

	b.Run("false", func(b *testing.B) {
		c := Parse("2020-08-05")
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c.IsPisces()
		}
	})
}
