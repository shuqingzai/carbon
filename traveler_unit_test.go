package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type TravelerSuite struct {
	suite.Suite
}

func TestTravelerSuite(t *testing.T) {
	suite.Run(t, new(TravelerSuite))
}

func (s *TravelerSuite) TearDownTest() {
	CleanTestNow()
}

func (s *TravelerSuite) TestNow() {
	s.Run("error timezone", func() {
		s.Error(Now("xxx").Error)
		s.Empty(Now("xxx").ToString())
	})

	s.Run("valid timezone", func() {
		s.Equal(time.Now().Format(DateLayout), Now().Layout(DateLayout, Local))
		s.Equal(time.Now().In(time.UTC).Format(DateLayout), Now(UTC).Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-05", Now(UTC).Layout(DateLayout))
	})
}

func (s *TravelerSuite) TestTomorrow() {
	s.Run("error timezone", func() {
		s.Error(Tomorrow("xxx").Error)
		s.Empty(Tomorrow("xxx").ToString())
	})

	s.Run("valid timezone", func() {
		s.Equal(time.Now().Add(time.Hour*24).Format(DateLayout), Tomorrow().Layout(DateLayout, Local))
		s.Equal(time.Now().Add(time.Hour*24).In(time.UTC).Format(DateLayout), Tomorrow(UTC).Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-06", Tomorrow(UTC).Layout(DateLayout))
	})
}

func (s *TravelerSuite) TestYesterday() {
	s.Run("error timezone", func() {
		s.Error(Yesterday("xxx").Error)
		s.Empty(Yesterday("xxx").ToString())
	})

	s.Run("valid timezone", func() {
		s.Equal(time.Now().Add(time.Hour*-24).Format(DateLayout), Yesterday().Layout(DateLayout, Local))
		s.Equal(time.Now().Add(time.Hour*-24).In(time.UTC).Format(DateLayout), Yesterday(UTC).Layout(DateLayout))
	})

	s.Run("frozen time", func() {
		SetTestNow(Parse("2020-08-05"))
		s.Equal("2020-08-04", Yesterday(UTC).Layout(DateLayout))
	})
}

func (s *TravelerSuite) TestCarbon_AddDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.AddDuration("10h").ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 10:00:00 +0000 UTC", NewCarbon().AddDuration("10h").ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDuration("10h").ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDuration("10h").ToString())
	})

	s.Run("empty duration", func() {
		s.Empty(Parse("2020-08-05").AddDuration("").ToString())
	})

	s.Run("error duration", func() {
		s.Empty(Parse("2020-08-05").AddDuration("xxx").ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 23:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10h").ToString())
		s.Equal("2020-01-01 23:44:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10.5h").ToString())
		s.Equal("2020-01-01 13:24:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10m").ToString())
		s.Equal("2020-01-01 13:24:45 +0000 UTC", Parse("2020-01-01 13:14:15").AddDuration("10.5m").ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDuration() {
	s.Run("nil carbon", func() {
		var c *Carbon
		s.Empty(c.SubDuration("10h").ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 14:00:00 +0000 UTC", NewCarbon().SubDuration("10h").ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDuration("10h").ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDuration("10h").ToString())
	})

	s.Run("error duration", func() {
		s.Empty(Parse("2020-08-05").SubDuration("xxx").ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 03:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10h").ToString())
		s.Equal("2020-01-01 02:44:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5h").ToString())
		s.Equal("2020-01-01 13:04:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10m").ToString())
		s.Equal("2020-01-01 13:03:45 +0000 UTC", Parse("2020-01-01 13:14:15").SubDuration("10.5m").ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturies() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddCenturies(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0201-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturies(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddCenturies(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddCenturies(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(0).ToString())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(1).ToString())
		s.Equal("2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturies(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturiesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddCenturiesNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0201-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturiesNoOverflow(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddCenturiesNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddCenturiesNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(0).ToString())
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(1).ToString())
		s.Equal("2220-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturiesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddCentury().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0101-01-01 00:00:00 +0000 UTC", NewCarbon().AddCentury().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddCentury().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddCentury().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCentury().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddCenturyNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddCenturyNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0101-01-01 00:00:00 +0000 UTC", NewCarbon().AddCenturyNoOverflow().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddCenturyNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddCenturyNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2120-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddCenturyNoOverflow().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturies() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubCenturies(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0199-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturies(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubCenturies(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubCenturies(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(0).ToString())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(1).ToString())
		s.Equal("1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturies(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturiesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubCenturiesNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0199-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturiesNoOverflow(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubCenturiesNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubCenturiesNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(0).ToString())
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(1).ToString())
		s.Equal("1820-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturiesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCentury() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubCentury().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0099-01-01 00:00:00 +0000 UTC", NewCarbon().SubCentury().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubCentury().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubCentury().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCentury().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubCenturyNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubCenturyNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0099-01-01 00:00:00 +0000 UTC", NewCarbon().SubCenturyNoOverflow().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubCenturyNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubCenturyNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("1920-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubCenturyNoOverflow().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecades() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDecades(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0021-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecades(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDecades(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDecades(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(0).ToString())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(1).ToString())
		s.Equal("2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecades(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecadesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDecadesNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0021-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecadesNoOverflow(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDecadesNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDecadesNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(0).ToString())
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(1).ToString())
		s.Equal("2040-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDecade().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0011-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecade().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDecade().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDecade().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecade().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDecadeNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDecadeNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0011-01-01 00:00:00 +0000 UTC", NewCarbon().AddDecadeNoOverflow().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDecadeNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDecadeNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2030-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddDecadeNoOverflow().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecades() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDecades(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0019-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecades(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDecades(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDecades(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(0).ToString())
		s.Equal("2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(1).ToString())
		s.Equal("2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecades(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecadesNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDecadesNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0019-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecadesNoOverflow(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDecadesNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDecadesNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(0).ToString())
		s.Equal("2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(1).ToString())
		s.Equal("2000-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecadesNoOverflow(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecade() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDecade().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0009-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecade().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDecade().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDecade().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2010-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubDecade().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDecadeNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDecadeNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0009-01-01 00:00:00 +0000 UTC", NewCarbon().SubDecadeNoOverflow().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDecadeNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDecadeNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2010-01-01", Parse("2020-01-01").SubDecadeNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddYears(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0003-01-01", NewCarbon().AddYears(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddYears(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddYears(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddYears(0).ToDateString())
		s.Equal("2021-01-01", Parse("2020-01-01").AddYears(1).ToDateString())
		s.Equal("2022-01-01", Parse("2020-01-01").AddYears(2).ToDateString())
		s.Equal("2023-03-01", Parse("2020-02-29").AddYears(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYearsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddYearsNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0003-01-01", NewCarbon().AddYearsNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddYearsNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddYearsNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddYearsNoOverflow(0).ToDateString())
		s.Equal("2021-01-01", Parse("2020-01-01").AddYearsNoOverflow(1).ToDateString())
		s.Equal("2022-01-01", Parse("2020-01-01").AddYearsNoOverflow(2).ToDateString())
		s.Equal("2023-02-28", Parse("2020-02-29").AddYearsNoOverflow(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddYear().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0002-01-01", NewCarbon().AddYear().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddYear().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2021-01-01", Parse("2020-01-01").AddYear().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-28").AddYear().ToDateString())
		s.Equal("2021-03-01", Parse("2020-02-29").AddYear().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddYearNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddYearNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0002-01-01", NewCarbon().AddYearNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddYearNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddYearNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2021-01-01", Parse("2020-01-01").AddYearNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-28").AddYearNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-02-29").AddYearNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYears() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubYears(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0001-01-01", NewCarbon().SubYears(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubYears(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubYears(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubYears(0).ToDateString())
		s.Equal("2019-01-01", Parse("2020-01-01").SubYears(1).ToDateString())
		s.Equal("2018-01-01", Parse("2020-01-01").SubYears(2).ToDateString())
		s.Equal("2017-03-01", Parse("2020-02-29").SubYears(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYearsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubYearsNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("-0001-01-01", NewCarbon().SubYearsNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubYearsNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubYearsNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubYearsNoOverflow(0).ToDateString())
		s.Equal("2019-01-01", Parse("2020-01-01").SubYearsNoOverflow(1).ToDateString())
		s.Equal("2018-01-01", Parse("2020-01-01").SubYearsNoOverflow(2).ToDateString())
		s.Equal("2017-02-28", Parse("2020-02-29").SubYearsNoOverflow(3).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYear() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubYear().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-01-01", NewCarbon().SubYear().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubYear().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubYear().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01", Parse("2020-01-01").SubYear().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-28").SubYear().ToDateString())
		s.Equal("2019-03-01", Parse("2020-02-29").SubYear().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubYearNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubYearNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-01-01", NewCarbon().SubYearNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubYearNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubYearNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-01-01", Parse("2020-01-01").SubYearNoOverflow().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-28").SubYearNoOverflow().ToDateString())
		s.Equal("2019-02-28", Parse("2020-02-29").SubYearNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarters() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddQuarters(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-07-01", NewCarbon().AddQuarters(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddQuarters(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddQuarters(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddQuarters(0).ToDateString())
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarters(1).ToDateString())
		s.Equal("2020-07-01", Parse("2020-01-01").AddQuarters(2).ToDateString())
		s.Equal("2020-11-29", Parse("2020-02-29").AddQuarters(3).ToDateString())
		s.Equal("2021-03-03", Parse("2020-08-31").AddQuarters(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuartersNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddQuartersNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-07-01", NewCarbon().AddQuartersNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddQuartersNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddQuartersNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddQuartersNoOverflow(0).ToDateString())
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuartersNoOverflow(1).ToDateString())
		s.Equal("2020-07-01", Parse("2020-01-01").AddQuartersNoOverflow(2).ToDateString())
		s.Equal("2020-11-29", Parse("2020-02-29").AddQuartersNoOverflow(3).ToDateString())
		s.Equal("2021-02-28", Parse("2020-08-31").AddQuartersNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddQuarter().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-04-01", NewCarbon().AddQuarter().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddQuarter().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarter().ToDateString())
		s.Equal("2020-05-28", Parse("2020-02-28").AddQuarter().ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddQuarter().ToDateString())
		s.Equal("2021-03-02", Parse("2020-11-30").AddQuarter().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddQuarterNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddQuarterNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-04-01", NewCarbon().AddQuarterNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddQuarterNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddQuarterNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-04-01", Parse("2020-01-01").AddQuarterNoOverflow().ToDateString())
		s.Equal("2020-05-28", Parse("2020-02-28").AddQuarterNoOverflow().ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddQuarterNoOverflow().ToDateString())
		s.Equal("2021-02-28", Parse("2020-11-30").AddQuarterNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarters() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubQuarters(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-07-01", NewCarbon().SubQuarters(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubQuarters(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubQuarters(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubQuarters(0).ToDateString())
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarters(1).ToDateString())
		s.Equal("2019-07-01", Parse("2020-01-01").SubQuarters(2).ToDateString())
		s.Equal("2019-05-29", Parse("2020-02-29").SubQuarters(3).ToDateString())
		s.Equal("2020-03-02", Parse("2020-08-31").SubQuarters(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuartersNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubQuartersNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-07-01", NewCarbon().SubQuartersNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubQuartersNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubQuartersNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubQuartersNoOverflow(0).ToDateString())
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuartersNoOverflow(1).ToDateString())
		s.Equal("2019-07-01", Parse("2020-01-01").SubQuartersNoOverflow(2).ToDateString())
		s.Equal("2019-05-29", Parse("2020-02-29").SubQuartersNoOverflow(3).ToDateString())
		s.Equal("2020-02-29", Parse("2020-08-31").SubQuartersNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarter() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubQuarter().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-10-01", NewCarbon().SubQuarter().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubQuarter().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubQuarter().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarter().ToDateString())
		s.Equal("2019-11-28", Parse("2020-02-28").SubQuarter().ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubQuarter().ToDateString())
		s.Equal("2020-08-30", Parse("2020-11-30").SubQuarter().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubQuarterNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubQuarterNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-10-01", NewCarbon().SubQuarterNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubQuarterNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubQuarterNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-10-01", Parse("2020-01-01").SubQuarterNoOverflow().ToDateString())
		s.Equal("2019-11-28", Parse("2020-02-28").SubQuarterNoOverflow().ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubQuarterNoOverflow().ToDateString())
		s.Equal("2020-08-30", Parse("2020-11-30").SubQuarterNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMonths(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-03-01", NewCarbon().AddMonths(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMonths(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMonths(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddMonths(0).ToDateString())
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonths(1).ToDateString())
		s.Equal("2020-03-01", Parse("2020-01-01").AddMonths(2).ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddMonths(3).ToDateString())
		s.Equal("2020-10-31", Parse("2020-08-31").AddMonths(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonthsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMonthsNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-03-01", NewCarbon().AddMonthsNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMonthsNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMonthsNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddMonthsNoOverflow(0).ToDateString())
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonthsNoOverflow(1).ToDateString())
		s.Equal("2020-03-01", Parse("2020-01-01").AddMonthsNoOverflow(2).ToDateString())
		s.Equal("2020-05-29", Parse("2020-02-29").AddMonthsNoOverflow(3).ToDateString())
		s.Equal("2020-10-31", Parse("2020-08-31").AddMonthsNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMonth().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-02-01", NewCarbon().AddMonth().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMonth().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonth().ToDateString())
		s.Equal("2020-03-28", Parse("2020-02-28").AddMonth().ToDateString())
		s.Equal("2020-03-29", Parse("2020-02-29").AddMonth().ToDateString())
		s.Equal("2020-12-30", Parse("2020-11-30").AddMonth().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMonthNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMonthNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-02-01", NewCarbon().AddMonthNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMonthNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMonthNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-02-01", Parse("2020-01-01").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-03-28", Parse("2020-02-28").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-03-29", Parse("2020-02-29").AddMonthNoOverflow().ToDateString())
		s.Equal("2020-12-30", Parse("2020-11-30").AddMonthNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonths() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMonths(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-11-01", NewCarbon().SubMonths(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMonths(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMonths(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubMonths(0).ToDateString())
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonths(1).ToDateString())
		s.Equal("2019-11-01", Parse("2020-01-01").SubMonths(2).ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubMonths(3).ToDateString())
		s.Equal("2020-07-01", Parse("2020-08-31").SubMonths(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonthsNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMonthsNoOverflow(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-11-01", NewCarbon().SubMonthsNoOverflow(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMonthsNoOverflow(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMonthsNoOverflow(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubMonthsNoOverflow(0).ToDateString())
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonthsNoOverflow(1).ToDateString())
		s.Equal("2019-11-01", Parse("2020-01-01").SubMonthsNoOverflow(2).ToDateString())
		s.Equal("2019-11-29", Parse("2020-02-29").SubMonthsNoOverflow(3).ToDateString())
		s.Equal("2020-06-30", Parse("2020-08-31").SubMonthsNoOverflow(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonth() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMonth().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-01", NewCarbon().SubMonth().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMonth().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMonth().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonth().ToDateString())
		s.Equal("2020-01-28", Parse("2020-02-28").SubMonth().ToDateString())
		s.Equal("2020-01-29", Parse("2020-02-29").SubMonth().ToDateString())
		s.Equal("2020-10-30", Parse("2020-11-30").SubMonth().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMonthNoOverflow() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMonthNoOverflow().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-01", NewCarbon().SubMonthNoOverflow().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMonthNoOverflow().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMonthNoOverflow().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-01", Parse("2020-01-01").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-01-28", Parse("2020-02-28").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-01-29", Parse("2020-02-29").SubMonthNoOverflow().ToDateString())
		s.Equal("2020-10-30", Parse("2020-11-30").SubMonthNoOverflow().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddWeeks(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-15", NewCarbon().AddWeeks(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddWeeks(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddWeeks(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddWeeks(0).ToDateString())
		s.Equal("2020-01-08", Parse("2020-01-01").AddWeeks(1).ToDateString())
		s.Equal("2020-01-15", Parse("2020-01-01").AddWeeks(2).ToDateString())
		s.Equal("2020-03-21", Parse("2020-02-29").AddWeeks(3).ToDateString())
		s.Equal("2020-09-14", Parse("2020-08-31").AddWeeks(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddWeek().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-08", NewCarbon().AddWeek().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddWeek().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-08", Parse("2020-01-01").AddWeek().ToDateString())
		s.Equal("2020-03-06", Parse("2020-02-28").AddWeek().ToDateString())
		s.Equal("2020-03-07", Parse("2020-02-29").AddWeek().ToDateString())
		s.Equal("2020-12-07", Parse("2020-11-30").AddWeek().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeeks() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubWeeks(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-18", NewCarbon().SubWeeks(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubWeeks(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubWeeks(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubWeeks(0).ToDateString())
		s.Equal("2019-12-25", Parse("2020-01-01").SubWeeks(1).ToDateString())
		s.Equal("2019-12-18", Parse("2020-01-01").SubWeeks(2).ToDateString())
		s.Equal("2020-02-08", Parse("2020-02-29").SubWeeks(3).ToDateString())
		s.Equal("2020-08-17", Parse("2020-08-31").SubWeeks(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubWeek() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubWeek().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-25", NewCarbon().SubWeek().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubWeek().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubWeek().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-25", Parse("2020-01-01").SubWeek().ToDateString())
		s.Equal("2020-02-21", Parse("2020-02-28").SubWeek().ToDateString())
		s.Equal("2020-02-22", Parse("2020-02-29").SubWeek().ToDateString())
		s.Equal("2020-11-23", Parse("2020-11-30").SubWeek().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDays(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-03", NewCarbon().AddDays(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDays(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDays(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").AddDays(0).ToDateString())
		s.Equal("2020-01-02", Parse("2020-01-01").AddDays(1).ToDateString())
		s.Equal("2020-01-03", Parse("2020-01-01").AddDays(2).ToDateString())
		s.Equal("2020-03-03", Parse("2020-02-29").AddDays(3).ToDateString())
		s.Equal("2020-09-02", Parse("2020-08-31").AddDays(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddDay().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-02", NewCarbon().AddDay().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddDay().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-02", Parse("2020-01-01").AddDay().ToDateString())
		s.Equal("2020-02-29", Parse("2020-02-28").AddDay().ToDateString())
		s.Equal("2020-03-01", Parse("2020-02-29").AddDay().ToDateString())
		s.Equal("2020-12-01", Parse("2020-11-30").AddDay().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDays() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDays(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-30", NewCarbon().SubDays(2).ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDays(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDays(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01", Parse("2020-01-01").SubDays(0).ToDateString())
		s.Equal("2019-12-31", Parse("2020-01-01").SubDays(1).ToDateString())
		s.Equal("2019-12-30", Parse("2020-01-01").SubDays(2).ToDateString())
		s.Equal("2020-02-26", Parse("2020-02-29").SubDays(3).ToDateString())
		s.Equal("2020-08-29", Parse("2020-08-31").SubDays(2).ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_SubDay() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubDay().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31", NewCarbon().SubDay().ToDateString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubDay().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubDay().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2019-12-31", Parse("2020-01-01").SubDay().ToDateString())
		s.Equal("2020-02-27", Parse("2020-02-28").SubDay().ToDateString())
		s.Equal("2020-02-28", Parse("2020-02-29").SubDay().ToDateString())
		s.Equal("2020-11-29", Parse("2020-11-30").SubDay().ToDateString())
	})
}

func (s *TravelerSuite) TestCarbon_AddHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddHours(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 02:00:00 +0000 UTC", NewCarbon().AddHours(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddHours(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddHours(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(0).ToString())
		s.Equal("2020-01-01 14:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(1).ToString())
		s.Equal("2020-01-01 15:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddHours(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddHour().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 01:00:00 +0000 UTC", NewCarbon().AddHour().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddHour().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 14:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").AddHour().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubHours() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubHours(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 22:00:00 +0000 UTC", NewCarbon().SubHours(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubHours(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubHours(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(0).ToString())
		s.Equal("2020-01-01 12:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(1).ToString())
		s.Equal("2020-01-01 11:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubHours(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubHour() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubHour().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:00:00 +0000 UTC", NewCarbon().SubHour().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubHour().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubHour().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 12:14:15 +0000 UTC", Parse("2020-08-05 13:14:15").SubHour().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMinutes(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:02:00 +0000 UTC", NewCarbon().AddMinutes(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMinutes(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMinutes(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(0).ToString())
		s.Equal("2020-01-01 13:15:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(1).ToString())
		s.Equal("2020-01-01 13:16:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMinutes(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMinute().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:01:00 +0000 UTC", NewCarbon().AddMinute().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMinute().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:15:15 +0000 UTC", Parse("2020-08-05 13:14:15").AddMinute().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMinutes() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMinutes(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:58:00 +0000 UTC", NewCarbon().SubMinutes(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMinutes(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMinutes(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(0).ToString())
		s.Equal("2020-01-01 13:13:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(1).ToString())
		s.Equal("2020-01-01 13:12:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMinutes(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMinute() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMinute().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:00 +0000 UTC", NewCarbon().SubMinute().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMinute().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMinute().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:13:15 +0000 UTC", Parse("2020-08-05 13:14:15").SubMinute().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddSeconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:02 +0000 UTC", NewCarbon().AddSeconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddSeconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddSeconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(0).ToString())
		s.Equal("2020-01-01 13:14:16 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(1).ToString())
		s.Equal("2020-01-01 13:14:17 +0000 UTC", Parse("2020-01-01 13:14:15").AddSeconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddSecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:01 +0000 UTC", NewCarbon().AddSecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddSecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:16 +0000 UTC", Parse("2020-08-05 13:14:15").AddSecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubSeconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubSeconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:58 +0000 UTC", NewCarbon().SubSeconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubSeconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubSeconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(0).ToString())
		s.Equal("2020-01-01 13:14:14 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(1).ToString())
		s.Equal("2020-01-01 13:14:13 +0000 UTC", Parse("2020-01-01 13:14:15").SubSeconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubSecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubSecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59 +0000 UTC", NewCarbon().SubSecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubSecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubSecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:14 +0000 UTC", Parse("2020-08-05 13:14:15").SubSecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMilliseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMilliseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.002 +0000 UTC", NewCarbon().AddMilliseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMilliseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMilliseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMilliseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMillisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMillisecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.001 +0000 UTC", NewCarbon().AddMillisecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMillisecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMillisecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15.001 +0000 UTC", Parse("2020-08-05 13:14:15").AddMillisecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMilliseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMilliseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.998 +0000 UTC", NewCarbon().SubMilliseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMilliseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMilliseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMilliseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMillisecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMillisecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999 +0000 UTC", NewCarbon().SubMillisecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMillisecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMillisecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:14.999 +0000 UTC", Parse("2020-08-05 13:14:15").SubMillisecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMicroseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMicroseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.000002 +0000 UTC", NewCarbon().AddMicroseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMicroseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMicroseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddMicroseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddMicrosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddMicrosecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.000001 +0000 UTC", NewCarbon().AddMicrosecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddMicrosecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddMicrosecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15.000001 +0000 UTC", Parse("2020-08-05 13:14:15").AddMicrosecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMicroseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMicroseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999998 +0000 UTC", NewCarbon().SubMicroseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMicroseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMicroseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubMicroseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubMicrosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubMicrosecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999999 +0000 UTC", NewCarbon().SubMicrosecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubMicrosecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubMicrosecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:14.999999 +0000 UTC", Parse("2020-08-05 13:14:15").SubMicrosecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddNanoseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddNanoseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.000000002 +0000 UTC", NewCarbon().AddNanoseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddNanoseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddNanoseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(0).ToString())
		s.Equal("2020-01-01 13:14:15.000000001 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(1).ToString())
		s.Equal("2020-01-01 13:14:15.000000002 +0000 UTC", Parse("2020-01-01 13:14:15").AddNanoseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_AddNanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.AddNanosecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0001-01-01 00:00:00.000000001 +0000 UTC", NewCarbon().AddNanosecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").AddNanosecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").AddNanosecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:15.000000001 +0000 UTC", Parse("2020-08-05 13:14:15").AddNanosecond().ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubNanoseconds() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubNanoseconds(2).ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999999998 +0000 UTC", NewCarbon().SubNanoseconds(2).ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubNanoseconds(2).ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubNanoseconds(2).ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-01-01 13:14:15 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(0).ToString())
		s.Equal("2020-01-01 13:14:14.999999999 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(1).ToString())
		s.Equal("2020-01-01 13:14:14.999999998 +0000 UTC", Parse("2020-01-01 13:14:15").SubNanoseconds(2).ToString())
	})
}

func (s *TravelerSuite) TestCarbon_SubNanosecond() {
	s.Run("nil carbon", func() {
		var c *Carbon
		c = nil
		s.Empty(c.SubNanosecond().ToString())
	})

	s.Run("zero carbon", func() {
		s.Equal("0000-12-31 23:59:59.999999999 +0000 UTC", NewCarbon().SubNanosecond().ToString())
	})

	s.Run("empty carbon", func() {
		s.Empty(Parse("").SubNanosecond().ToString())
	})

	s.Run("error carbon", func() {
		s.Empty(Parse("xxx").SubNanosecond().ToString())
	})

	s.Run("valid carbon", func() {
		s.Equal("2020-08-05 13:14:14.999999999 +0000 UTC", Parse("2020-08-05 13:14:15").SubNanosecond().ToString())
	})
}
