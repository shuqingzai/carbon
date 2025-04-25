package carbon

// Timestamp defines a Timestamp type.
// 定义 Timestamp 字段类型
type Timestamp int64

// PrecisionUnit implements TimestampFactory interface for Timestamp type.
// 实现 TimestampFactory 接口
func (t Timestamp) PrecisionUnit() int64 {
	return PrecisionSecond
}

// TimestampMilli defines a TimestampMilli type.
// 定义 TimestampMilli 字段类型
type TimestampMilli int64

// PrecisionUnit implements TimestampFactory interface for TimestampMilli type.
// 实现 TimestampFactory 接口
func (t TimestampMilli) PrecisionUnit() int64 {
	return PrecisionMillisecond
}

// TimestampMicro defines a TimestampMicro type.
// 定义 TimestampMicro 字段类型
type TimestampMicro int64

// PrecisionUnit implements TimestampFactory interface for TimestampMicro type.
// 实现 TimestampFactory 接口
func (t TimestampMicro) PrecisionUnit() int64 {
	return PrecisionMicrosecond
}

// TimestampNano defines a TimestampNano type.
// 定义 TimestampNano 字段类型
type TimestampNano int64

// PrecisionUnit implements TimestampFactory interface for TimestampNano type.
// 实现 TimestampFactory 接口
func (t TimestampNano) PrecisionUnit() int64 {
	return PrecisionNanosecond
}

// DateTime defines a DateTime type.
// 定义 DateTime 字段类型
type DateTime string

// FormatTemplate implements FormatFactory interface for DateTime type.
// 实现 FormatFactory 接口
func (t DateTime) FormatTemplate() string {
	return DateTimeFormat
}

// LayoutTemplate implements LayoutFactory interface for DateTime type.
// 实现 LayoutFactory 接口
func (t DateTime) LayoutTemplate() string {
	return DateTimeLayout
}

// DateTimeMilli defines a DateTimeMilli type.
// 定义 DateTimeMilli 字段类型
type DateTimeMilli string

// FormatTemplate implements FormatFactory interface for DateTimeMilli type.
// 实现 FormatFactory 接口
func (t DateTimeMilli) FormatTemplate() string {
	return DateTimeMilliFormat
}

// LayoutTemplate implements LayoutFactory interface for DateTimeMilli type.
// 实现 LayoutFactory 接口
func (t DateTimeMilli) LayoutTemplate() string {
	return DateTimeMilliLayout
}

// DateTimeMicro defines a DateTimeMicro type.
// 定义 DateTimeMicro 字段类型
type DateTimeMicro string

// FormatTemplate implements FormatFactory interface for DateTimeMicro type.
// 实现 FormatFactory 接口
func (t DateTimeMicro) FormatTemplate() string {
	return DateTimeMicroFormat
}

// LayoutTemplate implements LayoutFactory interface for DateTimeMicro type.
// 实现 LayoutFactory 接口
func (t DateTimeMicro) LayoutTemplate() string {
	return DateTimeMicroLayout
}

// DateTimeNano defines a DateTimeNano type.
// 定义 DateTimeNano 字段类型
type DateTimeNano string

// FormatTemplate implements FormatFactory interface for DateTimeNano type.
// 实现 FormatFactory 接口
func (t DateTimeNano) FormatTemplate() string {
	return DateTimeNanoFormat
}

// LayoutTemplate implements LayoutFactory interface for DateTimeNano type.
// 实现 LayoutFactory 接口
func (t DateTimeNano) LayoutTemplate() string {
	return DateTimeNanoLayout
}

// Date defines a Date type.
// 定义 Date 字段类型
type Date string

// FormatTemplate implements FormatFactory interface for Date type.
// 实现 FormatFactory 接口
func (t Date) FormatTemplate() string {
	return DateFormat
}

// LayoutTemplate implements LayoutFactory interface for Date type.
// 实现 LayoutFactory 接口
func (t Date) LayoutTemplate() string {
	return DateLayout
}

// DateMilli defines a DateMilli type.
// 定义 DateMilli 字段类型
type DateMilli string

// FormatTemplate implements FormatFactory interface for DateMilli type.
// 实现 FormatFactory 接口
func (t DateMilli) FormatTemplate() string {
	return DateMilliFormat
}

// LayoutTemplate implements LayoutFactory interface for DateMilli type.
// 实现 LayoutFactory 接口
func (t DateMilli) LayoutTemplate() string {
	return DateMilliLayout
}

// DateMicro defines a DateMicro type.
// 定义 DateMicro 字段类型
type DateMicro string

// FormatTemplate implements FormatFactory interface for DateMicro type.
// 实现 FormatFactory 接口
func (t DateMicro) FormatTemplate() string {
	return DateMicroFormat
}

// LayoutTemplate implements LayoutFactory interface for DateTimeMicro type.
// 实现 LayoutFactory 接口
func (t DateMicro) LayoutTemplate() string {
	return DateMicroLayout
}

// DateNano defines a DateNano type.
// 定义 DateNano 字段类型
type DateNano string

// FormatTemplate implements FormatFactory interface for DateNano type.
// 实现 FormatFactory 接口
func (t DateNano) FormatTemplate() string {
	return DateNanoFormat
}

// LayoutTemplate implements LayoutFactory interface for DateNano type.
// 实现 LayoutFactory 接口
func (t DateNano) LayoutTemplate() string {
	return DateNanoLayout
}

// Time defines a Time struct.
// 定义 Time 字段类型
type Time string

// FormatTemplate implements FormatFactory interface for Time type.
// 实现 FormatFactory 接口
func (t Time) FormatTemplate() string {
	return TimeFormat
}

// LayoutTemplate implements LayoutFactory interface for Time type.
// 实现 LayoutFactory 接口
func (t Time) LayoutTemplate() string {
	return TimeLayout
}

// TimeMilli defines a TimeMilli type.
// 定义 TimeMilli 字段类型
type TimeMilli string

// FormatTemplate implements FormatFactory interface for TimeMilli type.
// 实现 FormatFactory 接口
func (t TimeMilli) FormatTemplate() string {
	return TimeMilliFormat
}

// LayoutTemplate implements LayoutFactory interface for TimeMilli type.
// 实现 LayoutFactory 接口
func (t TimeMilli) LayoutTemplate() string {
	return TimeMilliLayout
}

// TimeMicro defines a TimeMicro type.
// 定义 TimeMicro 字段类型
type TimeMicro string

// FormatTemplate implements FormatFactory interface for TimeMicro type.
// 实现 FormatFactory 接口
func (t TimeMicro) FormatTemplate() string {
	return TimeMicroFormat
}

// LayoutTemplate implements LayoutFactory interface for TimeMicro type.
// 实现 LayoutFactory 接口
func (t TimeMicro) LayoutTemplate() string {
	return TimeMicroLayout
}

// TimeNano defines a TimeNano type.
// 定义 TimeNano 字段类型
type TimeNano string

// FormatTemplate implements FormatFactory interface for TimeNano type.
// 实现 FormatFactory 接口
func (t TimeNano) FormatTemplate() string {
	return TimeNanoFormat
}

// LayoutTemplate implements LayoutFactory interface for TimeNano type.
// 实现 LayoutFactory 接口
func (t TimeNano) LayoutTemplate() string {
	return TimeNanoLayout
}
