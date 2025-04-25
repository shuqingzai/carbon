package carbon

import (
	"bytes"
	"database/sql/driver"
	"fmt"
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
	setPrecision() int64
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
	ts := int64(0)
	c := NewCarbon()
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		ts, err = strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return ErrInvalidTimestamp(string(v))
		}
	case string:
		ts, err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return ErrInvalidTimestamp(v)
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
	v := int64(0)
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
	ts := int64(0)
	if t.IsNil() || t.IsZero() {
		return []byte(fmt.Sprintf(`%d`, ts)), nil
	}
	if t.HasError() {
		return []byte(fmt.Sprintf(`%d`, ts)), t.Error
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
	return []byte(fmt.Sprintf(`%d`, ts)), nil
}

// UnmarshalJSON implements json.Unmarshal interface for timestampType generic struct.
// 实现 json.Unmarshaler 接口
func (t *timestampType[T]) UnmarshalJSON(b []byte) error {
	value := string(bytes.Trim(b, `"`))
	println("value", value)
	c := NewCarbon()
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	ts, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return ErrInvalidTimestamp(value)
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

// String implements Stringer interface for timestampType generic struct.
// 实现 Stringer 接口
func (t *timestampType[T]) String() string {
	if t == nil || t.IsInvalid() || t.IsZero() {
		return "0"
	}
	return strconv.FormatInt(t.Int64(), 10)
}

// Int64 returns the timestamp value.
// 返回时间戳
func (t *timestampType[T]) Int64() int64 {
	ts := int64(0)
	if t == nil || t.IsInvalid() || t.IsZero() {
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
	return ts
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
	return factory.setPrecision()
}
