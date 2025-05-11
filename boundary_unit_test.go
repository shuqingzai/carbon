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
		c := NewCarbon().StartOfCentury()
		s.Nil(c.Error)
		s.Equal("0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfCentury()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfCentury()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfCentury()
		s.Nil(c.Error)
		s.Equal("0099-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfCentury()
		s.Nil(c.Error)
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfCentury()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		s.Empty(c.StartOfDecade().ToString())
	})

	s.Run("zero carbon", func() {
		c := NewCarbon().StartOfDecade()
		s.Nil(c.Error)
		s.Equal("0000-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfDecade()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfDecade()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfDecade()
		s.Nil(c.Error)
		s.Equal("0009-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfDecade()
		s.Nil(c.Error)
		s.Empty(c.String())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfDecade()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfYear()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfYear()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfYear()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfYear()
		s.Nil(c.Error)
		s.Equal("0001-12-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfYear()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfYear()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfQuarter()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfQuarter()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfQuarter()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfQuarter()
		s.Nil(c.Error)
		s.Equal("0001-03-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfQuarter()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfQuarter()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfMonth()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfMonth()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfMonth()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfMonth()
		s.Nil(c.Error)
		s.Equal("0001-01-31 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfMonth()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfMonth()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfWeek()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfWeek()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfWeek()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfWeek()
		s.Nil(c.Error)
		s.Equal("0001-01-07 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfWeek()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfWeek()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfDay()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfDay()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfDay()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfDay()
		s.Nil(c.Error)
		s.Equal("0001-01-01 23:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfDay()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfDay()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfHour()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfHour()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfHour()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfHour()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:59:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfHour()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfHour()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfMinute()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfMinute()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfMinute()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfMinute()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:59.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfMinute()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfMinute()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().StartOfSecond()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").StartOfSecond()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").StartOfSecond()
		s.Error(c.Error)
		s.Empty(c.ToString())
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
		c := NewCarbon().EndOfSecond()
		s.Nil(c.Error)
		s.Equal("0001-01-01 00:00:00.999999999 +0000 UTC", c.ToString())
	})

	s.Run("empty carbon", func() {
		c := Parse("").EndOfSecond()
		s.Nil(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("error carbon", func() {
		c := Parse("xxx").EndOfSecond()
		s.Error(c.Error)
		s.Empty(c.ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 00:00:00.999999999 +0000 UTC", Parse("2020-01-01 00:00:00").EndOfSecond().ToString())
		s.Equal("2020-08-15 12:30:30.999999999 +0000 UTC", Parse("2020-08-15 12:30:30.66666").EndOfSecond().ToString())
		s.Equal("2020-12-31 23:59:59.999999999 +0000 UTC", Parse("2020-12-31 23:59:59.999999999").EndOfSecond().ToString())
	})
}
