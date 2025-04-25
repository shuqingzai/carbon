// @Package carbon
// @Description a simple, semantic and developer-friendly time package for golang
// @Page github.com/dromara/carbon
// @Developer gouguoyin
// @Blog www.gouguoyin.com
// @Email 245629560@qq.com

// Package carbon is a simple, semantic and developer-friendly time package for golang.
package carbon

import (
	"time"
)

type StdTime = time.Time
type Weekday = time.Weekday
type Location = time.Location
type Duration = time.Duration

// Carbon defines a Carbon struct.
// 定义 Carbon 结构体
type Carbon struct {
	layout       string
	time         StdTime
	weekStartsAt Weekday
	weekendDays  []Weekday
	loc          *Location
	lang         *Language
	Error        error
}

// NewCarbon returns a new Carbon instance.
// 返回 Carbon 实例
func NewCarbon(time ...StdTime) *Carbon {
	c := new(Carbon)
	c.lang = NewLanguage().SetLocale(DefaultLocale)
	c.layout = DefaultLayout
	c.weekStartsAt = DefaultWeekStartsAt
	c.weekendDays = DefaultWeekendDays
	if len(time) > 0 {
		c.time = time[0]
		c.loc = c.time.Location()
		return c
	}
	c.loc, c.Error = parseTimezone(DefaultTimezone)
	return c
}

// Copy returns a new copy of the current Carbon instance
// 复制 Carbon 实例
func (c *Carbon) Copy() *Carbon {
	if c == nil {
		return nil
	}

	return &Carbon{
		layout:       c.layout,
		time:         time.Date(c.Year(), time.Month(c.Month()), c.Day(), c.Hour(), c.Minute(), c.Second(), c.Nanosecond(), c.loc),
		weekStartsAt: c.weekStartsAt,
		weekendDays:  c.weekendDays,
		loc:          c.loc,
		lang:         c.lang.Copy(),
		Error:        c.Error,
	}
}
