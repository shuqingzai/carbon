package carbon

import (
	"bytes"
	"database/sql/driver"
	"strconv"
	"time"
)

// timestamp precision constants
// 时间戳精度常量
const (
	precisionSecond      = 9
	precisionMillisecond = 999
	precisionMicrosecond = 999999
	precisionNanosecond  = 999999999
)

// defines a timestampFactory interface.
// 定义 timestampFactory 接口
type timestampFactory interface {
	~int64
	precision() int64
}

// defines a timestampType generic struct.
// 定义 timestampType 泛型结构体
type timestampType[T timestampFactory] struct {
	*Carbon
}

// returns a new timestampType generic instance.
// 返回 timestampType 泛型实例
func newTimestampType[T timestampFactory](c *Carbon) *timestampType[T] {
	return &timestampType[T]{
		Carbon: c,
	}
}

// Scan implements driver.Scanner interface for timestampType generic struct.
// 实现 driver.Scanner 接口
func (t *timestampType[T]) Scan(src any) (err error) {
	var (
		ts int64
		c  *Carbon
	)
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		if ts, err = parseTimestamp(string(v)); err != nil {
			return err
		}
	case string:
		if ts, err = parseTimestamp(v); err != nil {
			return err
		}
	case int64:
		ts = v
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
		*t = *newTimestampType[T](c)
		return t.Error
	default:
		return ErrFailedScan(src)
	}
	switch t.getPrecision() {
	case precisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case precisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case precisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case precisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = *newTimestampType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for timestampType generic struct.
// 实现 driver.Valuer 接口
func (t *timestampType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	var v int64
	switch t.getPrecision() {
	case precisionSecond:
		v = t.Timestamp()
	case precisionMillisecond:
		v = t.TimestampMilli()
	case precisionMicrosecond:
		v = t.TimestampMicro()
	case precisionNanosecond:
		v = t.TimestampNano()
	}
	return v, nil
}

// MarshalJSON implements json.Marshal interface for TimestampType generic struct.
// 实现 json.Marshaler 接口
func (t *timestampType[T]) MarshalJSON() ([]byte, error) {
	if t.IsNil() || t.IsZero() {
		return []byte(`0`), nil
	}
	if t.HasError() {
		return []byte(`0`), t.Error
	}
	var ts int64
	switch t.getPrecision() {
	case precisionSecond:
		ts = t.Timestamp()
	case precisionMillisecond:
		ts = t.TimestampMilli()
	case precisionMicrosecond:
		ts = t.TimestampMicro()
	case precisionNanosecond:
		ts = t.TimestampNano()
	}
	return []byte(strconv.FormatInt(ts, 10)), nil
}

// UnmarshalJSON implements json.Unmarshal interface for timestampType generic struct.
// 实现 json.Unmarshaler 接口
func (t *timestampType[T]) UnmarshalJSON(bs []byte) error {
	value := string(bytes.Trim(bs, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	ts, err := parseTimestamp(value)
	if err != nil {
		return err
	}
	var c *Carbon
	switch t.getPrecision() {
	case precisionSecond:
		c = CreateFromTimestamp(ts, DefaultTimezone)
	case precisionMillisecond:
		c = CreateFromTimestampMilli(ts, DefaultTimezone)
	case precisionMicrosecond:
		c = CreateFromTimestampMicro(ts, DefaultTimezone)
	case precisionNanosecond:
		c = CreateFromTimestampNano(ts, DefaultTimezone)
	}
	*t = *newTimestampType[T](c)
	return t.Error
}

// String implements Stringer interface for timestampType generic struct.
// 实现 Stringer 接口
func (t *timestampType[T]) String() string {
	if t.IsInvalid() || t.IsZero() {
		return "0"
	}
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t *timestampType[T]) Int64() (ts int64) {
	if t.IsInvalid() || t.IsZero() {
		return ts
	}
	switch t.getPrecision() {
	case precisionSecond:
		ts = t.Timestamp()
	case precisionMillisecond:
		ts = t.TimestampMilli()
	case precisionMicrosecond:
		ts = t.TimestampMicro()
	case precisionNanosecond:
		ts = t.TimestampNano()
	}
	return
}

// GormDataType sets gorm data type for timestampType generic struct.
// 设置 gorm 数据类型
func (t *timestampType[T]) GormDataType() string {
	return "time"
}

// getPrecision returns the set timestamp precision.
// 返回设置的时间戳精度
func (t *timestampType[T]) getPrecision() int64 {
	var factory T
	return factory.precision()
}
