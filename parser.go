package carbon

import (
	"fmt"
	"time"
)

// Parse parses a standard time string as a Carbon instance.
// 将标准格式时间字符串解析成 Carbon 实例
func Parse(value string, timezone ...string) *Carbon {
	if value == "" {
		return nil
	}
	var (
		loc *Location
		err error
	)
	tz := DefaultTimezone
	if len(timezone) > 0 {
		tz = timezone[0]
	}
	loc, err = parseTimezone(tz)
	if err != nil {
		return &Carbon{Error: err}
	}
	switch value {
	case "now":
		return Now(tz)
	case "yesterday":
		return Yesterday(tz)
	case "tomorrow":
		return Tomorrow(tz)
	}
	c := NewCarbon().SetLocation(loc)
	for i := range defaultLayouts {
		if tt, err := time.ParseInLocation(defaultLayouts[i], value, loc); err == nil {
			c.time = tt
			c.layout = defaultLayouts[i]
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseByLayout parses a time string as a Carbon instance by a confirmed layout
// 通过一个确认的 布局模板 将时间字符串解析成 Carbon 实例
func ParseByLayout(value, layout string, timezone ...string) *Carbon {
	if value == "" {
		return nil
	}
	if layout == "" {
		return &Carbon{Error: ErrEmptyLayout()}
	}
	var (
		loc *Location
		err error
	)
	tz := DefaultTimezone
	if len(timezone) > 0 {
		tz = timezone[0]
	}
	loc, err = parseTimezone(tz)
	if err != nil {
		return &Carbon{Error: err}
	}
	c := NewCarbon().SetLocation(loc)

	// timestamp layouts
	switch layout {
	case TimestampLayout:
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestamp(ts, c.Timezone())
	case TimestampMilliLayout:
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampMilli(ts, c.Timezone())
	case TimestampMicroLayout:
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampMicro(ts, c.Timezone())
	case TimestampNanoLayout:
		ts, err := parseTimestamp(value)
		if err != nil {
			c.Error = err
			return c
		}
		return CreateFromTimestampNano(ts, c.Timezone())
	}

	// other layouts
	tt, err := time.ParseInLocation(layout, value, c.loc)
	if err != nil {
		c.Error = fmt.Errorf("%w: %w", ErrMismatchedLayout(value, layout), c.Error)
		return c
	}
	c.time = tt
	c.layout = layout
	return c
}

// ParseByFormat parses a time string as a Carbon instance by a confirmed format.
// 通过一个确认的 格式模板 将时间字符串解析成 Carbon 实例
func ParseByFormat(value, format string, timezone ...string) *Carbon {
	if value == "" {
		return nil
	}
	if format == "" {
		return &Carbon{Error: ErrEmptyFormat()}
	}
	c := ParseByLayout(value, format2layout(format), timezone...)
	if c.HasError() {
		c.Error = fmt.Errorf("%w: %w", ErrMismatchedFormat(value, format), c.Error)
	}
	return c
}

// ParseWithLayouts parses a time string as a Carbon instance with multiple fuzzy layouts.
// 通过多个模糊的 布局模板 将时间字符串解析成 Carbon 实例
func ParseWithLayouts(value string, layouts []string, timezone ...string) *Carbon {
	if value == "" {
		return nil
	}
	if len(layouts) == 0 {
		return Parse(value, timezone...)
	}
	var (
		loc *Location
		err error
	)
	tz := DefaultTimezone
	if len(timezone) > 0 {
		tz = timezone[0]
	}
	loc, err = parseTimezone(tz)
	if err != nil {
		return &Carbon{Error: err}
	}
	c := NewCarbon().SetLocation(loc)
	for i := range layouts {
		if tt, err := time.ParseInLocation(layouts[i], value, c.loc); err == nil {
			c.time = tt
			c.layout = layouts[i]
			return c
		}
	}
	c.Error = ErrFailedParse(value)
	return c
}

// ParseWithFormats parses a time string as a Carbon instance with multiple fuzzy formats.
// 通过多个模糊的 格式模板 将时间字符串解析成 Carbon 实例
func ParseWithFormats(value string, formats []string, timezone ...string) *Carbon {
	if value == "" {
		return nil
	}
	if len(formats) == 0 {
		return Parse(value, timezone...)
	}
	var (
		err error
	)
	tz := DefaultTimezone
	if len(timezone) > 0 {
		tz = timezone[0]
	}
	if _, err = parseTimezone(tz); err != nil {
		return &Carbon{Error: err}
	}
	var layouts []string
	for i := range formats {
		layouts = append(layouts, format2layout(formats[i]))
	}
	return ParseWithLayouts(value, layouts, tz)
}
