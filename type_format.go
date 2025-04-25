package carbon

import (
	"bytes"
	"database/sql/driver"
	"time"
)

// FormatFactory defines a FormatFactory interface.
// 定义 FormatFactory 接口
type FormatFactory interface {
	~string
	Format() string
}

// FormatType defines a FormatType generic struct.
// 定义 FormatType 泛型结构体
type FormatType[T FormatFactory] struct {
	*Carbon
}

// NewFormatType returns a new FormatType generic instance.
// 返回 FormatType 泛型实例
func NewFormatType[T FormatFactory](carbon *Carbon) *FormatType[T] {
	return &FormatType[T]{
		Carbon: carbon,
	}
}

// Scan implements driver.Scanner interface for FormatType generic struct.
// 实现 driver.Scanner 接口
func (t *FormatType[T]) Scan(src any) error {
	var c *Carbon
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		c = Parse(string(v), DefaultTimezone)
	case string:
		c = Parse(v, DefaultTimezone)
	case time.Time:
		c = CreateFromStdTime(v, DefaultTimezone)
	case int64:
		c = CreateFromTimestamp(v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	*t = *NewFormatType[T](c)
	return t.Error
}

// Value implements driver.Valuer interface for FormatType generic struct.
// 实现 driver.Valuer 接口
func (t *FormatType[T]) Value() (driver.Value, error) {
	if t.IsNil() || t.IsZero() {
		return nil, nil
	}
	if t.HasError() {
		return nil, t.Error
	}
	return t.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for FormatType generic struct.
// 实现 json.Marshaler 接口
func (t *FormatType[T]) MarshalJSON() ([]byte, error) {
	if t.IsNil() || t.IsZero() {
		return []byte(`""`), nil
	}
	if t.HasError() {
		return []byte(`""`), t.Error
	}
	value := t.Format(t.getFormat(), t.Timezone())
	bs := make([]byte, 0, len(value)+2)
	bs = append(bs, '"')
	bs = append(bs, value...)
	bs = append(bs, '"')
	return bs, nil
}

// UnmarshalJSON implements json.Unmarshal interface for FormatType generic struct.
// 实现 json.Unmarshaler 接口
func (t *FormatType[T]) UnmarshalJSON(bs []byte) error {
	value := string(bytes.Trim(bs, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	*t = *NewFormatType[T](ParseByFormat(value, t.getFormat()))
	return t.Error
}

// String implements Stringer interface for FormatType generic struct.
// 实现 Stringer 接口
func (t *FormatType[T]) String() string {
	if t.IsInvalid() || t.IsZero() {
		return ""
	}
	return t.Format(t.getFormat(), t.Timezone())
}

// GormDataType sets gorm data type for FormatType generic struct.
// 设置 gorm 数据类型
func (t *FormatType[T]) GormDataType() string {
	return "time"
}

// getFormat returns the set format.
// 返回设置的格式模板
func (t *FormatType[T]) getFormat() string {
	var factory T
	return factory.Format()
}
