package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type CreatorSuite struct {
	suite.Suite
}

func TestCreatorSuite(t *testing.T) {
	suite.Run(t, new(CreatorSuite))
}

func (s *CreatorSuite) TestCreateFromStdTime() {
	s.Run("error timezone", func() {
		s.Error(CreateFromStdTime(time.Now(), "xxx").Error)
	})

	s.Run("without timezone", func() {
		now := time.Now()
		s.Equal(now.Unix(), CreateFromStdTime(time.Now()).Timestamp())
	})

	s.Run("with timezone", func() {
		now := time.Now().In(time.Local)
		s.Equal(now.Unix(), CreateFromStdTime(now).Timestamp())
		s.Equal(now.Unix(), CreateFromStdTime(now, UTC).Timestamp())
	})
}

func (s *CreatorSuite) TestCreateFromTimestamp() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimestamp(0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("1969-12-31 23:59:59 +0000 UTC", CreateFromTimestamp(-1).ToString())
		s.Equal("1970-01-01 00:00:00 +0000 UTC", CreateFromTimestamp(0).ToString())
		s.Equal("1970-01-01 00:00:01 +0000 UTC", CreateFromTimestamp(1).ToString())
		s.Equal("2022-04-12 03:55:55 +0000 UTC", CreateFromTimestamp(1649735755).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("1970-01-01 07:59:59 +0800 CST", CreateFromTimestamp(-1, PRC).ToString())
		s.Equal("1970-01-01 08:00:00 +0800 CST", CreateFromTimestamp(0, PRC).ToString())
		s.Equal("1970-01-01 08:00:01 +0800 CST", CreateFromTimestamp(1, PRC).ToString())
		s.Equal("2022-04-12 11:55:55 +0800 CST", CreateFromTimestamp(1649735755, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromTimestampMilli() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimestampMilli(0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("1969-12-31 23:59:59.999 +0000 UTC", CreateFromTimestampMilli(-1).ToString())
		s.Equal("1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMilli(0).ToString())
		s.Equal("1970-01-01 00:00:00.001 +0000 UTC", CreateFromTimestampMilli(1).ToString())
		s.Equal("2022-04-12 03:55:55.981 +0000 UTC", CreateFromTimestampMilli(1649735755981).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("1970-01-01 07:59:59.999 +0800 CST", CreateFromTimestampMilli(-1, PRC).ToString())
		s.Equal("1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMilli(0, PRC).ToString())
		s.Equal("1970-01-01 08:00:00.001 +0800 CST", CreateFromTimestampMilli(1, PRC).ToString())
		s.Equal("2022-04-12 11:55:55.981 +0800 CST", CreateFromTimestampMilli(1649735755981, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromTimestampMicro() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimestampMicro(0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("1969-12-31 23:59:59.999999 +0000 UTC", CreateFromTimestampMicro(-1).ToString())
		s.Equal("1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampMicro(0).ToString())
		s.Equal("1970-01-01 00:00:00.000001 +0000 UTC", CreateFromTimestampMicro(1).ToString())
		s.Equal("2022-04-12 03:55:55.981566 +0000 UTC", CreateFromTimestampMicro(1649735755981566).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("1970-01-01 07:59:59.999999 +0800 CST", CreateFromTimestampMicro(-1, PRC).ToString())
		s.Equal("1970-01-01 08:00:00 +0800 CST", CreateFromTimestampMicro(0, PRC).ToString())
		s.Equal("1970-01-01 08:00:00.000001 +0800 CST", CreateFromTimestampMicro(1, PRC).ToString())
		s.Equal("2022-04-12 11:55:55.981566 +0800 CST", CreateFromTimestampMicro(1649735755981566, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromTimestampNano() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimestampNano(0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("1969-12-31 23:59:59.999999999 +0000 UTC", CreateFromTimestampNano(-1).ToString())
		s.Equal("1970-01-01 00:00:00 +0000 UTC", CreateFromTimestampNano(0).ToString())
		s.Equal("1970-01-01 00:00:00.000000001 +0000 UTC", CreateFromTimestampNano(1).ToString())
		s.Equal("2022-04-12 03:55:55.981566888 +0000 UTC", CreateFromTimestampNano(1649735755981566888).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("1970-01-01 07:59:59.999999999 +0800 CST", CreateFromTimestampNano(-1, PRC).ToString())
		s.Equal("1970-01-01 08:00:00 +0800 CST", CreateFromTimestampNano(0, PRC).ToString())
		s.Equal("1970-01-01 08:00:00.000000001 +0800 CST", CreateFromTimestampNano(1, PRC).ToString())
		s.Equal("2022-04-12 11:55:55.981566888 +0800 CST", CreateFromTimestampNano(1649735755981566888, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateTime() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateTime(0, 0, 0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTime(0, 0, 0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 13:14:15 +0000 UTC", CreateFromDateTime(2020, 8, 5, 13, 14, 15).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTime(0, 0, 0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 13:14:15 +0800 CST", CreateFromDateTime(2020, 8, 5, 13, 14, 15, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateTimeMilli() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 13:14:15.999 +0000 UTC", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMilli(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 13:14:15.999 +0800 CST", CreateFromDateTimeMilli(2020, 8, 5, 13, 14, 15, 999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateTimeMicro() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 13:14:15.999999 +0000 UTC", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeMicro(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 13:14:15.999999 +0800 CST", CreateFromDateTimeMicro(2020, 8, 5, 13, 14, 15, 999999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateTimeNano() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 13:14:15.999999999 +0000 UTC", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateTimeNano(0, 0, 0, 0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 13:14:15.999999999 +0800 CST", CreateFromDateTimeNano(2020, 8, 5, 13, 14, 15, 999999999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDate() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDate(0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDate(0, 0, 0).ToString())
		s.Equal("2020-08-05 00:00:00 +0000 UTC", CreateFromDate(2020, 8, 5).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDate(0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 00:00:00 +0800 CST", CreateFromDate(2020, 8, 5, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateMilli() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateMilli(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMilli(0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 00:00:00.999 +0000 UTC", CreateFromDateMilli(2020, 8, 5, 999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMilli(0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 00:00:00.999 +0800 CST", CreateFromDateMilli(2020, 8, 5, 999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateMicro() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateMicro(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateMicro(0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 00:00:00.999999 +0000 UTC", CreateFromDateMicro(2020, 8, 5, 999999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateMicro(0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 00:00:00.999999 +0800 CST", CreateFromDateMicro(2020, 8, 5, 999999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromDateNano() {
	s.Run("error timezone", func() {
		s.Error(CreateFromDateNano(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0000 UTC", CreateFromDateNano(0, 0, 0, 0).ToString())
		s.Equal("2020-08-05 00:00:00.999999999 +0000 UTC", CreateFromDateNano(2020, 8, 5, 999999999).ToString())
	})

	s.Run("with timezone", func() {
		s.Equal("-0001-11-30 00:00:00 +0805 LMT", CreateFromDateNano(0, 0, 0, 0, PRC).ToString())
		s.Equal("2020-08-05 00:00:00.999999999 +0800 CST", CreateFromDateNano(2020, 8, 5, 999999999, PRC).ToString())
	})
}

func (s *CreatorSuite) TestCreateFromTime() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTime(0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("00:00:00", CreateFromTime(0, 0, 0).ToTimeString())
		s.Equal("13:14:15", CreateFromTime(13, 14, 15).ToTimeString())
	})

	s.Run("with timezone", func() {
		s.Equal("00:00:00", CreateFromTime(0, 0, 0, PRC).ToTimeString())
		s.Equal("13:14:15", CreateFromTime(13, 14, 15, PRC).ToTimeString())
	})
}

func (s *CreatorSuite) TestCreateFromTimeMilli() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimeMilli(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("00:00:00", CreateFromTimeMilli(0, 0, 0, 0).ToTimeMilliString())
		s.Equal("13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999).ToTimeMilliString())
	})

	s.Run("with timezone", func() {
		s.Equal("00:00:00", CreateFromTimeMilli(0, 0, 0, 0, PRC).ToTimeMilliString())
		s.Equal("13:14:15.999", CreateFromTimeMilli(13, 14, 15, 999, PRC).ToTimeMilliString())
	})
}

func (s *CreatorSuite) TestCreateFromTimeMicro() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimeMicro(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("00:00:00", CreateFromTimeMicro(0, 0, 0, 0).ToTimeMicroString())
		s.Equal("13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999).ToTimeMicroString())
	})

	s.Run("with timezone", func() {
		s.Equal("00:00:00", CreateFromTimeMicro(0, 0, 0, 0, PRC).ToTimeMicroString())
		s.Equal("13:14:15.999999", CreateFromTimeMicro(13, 14, 15, 999999, PRC).ToTimeMicroString())
	})
}

func (s *CreatorSuite) TestCreateFromTimeNano() {
	s.Run("error timezone", func() {
		s.Error(CreateFromTimeNano(0, 0, 0, 0, "xxx").Error)
	})

	s.Run("without timezone", func() {
		s.Equal("00:00:00", CreateFromTimeNano(0, 0, 0, 0).ToTimeNanoString())
		s.Equal("13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999).ToTimeNanoString())
	})

	s.Run("with timezone", func() {
		s.Equal("00:00:00", CreateFromTimeNano(0, 0, 0, 0, PRC).ToTimeNanoString())
		s.Equal("13:14:15.999999999", CreateFromTimeNano(13, 14, 15, 999999999, PRC).ToTimeNanoString())
	})
}
