package carbon

import (
	"gitee.com/golang-package/carbon/v2/calendar/julian"
	"gitee.com/golang-package/carbon/v2/calendar/lunar"
	"gitee.com/golang-package/carbon/v2/calendar/persian"
)

// Lunar converts Carbon instance to Lunar instance.
// 将 Carbon 实例转化为 Lunar 实例
func (c *Carbon) Lunar() *lunar.Lunar {
	if c.IsNil() {
		return nil
	}
	if c.IsZero() || c.IsEmpty() {
		return &lunar.Lunar{}
	}
	if c.HasError() {
		return &lunar.Lunar{Error: c.Error}
	}
	return lunar.FromStdTime(c.StdTime())
}

// CreateFromLunar creates a Carbon instance from Lunar date.
// 从 农历日期 创建 Carbon 实例
func CreateFromLunar(year, month, day int, isLeapMonth bool) *Carbon {
	l := lunar.NewLunar(year, month, day, isLeapMonth)
	if l.Error != nil {
		return &Carbon{Error: l.Error}
	}
	return NewCarbon(l.ToGregorian(DefaultTimezone).Time)
}

// Julian converts Carbon instance to Julian instance.
// 将 Carbon 实例转化为 Julian 实例
func (c *Carbon) Julian() *julian.Julian {
	if c.IsNil() {
		return nil
	}
	if c.IsEmpty() {
		return &julian.Julian{}
	}
	if c.HasError() {
		return &julian.Julian{}
	}
	return julian.FromStdTime(c.StdTime())
}

// CreateFromJulian creates a Carbon instance from Julian Day or Modified Julian Day.
// 从 儒略日/简化儒略日 创建 Carbon 实例
func CreateFromJulian(f float64) *Carbon {
	return NewCarbon(julian.NewJulian(f).ToGregorian(DefaultTimezone).Time)
}

// Persian converts Carbon instance to Persian instance.
// 将 Carbon 实例转化为 Persian 实例
func (c *Carbon) Persian() *persian.Persian {
	if c.IsNil() {
		return nil
	}
	if c.IsZero() || c.IsEmpty() {
		return &persian.Persian{}
	}
	if c.HasError() {
		return &persian.Persian{Error: c.Error}
	}
	return persian.FromStdTime(c.StdTime())
}

// CreateFromPersian creates a Carbon instance from Persian date.
// 从 波斯日期 创建 Carbon 实例
func CreateFromPersian(year, month, day int) *Carbon {
	p := persian.NewPersian(year, month, day)
	if p.Error != nil {
		return &Carbon{Error: p.Error}
	}
	return NewCarbon(p.ToGregorian(DefaultTimezone).Time)
}
