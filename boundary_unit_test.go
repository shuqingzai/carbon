package carbon

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoundarySuite struct {
	suite.Suite
}

func TestBoundarySuite(t *testing.T) {
	suite.Run(t, new(BoundarySuite))
}

func (s *BoundarySuite) TestCarbon_StartOfCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfCentury())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfCentury().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfCentury().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfCentury().ToString())
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfCentury().ToString())
		s.Equal("2000-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfCentury().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfCentury())
	})

	s.Run("zero carbon", func() {
		s.Equal("0099-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfCentury().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfCentury().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfCentury().ToString())
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfCentury().ToString())
		s.Equal("2099-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfCentury().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfDecade())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDecade().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfDecade().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDecade().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDecade().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDecade().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfDecade())
	})

	s.Run("zero carbon", func() {
		s.Equal("0009-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDecade().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfDecade().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDecade().ToString())
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDecade().ToString())
		s.Equal("2029-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDecade().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfYear())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfYear().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfYear().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfYear().ToString())
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfYear().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfYear())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfYear().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfYear().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfYear().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfYear().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfQuarter())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfQuarter().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfQuarter().ToString())
		s.Equal("2020-07-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfQuarter().ToString())
		s.Equal("2020-10-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfQuarter().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfQuarter())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-03-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfQuarter().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-03-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfQuarter().ToString())
		s.Equal("2020-09-30 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfQuarter().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfQuarter().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfMonth())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfMonth().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMonth().ToString())
		s.Equal("2020-08-01 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMonth().ToString())
		s.Equal("2020-12-01 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMonth().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfMonth())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfMonth().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-31 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMonth().ToString())
		s.Equal("2020-08-31 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMonth().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMonth().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfWeek())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfWeek().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-30 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfWeek().ToString())
		s.Equal("2020-08-10 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfWeek().ToString())
		s.Equal("2020-12-28 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfWeek().ToString())
		s.Equal("2025-04-07 00:00:00 +0000 UTC", Parse("2025-04-07 00:00:00").StartOfWeek().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfWeek())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-07 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfWeek().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-05 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfWeek().ToString())
		s.Equal("2020-08-16 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfWeek().ToString())
		s.Equal("2021-01-03 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfWeek().ToString())
		s.Equal("2025-04-13 23:59:59.999999999 +0000 UTC", Parse("2025-04-13 00:00:00").EndOfWeek().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfDay())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfDay().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfDay().ToString())
		s.Equal("2020-08-15 00:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfDay().ToString())
		s.Equal("2020-12-31 00:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfDay().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfDay())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfDay().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 23:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfDay().ToString())
		s.Equal("2020-08-15 23:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfDay().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfDay().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfHour())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfHour().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfHour().ToString())
		s.Equal("2020-08-15 12:00:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfHour().ToString())
		s.Equal("2020-12-31 23:00:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfHour().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfHour())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:59:59.999999999 +0000 UTC", NewCarbon().EndOfHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfHour().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:59:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfHour().ToString())
		s.Equal("2020-08-15 12:59:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfHour().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfHour().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfMinute())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfMinute().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfMinute().ToString())
		s.Equal("2020-08-15 12:30:00 +0000 UTC", Parse("2020-08-15 12:30:30").StartOfMinute().ToString())
		s.Equal("2020-12-31 23:59:00 +0000 UTC", Parse("2020-12-31 23:59:59").StartOfMinute().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfMinute())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:59.999999999 +0000 UTC", NewCarbon().EndOfMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfMinute().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:59.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfMinute().ToString())
		s.Equal("2020-08-15 12:30:59.999999999 +0000 UTC", Parse("2020-08-15 12:30:30").EndOfMinute().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59").EndOfMinute().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_StartOfSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.StartOfSecond())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfSecond().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00 +0000 UTC", Parse("2020-01-01 00:00:00").StartOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30 +0000 UTC", Parse("2020-08-15 12:30:30.66666").StartOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").StartOfSecond().ToString())
	})
}

func (s *BoundarySuite) TestCarbon_EndOfSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Nil(c.EndOfSecond())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.999999999 +0000 UTC", NewCarbon().EndOfSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfSecond().Error)
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
