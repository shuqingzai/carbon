package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CalendarSuite struct {
	suite.Suite
}

func TestCalendarSuite(t *testing.T) {
	suite.Run(t, new(CalendarSuite))
}

func (s *CalendarSuite) TestCarbon_Julian() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Julian())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon()
		s.Equal(1.7214235e+06, c.Julian().JD())
		s.Equal(float64(-678577), c.Julian().MJD())
	})

	s.Run("empty carbon", func() {
		c := Parse("")
		s.Zero(c.Julian().JD())
		s.Zero(c.Julian().MJD())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx")
		s.Zero(c.Julian().JD())
		s.Zero(c.Julian().MJD())
	})

	s.Run("valid carbon", func() {
		c := Parse("2024-01-23 13:14:15")

		s.Equal(2460333.051563, c.Julian().JD())
		s.Equal(60332.551563, c.Julian().MJD())

		s.Equal(2460333.0516, c.Julian().JD(4))
		s.Equal(60332.5516, c.Julian().MJD(4))

		s.Equal(2460333.0515625, c.Julian().JD(7))
		s.Equal(60332.5515625, c.Julian().MJD(7))
	})
}

func (s *CalendarSuite) TestCreateFromJulian() {
	s.Run("error julian", func() {
		s.Equal("-4712-01-01 12:00:00", CreateFromJulian(0).String())
		s.Equal("-4712-01-01 12:00:00", CreateFromJulian(-1).String())
	})

	s.Run("valid julian", func() {
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.051563).ToString())
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.551563).ToString())

		s.Equal("2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(2460333.0516).ToString())
		s.Equal("2024-01-23 13:14:18 +0000 UTC", CreateFromJulian(60332.5516).ToString())

		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(2460333.0515625).ToString())
		s.Equal("2024-01-23 13:14:15 +0000 UTC", CreateFromJulian(60332.5515625).ToString())
	})
}

func (s *CalendarSuite) TestCarbon_Lunar() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Lunar())
	})

	s.Run("zero carbon", func() {
		s.Empty(NewCarbon().Lunar().String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Lunar().String())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").Lunar().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2023-12-08", Parse("2024-01-18", PRC).Lunar().String())
		s.Equal("2023-12-11", Parse("2024-01-21", PRC).Lunar().String())
		s.Equal("2023-12-14", Parse("2024-01-24", PRC).Lunar().String())
	})
}

func (s *CalendarSuite) TestCreateFromLunar() {
	s.Run("error lunar", func() {
		s.Empty(CreateFromLunar(2200, 12, 14, false).ToString())
	})

	s.Run("valid lunar", func() {
		s.Equal("2024-01-21 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 11, false).ToString(PRC))
		s.Equal("2024-01-18 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 8, false).ToString(PRC))
		s.Equal("2024-01-24 00:00:00 +0800 CST", CreateFromLunar(2023, 12, 14, false).ToString(PRC))
	})
}

func (s *CalendarSuite) TestCarbon_Persian() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.Persian())
	})

	s.Run("zero carbon", func() {
		s.Empty(NewCarbon().Persian().String())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").Persian().String())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").Persian().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("1178-10-11", Parse("1800-01-01 00:00:00").Persian().String())
		s.Equal("1399-05-15", Parse("2020-08-05 13:14:15").Persian().String())
		s.Equal("1402-10-11", Parse("2024-01-01 00:00:00").Persian().String())
	})
}

func (s *CalendarSuite) TestCreateFromPersian() {
	s.Run("error persian", func() {
		s.Empty(CreateFromPersian(9999, 12, 14).ToDateTimeString())
	})

	s.Run("valid persian", func() {
		s.Equal("1800-01-01 00:00:00", CreateFromPersian(1178, 10, 11).ToDateTimeString())
		s.Equal("2024-01-01 00:00:00", CreateFromPersian(1402, 10, 11).ToDateTimeString())
		s.Equal("2024-08-05 00:00:00", CreateFromPersian(1403, 5, 15).ToDateTimeString())
	})
}
