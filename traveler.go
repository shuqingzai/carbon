package carbon

import (
	"time"
)

// Now returns a Carbon instance for now.
// 当前
func Now(timezone ...string) *Carbon {
	if IsTestNow() {
		return frozenNow.testNow
	}
	var (
		loc *Location
		err error
	)
	tz := DefaultTimezone
	if len(timezone) > 0 {
		tz = timezone[0]
	}
	if loc, err = parseTimezone(tz); err != nil {
		return &Carbon{Error: err}
	}
	return CreateFromStdTime(time.Now().In(loc))
}

// Tomorrow returns a Carbon instance for tomorrow.
// 明天
func Tomorrow(timezone ...string) *Carbon {
	now := Now(timezone...)
	if now.IsInvalid() {
		return now
	}
	return now.Copy().AddDay()
}

// Yesterday returns a Carbon instance for yesterday.
// 昨天
func Yesterday(timezone ...string) *Carbon {
	now := Now(timezone...)
	if now.IsInvalid() {
		return now
	}
	return now.Copy().SubDay()
}

// AddDuration adds duration.
// 按照时长增加时间,支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c *Carbon) AddDuration(duration string) *Carbon {
	if c.IsInvalid() {
		return c
	}
	var (
		td  Duration
		err error
	)
	if td, err = parseDuration(duration); err != nil {
		c.Error = err
		return c
	}
	c.time = c.StdTime().Add(td)
	return c
}

// SubDuration subtracts duration.
// 按照时长减少时间,支持整数/浮点数和符号ns(纳秒)、us(微妙)、ms(毫秒)、s(秒)、m(分钟)、h(小时)的组合
func (c *Carbon) SubDuration(duration string) *Carbon {
	return c.AddDuration("-" + duration)
}

// AddCenturies adds some centuries.
// N 个世纪后
func (c *Carbon) AddCenturies(centuries int) *Carbon {
	return c.AddYears(centuries * YearsPerCentury)
}

// AddCenturiesNoOverflow adds some centuries without overflowing month.
// N 个世纪后(月份不溢出)
func (c *Carbon) AddCenturiesNoOverflow(centuries int) *Carbon {
	return c.AddYearsNoOverflow(centuries * YearsPerCentury)
}

// AddCentury adds one century.
// 1 个世纪后
func (c *Carbon) AddCentury() *Carbon {
	return c.AddCenturies(1)
}

// AddCenturyNoOverflow adds one century without overflowing month.
// 1 个世纪后(月份不溢出)
func (c *Carbon) AddCenturyNoOverflow() *Carbon {
	return c.AddCenturiesNoOverflow(1)
}

// SubCenturies subtracts some centuries.
// N 个世纪前
func (c *Carbon) SubCenturies(centuries int) *Carbon {
	return c.SubYears(centuries * YearsPerCentury)
}

// SubCenturiesNoOverflow subtracts some centuries without overflowing month.
// N 个世纪前(月份不溢出)
func (c *Carbon) SubCenturiesNoOverflow(centuries int) *Carbon {
	return c.SubYearsNoOverflow(centuries * YearsPerCentury)
}

// SubCentury subtracts one century.
// 1 个世纪前
func (c *Carbon) SubCentury() *Carbon {
	return c.SubCenturies(1)
}

// SubCenturyNoOverflow subtracts one century without overflowing month.
// 1 个世纪前(月份不溢出)
func (c *Carbon) SubCenturyNoOverflow() *Carbon {
	return c.SubCenturiesNoOverflow(1)
}

// AddDecades adds some decades.
// N 个年代后
func (c *Carbon) AddDecades(decades int) *Carbon {
	return c.AddYears(decades * YearsPerDecade)
}

// AddDecadesNoOverflow adds some decades without overflowing month.
// N 个年代后(月份不溢出)
func (c *Carbon) AddDecadesNoOverflow(decades int) *Carbon {
	return c.AddYearsNoOverflow(decades * YearsPerDecade)
}

// AddDecade adds one decade.
// 1 个年代后
func (c *Carbon) AddDecade() *Carbon {
	return c.AddDecades(1)
}

