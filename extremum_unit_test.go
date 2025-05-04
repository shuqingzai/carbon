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
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Closest(Now(), Now()).ToString())
		s.Empty(Now().Closest(c, Now()).ToString())
		s.Empty(Now().Closest(Now(), c).ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2020-08-05 00:00:00 +0000 UTC", c.Closest(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(c, Parse("2020-08-06")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-06"), c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Empty(c.Closest(Now(), Now()).ToString())
		s.Empty(Now().Closest(c, Now()).ToString())
		s.Empty(Now().Closest(Now(), c).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-05"), Parse("2020-08-05")).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-06"), Parse("2020-08-06")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Closest(Parse("2020-08-06"), Parse("2020-08-07")).ToString())
	})
}

func (s *ExtremumSuite) TestCarbon_Farthest() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.Farthest(Now(), Now()).ToString())
		s.Empty(Now().Farthest(c, Now()).ToString())
		s.Empty(Now().Farthest(Now(), c).ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal("2020-08-06 00:00:00 +0000 UTC", c.Farthest(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(c, Parse("2020-08-06")).ToString())
		s.Equal("0001-01-01 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-06"), c).ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Empty(c.Farthest(Now(), Now()).ToString())
		s.Empty(Now().Farthest(c, Now()).ToString())
		s.Empty(Now().Farthest(Now(), c).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-05"), Parse("2020-08-05")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-05"), Parse("2020-08-06")).ToString())
		s.Equal("2020-08-06 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-06"), Parse("2020-08-06")).ToString())
		s.Equal("2020-08-07 00:00:00 +0000 UTC", Parse("2020-08-05").Farthest(Parse("2020-08-06"), Parse("2020-08-07")).ToString())
	})
}
