package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type FrozenSuite struct {
	suite.Suite
}

func TestFrozenSuite(t *testing.T) {
	suite.Run(t, new(FrozenSuite))
}

func (s *FrozenSuite) TestSetTestNow() {
	now := Parse("2020-08-05")

	SetTestNow(nil)
	s.NotEqual(now, Now())

	SetTestNow(now)
	s.Equal("2020-08-05", Now().ToDateString())
	s.Equal("2020-08-04", Yesterday().ToDateString())
	s.Equal("2020-08-06", Tomorrow().ToDateString())
	s.Equal("just now", Now().DiffForHumans())
	s.Equal("1 day ago", Yesterday().DiffForHumans())
	s.Equal("1 day from now", Tomorrow().DiffForHumans())
	s.Equal("2 months from now", Parse("2020-10-05").DiffForHumans())
	s.Equal("2 months before", now.DiffForHumans(Parse("2020-10-05")))
	s.True(IsTestNow())

	ClearTestNow()
	s.Equal(time.Now().In(time.UTC).Format(DateLayout), Now().ToDateString())
	s.Equal(time.Now().In(time.UTC).Add(time.Hour*-24).Format(DateLayout), Yesterday().ToDateString())
	s.Equal(time.Now().In(time.UTC).Add(time.Hour*24).Format(DateLayout), Tomorrow().ToDateString())
	s.False(IsTestNow())
}