// AddDecadeNoOverflow adds one decade without overflowing month.
// 1 个年代后(月份不溢出)
func (c *Carbon) AddDecadeNoOverflow() *Carbon {
	return c.AddDecadesNoOverflow(1)
}

// SubDecades subtracts some decades.
// N 个年代后
func (c *Carbon) SubDecades(decades int) *Carbon {
	return c.SubYears(decades * YearsPerDecade)
}

// SubDecadesNoOverflow subtracts some decades without overflowing month.
// N 个年代后(月份不溢出)
func (c *Carbon) SubDecadesNoOverflow(decades int) *Carbon {
	return c.SubYearsNoOverflow(decades * YearsPerDecade)
}

// SubDecade subtracts one decade.
// 1 个年代后
func (c *Carbon) SubDecade() *Carbon {
	return c.SubDecades(1)
}

// SubDecadeNoOverflow subtracts one decade without overflowing month.
// 1 个年代后(月份不溢出)
func (c *Carbon) SubDecadeNoOverflow() *Carbon {
	return c.SubDecadesNoOverflow(1)
}

// AddYears adds some years.
// N 年后
func (c *Carbon) AddYears(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(years, 0, 0)
	return c
}

// AddYearsNoOverflow adds some years without overflowing month.
// N 年后(月份不溢出)
func (c *Carbon) AddYearsNoOverflow(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	nanosecond := c.Nanosecond()
	year, month, day, hour, minute, second := c.DateTime()
	// 获取N年后本月的最后一天
	lastYear, lastMonth, lastDay := create(year+years, month+1, 0, hour, minute, second, nanosecond).SetLocation(c.loc).Date()
	if day > lastDay {
		day = lastDay
	}
	return create(lastYear, lastMonth, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// AddYear adds one year.
// 1 年后
func (c *Carbon) AddYear() *Carbon {
	return c.AddYears(1)
}

// AddYearNoOverflow adds one year without overflowing month.
// 1 年后(月份不溢出)
func (c *Carbon) AddYearNoOverflow() *Carbon {
	return c.AddYearsNoOverflow(1)
}

// SubYears subtracts some years.
// N 年前
func (c *Carbon) SubYears(years int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	return c.AddYears(-years)
}

// SubYearsNoOverflow subtracts some years without overflowing month.
// N 年前(月份不溢出)
func (c *Carbon) SubYearsNoOverflow(years int) *Carbon {
	return c.AddYearsNoOverflow(-years)
}

// SubYear subtracts one year.
// 1 年前
func (c *Carbon) SubYear() *Carbon {
	return c.SubYears(1)
}

// SubYearNoOverflow subtracts one year without overflowing month.
// 1 年前(月份不溢出)
func (c *Carbon) SubYearNoOverflow() *Carbon {
	return c.SubYearsNoOverflow(1)
}

// AddQuarters adds some quarters
// N 个季度后
func (c *Carbon) AddQuarters(quarters int) *Carbon {
	return c.AddMonths(quarters * MonthsPerQuarter)
}

// AddQuartersNoOverflow adds quarters without overflowing month.
// N 个季度后(月份不溢出)
func (c *Carbon) AddQuartersNoOverflow(quarters int) *Carbon {
	return c.AddMonthsNoOverflow(quarters * MonthsPerQuarter)
}

// AddQuarter adds one quarter
// 1 个季度后
func (c *Carbon) AddQuarter() *Carbon {
	return c.AddQuarters(1)
}

// AddQuarterNoOverflow adds one quarter without overflowing month.
// 1 个季度后(月份不溢出)
func (c *Carbon) AddQuarterNoOverflow() *Carbon {
	return c.AddQuartersNoOverflow(1)
}

// SubQuarters subtracts some quarters.
// N 个季度前
func (c *Carbon) SubQuarters(quarters int) *Carbon {
	return c.AddQuarters(-quarters)
}

// SubQuartersNoOverflow subtracts some quarters without overflowing month.
// N 个季度前(月份不溢出)
func (c *Carbon) SubQuartersNoOverflow(quarters int) *Carbon {
	return c.AddMonthsNoOverflow(-quarters * MonthsPerQuarter)
}

// SubQuarter subtracts one quarter.
// 1 个季度前
func (c *Carbon) SubQuarter() *Carbon {
	return c.SubQuarters(1)
}

// SubQuarterNoOverflow subtracts one quarter without overflowing month.
// 1 个季度前(月份不溢出)
func (c *Carbon) SubQuarterNoOverflow() *Carbon {
	return c.SubQuartersNoOverflow(1)
}

// AddMonths adds some months.
// N 个月后
func (c *Carbon) AddMonths(months int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(0, months, 0)
	return c
}

// AddMonthsNoOverflow adds some months without overflowing month.
// N 个月后(月份不溢出)
func (c *Carbon) AddMonthsNoOverflow(months int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	nanosecond := c.Nanosecond()
	year, month, day, hour, minute, second := c.DateTime()
	// 获取N月后的最后一天
	lastYear, lastMonth, lastDay := create(year, month+months+1, 0, hour, minute, second, nanosecond).SetLocation(c.loc).Date()
	if day > lastDay {
		day = lastDay
	}
	return create(lastYear, lastMonth, day, hour, minute, second, nanosecond).SetLocation(c.loc)
}

// AddMonth adds one month.
// 1 个月后
func (c *Carbon) AddMonth() *Carbon {
	return c.AddMonths(1)
}

// AddMonthNoOverflow adds one month without overflowing month.
// 1 个月后(月份不溢出)
func (c *Carbon) AddMonthNoOverflow() *Carbon {
	return c.AddMonthsNoOverflow(1)
}

// SubMonths subtracts some months.
// N 个月前
func (c *Carbon) SubMonths(months int) *Carbon {
	return c.AddMonths(-months)
}

// SubMonthsNoOverflow subtracts some months without overflowing month.
// N 个月前(月份不溢出)
func (c *Carbon) SubMonthsNoOverflow(months int) *Carbon {
	return c.AddMonthsNoOverflow(-months)
}

// SubMonth subtracts one month.
// 1 个月前
func (c *Carbon) SubMonth() *Carbon {
	return c.SubMonths(1)
}

// SubMonthNoOverflow subtracts one month without overflowing month.
// 1 个月前(月份不溢出)
func (c *Carbon) SubMonthNoOverflow() *Carbon {
	return c.SubMonthsNoOverflow(1)
}

// AddWeeks adds some weeks.
// N 周后
func (c *Carbon) AddWeeks(weeks int) *Carbon {
	return c.AddDays(weeks * DaysPerWeek)
}

// AddWeek adds one week.
// 1 周后
func (c *Carbon) AddWeek() *Carbon {
	return c.AddWeeks(1)
}

// SubWeeks subtracts some weeks.
// N 周前
func (c *Carbon) SubWeeks(weeks int) *Carbon {
	return c.SubDays(weeks * DaysPerWeek)
}

// SubWeek subtracts one week.
// 1 周前
func (c *Carbon) SubWeek() *Carbon {
	return c.SubWeeks(1)
}

// AddDays adds some days.
// N 天后
func (c *Carbon) AddDays(days int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	c.time = c.StdTime().AddDate(0, 0, days)
	return c
}

// AddDay adds one day.
// 1 天后
func (c *Carbon) AddDay() *Carbon {
	return c.AddDays(1)
}

// SubDays subtracts some days.
// N 天前
func (c *Carbon) SubDays(days int) *Carbon {
	return c.AddDays(-days)
}

// SubDay subtracts one day.
// 1 天前
func (c *Carbon) SubDay() *Carbon {
	return c.SubDays(1)
}

// AddHours adds some hours.
// N 小时后
func (c *Carbon) AddHours(hours int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(hours) * time.Hour
	c.time = c.StdTime().Add(td)
	return c
}

// AddHour adds one hour.
// 1 小时后
func (c *Carbon) AddHour() *Carbon {
	return c.AddHours(1)
}

// SubHours subtracts some hours.
// N 小时前
func (c *Carbon) SubHours(hours int) *Carbon {
	return c.AddHours(-hours)
}

// SubHour subtracts one hour.
// 1 小时前
func (c *Carbon) SubHour() *Carbon {
	return c.SubHours(1)
}

// AddMinutes adds some minutes.
// N 分钟后
func (c *Carbon) AddMinutes(minutes int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(minutes) * time.Minute
	c.time = c.StdTime().Add(td)
	return c
}

// AddMinute adds one minute.
// 1 分钟后
func (c *Carbon) AddMinute() *Carbon {
	return c.AddMinutes(1)
}

// SubMinutes subtracts some minutes.
// N 分钟前
func (c *Carbon) SubMinutes(minutes int) *Carbon {
	return c.AddMinutes(-minutes)
}

// SubMinute subtracts one minute.
// 1 分钟前
func (c *Carbon) SubMinute() *Carbon {
	return c.SubMinutes(1)
}

// AddSeconds adds some seconds.
// N 秒钟后
func (c *Carbon) AddSeconds(seconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(seconds) * time.Second
	c.time = c.StdTime().Add(td)
	return c
}

// AddSecond adds one second.
// 1 秒钟后
func (c *Carbon) AddSecond() *Carbon {
	return c.AddSeconds(1)
}

// SubSeconds subtracts some seconds.
// N 秒钟前
func (c *Carbon) SubSeconds(seconds int) *Carbon {
	return c.AddSeconds(-seconds)
}

// SubSecond subtracts one second.
// 1 秒钟前
func (c *Carbon) SubSecond() *Carbon {
	return c.SubSeconds(1)
}

// AddMilliseconds adds some milliseconds.
// N 毫秒后
func (c *Carbon) AddMilliseconds(milliseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(milliseconds) * time.Millisecond
	c.time = c.StdTime().Add(td)
	return c
}

// AddMillisecond adds one millisecond.
// 1 毫秒后
func (c *Carbon) AddMillisecond() *Carbon {
	return c.AddMilliseconds(1)
}

// SubMilliseconds subtracts some milliseconds.
// N 毫秒前
func (c *Carbon) SubMilliseconds(milliseconds int) *Carbon {
	return c.AddMilliseconds(-milliseconds)
}

// SubMillisecond subtracts one millisecond.
// 1 毫秒前
func (c *Carbon) SubMillisecond() *Carbon {
	return c.SubMilliseconds(1)
}

// AddMicroseconds adds some microseconds.
// N 微秒后
func (c *Carbon) AddMicroseconds(microseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(microseconds) * time.Microsecond
	c.time = c.StdTime().Add(td)
	return c
}

// AddMicrosecond adds one microsecond.
// 1 微秒后
func (c *Carbon) AddMicrosecond() *Carbon {
	return c.AddMicroseconds(1)
}

// SubMicroseconds subtracts some microseconds.
// N 微秒前
func (c *Carbon) SubMicroseconds(microseconds int) *Carbon {
	return c.AddMicroseconds(-microseconds)
}

// SubMicrosecond subtracts one microsecond.
// 1 微秒前
func (c *Carbon) SubMicrosecond() *Carbon {
	return c.SubMicroseconds(1)
}

// AddNanoseconds adds some nanoseconds.
// N 纳秒后
func (c *Carbon) AddNanoseconds(nanoseconds int) *Carbon {
	if c.IsInvalid() {
		return c
	}
	td := Duration(nanoseconds) * time.Nanosecond
	c.time = c.StdTime().Add(td)
	return c
}

// AddNanosecond adds one nanosecond.
// 1 纳秒后
func (c *Carbon) AddNanosecond() *Carbon {
	return c.AddNanoseconds(1)
}

// SubNanoseconds subtracts some nanoseconds.
// N 纳秒前
func (c *Carbon) SubNanoseconds(nanoseconds int) *Carbon {
	return c.AddNanoseconds(-nanoseconds)
}

// SubNanosecond subtracts one nanosecond.
// 1 纳秒前
func (c *Carbon) SubNanosecond() *Carbon {
	return c.SubNanoseconds(1)
}
