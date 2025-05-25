package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExtremumSuite struct {
	suite.Suite
}

func TestExtremumSuite(t *testing.T) {
	suite.Run(t, new(ExtremumSuite))
}

func (s *ExtremumSuite) TestZeroValue() {
	s.Equal("0001-01-01 00:00:00 +0000 UTC", ZeroValue().ToString())
	s.Equal("January", ZeroValue().ToMonthString())
	s.Equal("一月", ZeroValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestEpochValue() {
	s.Equal("1970-01-01 00:00:00 +0000 UTC", EpochValue().ToString())
	s.Equal("January", EpochValue().ToMonthString())
	s.Equal("一月", EpochValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMaxValue() {
	s.Equal("9999-12-31 23:59:59.999999999 +0000 UTC", MaxValue().ToString())
	s.Equal("December", MaxValue().ToMonthString())
	s.Equal("十二月", MaxValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMinValue() {
	s.Equal("0001-01-01 00:00:00 +0000 UTC", MinValue().ToString())
	s.Equal("January", MinValue().ToMonthString())
	s.Equal("一月", MinValue().SetLocale("zh-CN").ToMonthString())
}

func (s *ExtremumSuite) TestMaxDuration() {
	s.Equal(9.223372036854776e+09, MaxDuration().Seconds())
}

func (s *ExtremumSuite) TestMinDuration() {
	s.Equal(-9.223372036854776e+09, MinDuration().Seconds())
}

func (s *ExtremumSuite) TestMax() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(Max(c, Now()))
		s.Nil(Max(Now(), c))
		s.Nil(Max(c, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Max(c, Parse("2020-08-06")).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Max(Parse("2020-08-05"), c).ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Max(c, c).ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Empty(Max(c, Now()).ToString())
		s.Empty(Max(Now(), c).ToString())
		s.Empty(Max(c, c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Error(Max(c, Now()).Error)
		s.Error(Max(Now(), c).Error)
		s.Error(Max(c, c).Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Max(Parse("2020-08-05")).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Max(Parse("2020-08-05"), Parse("2020-08-05")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Max(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
	})
}

func (s *ExtremumSuite) TestMin() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(Min(c, Now()))
		s.Nil(Min(Now(), c))
		s.Nil(Min(c, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Min(c, Parse("2020-08-06")).ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Min(Parse("2020-08-05"), c).ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Min(c, c).ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Empty(Min(c, Now()).ToString())
		s.Empty(Min(Now(), c).ToString())
		s.Empty(Min(c, c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Error(Min(c, Now()).Error)
		s.Error(Min(Now(), c).Error)
		s.Error(Min(c, c).Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Min(Parse("2020-08-05")).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Min(Parse("2020-08-05"), Parse("2020-08-05")).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Min(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
	})
}

func (s *ExtremumSuite) TestCarbon_Closest() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c1, c.Closest(c1, c2, c3))
		s.Equal(c2, c1.Closest(c, c2, c3))
		s.Equal(c1, c2.Closest(c1, c, c3))
		s.Equal(c2, c3.Closest(c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Equal(c, c.Closest(c1, c2, c3))
		s.Equal(c, c1.Closest(c, c2, c3))
		s.Equal(c, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")
		s.Equal(c, c.Closest(c))
		s.Equal(c1, c.Closest(c1))
		s.Equal(c2, c.Closest(c1, c2))
		s.Equal(c3, c.Closest(c1, c3))
		s.Equal(c3, c.Closest(c1, c2, c3))
		s.Equal(c2, c1.Closest(c, c2, c3))
		s.Equal(c1, c2.Closest(c1, c, c3))
		s.Equal(c, c3.Closest(c1, c2, c))
	})
}

func (s *ExtremumSuite) TestCarbon_Farthest() {
	c1 := Parse("2020-08-01")
	c2 := Parse("2020-08-03")
	c3 := Parse("2020-08-06")

	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(c3, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Equal(c, c.Farthest(c1, c2, c3))
		s.Equal(c, c1.Farthest(c, c2, c3))
		s.Equal(c, c2.Farthest(c1, c, c3))
		s.Equal(c, c3.Farthest(c1, c2, c))
	})

	s.Run("valid carbon", func() {
		c := Parse("2020-08-05")
		s.Equal(c, c.Farthest(c))
		s.Equal(c1, c.Farthest(c1))
		s.Equal(c1, c.Farthest(c1, c2))
		s.Equal(c1, c.Farthest(c1, c3))
		s.Equal(c1, c.Farthest(c1, c2, c3))
		s.Equal(c3, c1.Farthest(c, c2, c3))
		s.Equal(c3, c2.Farthest(c1, c, c3))
		s.Equal(c1, c3.Farthest(c1, c2, c))
	})
}
