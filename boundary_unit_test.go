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
		s.Empty(c.StartOfCentury().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfCentury().Error)
		s.Equal("0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfCentury().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfCentury().Error)
		s.Empty(Parse("").StartOfCentury().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfCentury().Error)
		s.Empty(Parse("xxx").StartOfCentury().ToString())
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
		s.Empty(c.EndOfCentury().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfCentury().Error)
		s.Equal("0099-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfCentury().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(NewCarbon().EndOfCentury().Error)
		s.Empty(Parse("").EndOfCentury().String())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfCentury().Error)
		s.Empty(Parse("xxx").EndOfCentury().ToString())
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
		s.Nil(c.EndOfDecade())
		s.Empty(c.EndOfDecade().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfDecade().Error)
		s.Equal("0000-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDecade().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(NewCarbon().StartOfDecade().Error)
		s.Empty(Parse("").StartOfDecade().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfDecade().Error)
		s.Empty(Parse("xxx").StartOfDecade().ToString())
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
		s.Empty(c.EndOfDecade().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfDecade().Error)
		s.Equal("0009-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDecade().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(NewCarbon().EndOfDecade().Error)
		s.Empty(Parse("").EndOfDecade().String())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfDecade().Error)
		s.Empty(Parse("xxx").EndOfDecade().ToString())
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
		s.Empty(c.StartOfYear().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfYear().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfYear().Error)
		s.Empty(Parse("").StartOfYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfYear().Error)
		s.Empty(Parse("xxx").StartOfYear().ToString())
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
		s.Empty(c.EndOfYear().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfYear().Error)
		s.Equal("0001-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfYear().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfYear().Error)
		s.Empty(Parse("").EndOfYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfYear().Error)
		s.Empty(Parse("xxx").EndOfYear().ToString())
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
		s.Empty(c.StartOfQuarter().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfQuarter().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfQuarter().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfQuarter().Error)
		s.Empty(Parse("").StartOfQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfQuarter().Error)
		s.Empty(Parse("xxx").StartOfQuarter().ToString())
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
		s.Empty(c.EndOfQuarter().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfQuarter().Error)
		s.Equal("0001-03-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfQuarter().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfQuarter().Error)
		s.Empty(Parse("").EndOfQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfQuarter().Error)
		s.Empty(Parse("xxx").EndOfQuarter().ToString())
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
		s.Empty(c.StartOfMonth().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfMonth().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMonth().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfMonth().Error)
		s.Empty(Parse("").StartOfMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfMonth().Error)
		s.Empty(Parse("xxx").StartOfMonth().ToString())
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
		s.Empty(c.EndOfMonth().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfMonth().Error)
		s.Equal("0001-01-31 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfMonth().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfMonth().Error)
		s.Empty(Parse("").EndOfMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfMonth().Error)
		s.Empty(Parse("xxx").EndOfMonth().ToString())
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
		s.Empty(c.StartOfWeek().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfWeek().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfWeek().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfWeek().Error)
		s.Empty(Parse("").StartOfWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfWeek().Error)
		s.Empty(Parse("xxx").StartOfWeek().ToString())
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
		s.Empty(c.EndOfWeek().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfWeek().Error)
		s.Equal("0001-01-07 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfWeek().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfWeek().Error)
		s.Empty(Parse("").EndOfWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfWeek().Error)
		s.Empty(Parse("xxx").EndOfWeek().ToString())
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
		s.Empty(c.StartOfDay().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfDay().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfDay().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfDay().Error)
		s.Empty(Parse("").StartOfDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfDay().Error)
		s.Empty(Parse("xxx").StartOfDay().ToString())
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
		s.Empty(c.EndOfDay().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfDay().Error)
		s.Equal("0001-01-01 23:59:59.999999999 +0000 UTC", NewCarbon().EndOfDay().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfDay().Error)
		s.Empty(Parse("").EndOfDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfDay().Error)
		s.Empty(Parse("xxx").EndOfDay().ToString())
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
		s.Empty(c.StartOfHour().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfHour().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfHour().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfHour().Error)
		s.Empty(Parse("").StartOfHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfHour().Error)
		s.Empty(Parse("xxx").StartOfHour().ToString())
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
		s.Empty(c.EndOfHour().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfHour().Error)
		s.Equal("0001-01-01 00:59:59.999999999 +0000 UTC", NewCarbon().EndOfHour().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfHour().Error)
		s.Empty(Parse("").EndOfHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfHour().Error)
		s.Empty(Parse("xxx").EndOfHour().ToString())
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
		s.Empty(c.StartOfMinute().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfMinute().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfMinute().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfMinute().Error)
		s.Empty(Parse("").StartOfMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfMinute().Error)
		s.Empty(Parse("xxx").StartOfMinute().ToString())
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
		s.Empty(c.EndOfMinute().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfMinute().Error)
		s.Equal("0001-01-01 00:00:59.999999999 +0000 UTC", NewCarbon().EndOfMinute().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfMinute().Error)
		s.Empty(Parse("").EndOfMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfMinute().Error)
		s.Empty(Parse("xxx").EndOfMinute().ToString())
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
		s.Empty(c.StartOfSecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().StartOfSecond().Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", NewCarbon().StartOfSecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").StartOfSecond().Error)
		s.Empty(Parse("").StartOfSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").StartOfSecond().Error)
		s.Empty(Parse("xxx").StartOfSecond().ToString())
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
		s.Empty(c.EndOfSecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Nil(NewCarbon().EndOfSecond().Error)
		s.Equal("0001-01-01 00:00:00.999999999 +0000 UTC", NewCarbon().EndOfSecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Nil(Parse("").EndOfSecond().Error)
		s.Empty(Parse("").EndOfSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Error(Parse("xxx").EndOfSecond().Error)
		s.Empty(Parse("xxx").EndOfSecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
