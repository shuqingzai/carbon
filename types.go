package carbon

// Timestamp defines a Timestamp type.
// 定义 Timestamp 字段类型
type Timestamp int64

// SetPrecision implements TimestampFactory interface for Timestamp type.
// 实现 TimestampFactory 接口
func (t Timestamp) SetPrecision() int64 {
	return PrecisionSecond
}

// TimestampMilli defines a TimestampMilli type.
// 定义 TimestampMilli 字段类型
type TimestampMilli int64

// SetPrecision implements TimestampFactory interface for TimestampMilli type.
// 实现 TimestampFactory 接口
func (t TimestampMilli) SetPrecision() int64 {
	return PrecisionMillisecond
}

// TimestampMicro defines a TimestampMicro type.
// 定义 TimestampMicro 字段类型
type TimestampMicro int64

// SetPrecision implements TimestampFactory interface for TimestampMicro type.
// 实现 TimestampFactory 接口
func (t TimestampMicro) SetPrecision() int64 {
	return PrecisionMicrosecond
}

// TimestampNano defines a TimestampNano type.
// 定义 TimestampNano 字段类型
type TimestampNano int64

// SetPrecision implements TimestampFactory interface for TimestampNano type.
// 实现 TimestampFactory 接口
func (t TimestampNano) SetPrecision() int64 {
	return PrecisionNanosecond
}

// DateTime defines a DateTime type.
// 定义 DateTime 字段类型
type DateTime string

// SetFormat implements FormatFactory interface for DateTime type.
// 实现 FormatFactory 接口
func (t DateTime) SetFormat() string {
	return DateTimeFormat
}

// SetLayout implements LayoutFactory interface for DateTime type.
// 实现 LayoutFactory 接口
func (t DateTime) SetLayout() string {
	return DateTimeLayout
}

// DateTimeMilli defines a DateTimeMilli type.
// 定义 DateTimeMilli 字段类型
type DateTimeMilli string

// SetFormat implements FormatFactory interface for DateTimeMilli type.
// 实现 FormatFactory 接口
func (t DateTimeMilli) SetFormat() string {
	return DateTimeMilliFormat
}

// SetLayout implements LayoutFactory interface for DateTimeMilli type.
// 实现 LayoutFactory 接口
func (t DateTimeMilli) SetLayout() string {
	return DateTimeMilliLayout
}

// DateTimeMicro defines a DateTimeMicro type.
// 定义 DateTimeMicro 字段类型
type DateTimeMicro string

// SetFormat implements FormatFactory interface for DateTimeMicro type.
// 实现 FormatFactory 接口
func (t DateTimeMicro) SetFormat() string {
	return DateTimeMicroFormat
}

// SetLayout implements LayoutFactory interface for DateTimeMicro type.
// 实现 LayoutFactory 接口
func (t DateTimeMicro) SetLayout() string {
	return DateTimeMicroLayout
}

// DateTimeNano defines a DateTimeNano type.
// 定义 DateTimeNano 字段类型
type DateTimeNano string

// SetFormat implements FormatFactory interface for DateTimeNano type.
// 实现 FormatFactory 接口
func (t DateTimeNano) SetFormat() string {
	return DateTimeNanoFormat
}

// SetLayout implements LayoutFactory interface for DateTimeNano type.
// 实现 LayoutFactory 接口
func (t DateTimeNano) SetLayout() string {
	return DateTimeNanoLayout
}

// Date defines a Date type.
// 定义 Date 字段类型
type Date string

// SetFormat implements FormatFactory interface for Date type.
// 实现 FormatFactory 接口
func (t Date) SetFormat() string {
	return DateFormat
}

// SetLayout implements LayoutFactory interface for Date type.
// 实现 LayoutFactory 接口
func (t Date) SetLayout() string {
	return DateLayout
}

// DateMilli defines a DateMilli type.
// 定义 DateMilli 字段类型
type DateMilli string

// SetFormat implements FormatFactory interface for DateMilli type.
// 实现 FormatFactory 接口
func (t DateMilli) SetFormat() string {
	return DateMilliFormat
}

// SetLayout implements LayoutFactory interface for DateMilli type.
// 实现 LayoutFactory 接口
func (t DateMilli) SetLayout() string {
	return DateMilliLayout
}

// DateMicro defines a DateMicro type.
// 定义 DateMicro 字段类型
type DateMicro string

// SetFormat implements FormatFactory interface for DateMicro type.
// 实现 FormatFactory 接口
func (t DateMicro) SetFormat() string {
	return DateMicroFormat
}

// SetLayout implements LayoutFactory interface for DateTimeMicro type.
// 实现 LayoutFactory 接口
func (t DateMicro) SetLayout() string {
	return DateMicroLayout
}

// DateNano defines a DateNano type.
// 定义 DateNano 字段类型
type DateNano string

// SetFormat implements FormatFactory interface for DateNano type.
// 实现 FormatFactory 接口
func (t DateNano) SetFormat() string {
	return DateNanoFormat
}

// SetLayout implements LayoutFactory interface for DateNano type.
// 实现 LayoutFactory 接口
func (t DateNano) SetLayout() string {
	return DateNanoLayout
}

// Time defines a Time struct.
// 定义 Time 字段类型
type Time string

// SetFormat implements FormatFactory interface for Time type.
// 实现 FormatFactory 接口
func (t Time) SetFormat() string {
	return TimeFormat
}

// SetLayout implements LayoutFactory interface for Time type.
// 实现 LayoutFactory 接口
func (t Time) SetLayout() string {
	return TimeLayout
}

// TimeMilli defines a TimeMilli type.
// 定义 TimeMilli 字段类型
type TimeMilli string

// SetFormat implements FormatFactory interface for TimeMilli type.
// 实现 FormatFactory 接口
func (t TimeMilli) SetFormat() string {
	return TimeMilliFormat
}

// SetLayout implements LayoutFactory interface for TimeMilli type.
// 实现 LayoutFactory 接口
func (t TimeMilli) SetLayout() string {
	return TimeMilliLayout
}

// TimeMicro defines a TimeMicro type.
// 定义 TimeMicro 字段类型
type TimeMicro string

// SetFormat implements FormatFactory interface for TimeMicro type.
// 实现 FormatFactory 接口
func (t TimeMicro) SetFormat() string {
	return TimeMicroFormat
}

// SetLayout implements LayoutFactory interface for TimeMicro type.
// 实现 LayoutFactory 接口
func (t TimeMicro) SetLayout() string {
	return TimeMicroLayout
}

// TimeNano defines a TimeNano type.
// 定义 TimeNano 字段类型
type TimeNano string

// SetFormat implements FormatFactory interface for TimeNano type.
// 实现 FormatFactory 接口
func (t TimeNano) SetFormat() string {
	return TimeNanoFormat
}

// SetLayout implements LayoutFactory interface for TimeNano type.
// 实现 LayoutFactory 接口
func (t TimeNano) SetLayout() string {
	return TimeNanoLayout
}
