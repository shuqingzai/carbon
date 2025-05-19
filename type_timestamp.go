package carbon

import (
	"bytes"
	"database/sql/driver"
	"strconv"
)

// timestamp precision constants
const (
	PrecisionSecond      = 9
	PrecisionMillisecond = 999
	PrecisionMicrosecond = 999999
	PrecisionNanosecond  = 999999999
)

// TimestampTyper defines a TimestampTyper interface.
type TimestampTyper interface {
	~int64 | ~string
	DataType() string
	Precision() int64
}

// TimestampType defines a TimestampType generic struct.
type TimestampType[T TimestampTyper] struct {
	*Carbon
}

// NewTimestampType returns a new TimestampType generic instance.
func NewTimestampType[T TimestampTyper](c *Carbon) *TimestampType[T] {
	return &TimestampType[T]{
		Carbon: c,
	}
}

// Scan implements driver.Scanner interface for TimestampType generic struct.
func (t *TimestampType[T]) Scan(src any) (err error) {
	var (
		ts int64
		c  *Carbon
	)
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		c = Parse(string(v))
		*t = *NewTimestampType[T](c)
		return t.Error
	case string:
		c = Parse(v)
		*t = *NewTimestampType[T](c)
		return t.Error
	case int64:
		ts = v
	case StdTime:
		c = CreateFromStdTime(v, DefaultTimezone)
		*t = *NewTimestampType[T](c)
		return t.Error
	case *StdTime:
		c = CreateFromStdTime(*v, DefaultTimezone)
		*t = *NewTimestampType[T](c)
		return t.Error
	default:
		return ErrFailedScan(src)
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case PrecisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case PrecisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case PrecisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = *NewTimestampType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for TimestampType generic struct.
func (t TimestampType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() || t.IsEmpty() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for TimestampType generic struct.
func (t TimestampType[T]) MarshalJSON() ([]byte, error) {
	if t.IsNil() || t.IsZero() || t.IsEmpty() {
		return []byte(`null`), nil
	}
	if t.HasError() {
		return []byte(`null`), t.Error
	}
	var ts int64
	switch t.getPrecision() {
	case PrecisionSecond:
		ts = t.Timestamp()
	case PrecisionMillisecond:
		ts = t.TimestampMilli()
	case PrecisionMicrosecond:
		ts = t.TimestampMicro()
	case PrecisionNanosecond:
		ts = t.TimestampNano()
	}
	return []byte(strconv.FormatInt(ts, 10)), nil
}

// UnmarshalJSON implements json.Unmarshal interface for TimestampType generic struct.
func (t *TimestampType[T]) UnmarshalJSON(src []byte) error {
	v := string(bytes.Trim(src, `"`))
	if v == "" || v == "null" {
		return nil
	}
	var (
		ts  int64
		err error
	)
	if ts, err = parseTimestamp(v); err != nil {
		return err
	}
	var c *Carbon
	switch t.getPrecision() {
	case PrecisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case PrecisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case PrecisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case PrecisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = *NewTimestampType[T](c)
	return t.Error
}

// String implements Stringer interface for TimestampType generic struct.
func (t *TimestampType[T]) String() string {
	if t == nil || t.IsInvalid() || t.IsZero() {
		return "0"
	}
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
func (t *TimestampType[T]) Int64() (ts int64) {
	if t == nil || t.IsInvalid() || t.IsZero() {
		return
	}
	switch t.getPrecision() {
	case PrecisionSecond:
		ts = t.Timestamp()
	case PrecisionMillisecond:
		ts = t.TimestampMilli()
	case PrecisionMicrosecond:
		ts = t.TimestampMicro()
	case PrecisionNanosecond:
		ts = t.TimestampNano()
	}
	return
}

// GormDataType implements GormDataType interface for TimestampType generic struct.
func (t *TimestampType[T]) GormDataType() string {
	return t.getDataType()
}

// getDataType returns data type of TimestampType generic struct.
func (t *TimestampType[T]) getDataType() string {
	var typer T
	return typer.DataType()
}

// getPrecision returns precision of TimestampType generic struct.
func (t *TimestampType[T]) getPrecision() int64 {
	var typer T
	return typer.Precision()
}
