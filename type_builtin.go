package carbon

type timestampSecondType int64

func (t timestampSecondType) setPrecision() int64 {
	return precisionSecond
}

type Timestamp = timestampType[timestampSecondType]

func NewTimestamp(c *Carbon) *Timestamp {
	return newTimestampType[timestampSecondType](c)
}

type timestampMilliType int64

func (t timestampMilliType) setPrecision() int64 {
	return precisionMillisecond
}

type TimestampMilli = timestampType[timestampMilliType]

func NewTimestampMilli(c *Carbon) *TimestampMilli {
	return newTimestampType[timestampMilliType](c)
}

type timestampMicroType int64

func (t timestampMicroType) setPrecision() int64 {
	return precisionMicrosecond
}

type TimestampMicro = timestampType[timestampMicroType]

func NewTimestampMicro(c *Carbon) *TimestampMicro {
	return newTimestampType[timestampMicroType](c)
}

type timestampNanoType int64

func (t timestampNanoType) setPrecision() int64 {
	return precisionNanosecond
}

type TimestampNano = timestampType[timestampNanoType]

func NewTimestampNano(c *Carbon) *TimestampNano {
	return newTimestampType[timestampNanoType](c)
}

type DateTimeType string

func (t DateTimeType) SetLayout() string {
	return DateTimeLayout
}

type DateTime = LayoutType[DateTimeType]

func NewDateTime(c *Carbon) *DateTime {
	return NewLayoutType[DateTimeType](c)
}

type DateTimeMilliType string

func (t DateTimeMilliType) SetLayout() string {
	return DateTimeMilliLayout
}

type DateTimeMilli = LayoutType[DateTimeMilliType]

func NewDateTimeMilli(c *Carbon) *DateTimeMilli {
	return NewLayoutType[DateTimeMilliType](c)
}

type DateTimeMicroType string

func (t DateTimeMicroType) SetLayout() string {
	return DateTimeMicroLayout
}

type DateTimeMicro = LayoutType[DateTimeMicroType]

func NewDateTimeMicro(c *Carbon) *DateTimeMicro {
	return NewLayoutType[DateTimeMicroType](c)
}

type DateTimeNanoType string

func (t DateTimeNanoType) SetLayout() string {
	return DateTimeNanoLayout
}

type DateTimeNano = LayoutType[DateTimeNanoType]

func NewDateTimeNano(c *Carbon) *DateTimeNano {
	return NewLayoutType[DateTimeNanoType](c)
}

type DateType string

func (t DateType) SetLayout() string {
	return DateLayout
}

type Date = LayoutType[DateType]

func NewDate(c *Carbon) *Date {
	return NewLayoutType[DateType](c)
}

type DateMilliType string

func (t DateMilliType) SetLayout() string {
	return DateMilliLayout
}

type DateMilli = LayoutType[DateMilliType]

func NewDateMilli(c *Carbon) *DateMilli {
	return NewLayoutType[DateMilliType](c)
}

type DateMicroType string

func (t DateMicroType) SetLayout() string {
	return DateMicroLayout
}

type DateMicro = LayoutType[DateMicroType]

func NewDateMicro(c *Carbon) *DateMicro {
	return NewLayoutType[DateMicroType](c)
}

type DateNanoType string

func (t DateNanoType) SetLayout() string {
	return DateNanoLayout
}

type DateNano = LayoutType[DateNanoType]

func NewDateNano(c *Carbon) *DateNano {
	return NewLayoutType[DateNanoType](c)
}

type TimeType string

func (t TimeType) SetLayout() string {
	return TimeLayout
}

type Time = LayoutType[TimeType]

func NewTime(c *Carbon) *Time {
	return NewLayoutType[TimeType](c)
}

type TimeMilliType string

func (t TimeMilliType) SetLayout() string {
	return TimeMilliLayout
}

type TimeMilli = LayoutType[TimeMilliType]

func NewTimeMilli(c *Carbon) *TimeMilli {
	return NewLayoutType[TimeMilliType](c)
}

type TimeMicroType string

func (t TimeMicroType) SetLayout() string {
	return TimeMicroLayout
}

type TimeMicro = LayoutType[TimeMicroType]

func NewTimeMicro(c *Carbon) *TimeMicro {
	return NewLayoutType[TimeMicroType](c)
}

type TimeNanoType string

func (t TimeNanoType) SetLayout() string {
	return TimeNanoLayout
}

type TimeNano = LayoutType[TimeNanoType]

func NewTimeNano(c *Carbon) *TimeNano {
	return NewLayoutType[TimeNanoType](c)
}
