package carbon

import (
	"bytes"
	"database/sql/driver"
	"time"
)

// Scan implements driver.Scanner interface for Carbon struct.
// 实现 driver.Scanner 接口
func (c *Carbon) Scan(src any) error {
	switch v := src.(type) {
	case nil:
		return nil
	case []byte:
		*c = *Parse(string(v), DefaultTimezone)
	case string:
		*c = *Parse(v, DefaultTimezone)
	case time.Time:
		*c = *CreateFromStdTime(v, DefaultTimezone)
	case int64:
		*c = *CreateFromTimestamp(v, DefaultTimezone)
	default:
		return ErrFailedScan(v)
	}
	return c.Error
}

// Value implements driver.Valuer interface for Carbon struct.
// 实现 driver.Valuer 接口
func (c *Carbon) Value() (driver.Value, error) {
	if c.IsNil() || c.IsZero() {
		return nil, nil
	}
	if c.HasError() {
		return nil, c.Error
	}
	return c.StdTime(), nil
}

// MarshalJSON implements json.Marshal interface for Carbon struct.
// 实现 json.Marshaler 接口
func (c *Carbon) MarshalJSON() ([]byte, error) {
	if c.IsNil() || c.IsZero() {
		return []byte(`""`), nil
	}
	if c.HasError() {
		return []byte(`""`), c.Error
	}
	value := c.Layout(DefaultLayout, c.Timezone())
	bs := make([]byte, 0, len(value)+2)
	bs = append(bs, '"')
	bs = append(bs, value...)
	bs = append(bs, '"')
	return bs, nil
}

// UnmarshalJSON implements json.Unmarshal interface for Carbon struct.
// 实现 json.Unmarshaler 接口
func (c *Carbon) UnmarshalJSON(bs []byte) error {
	value := string(bytes.Trim(bs, `"`))
	if value == "" || value == "null" || value == "0" {
		return nil
	}
	*c = *ParseByLayout(value, DefaultLayout)
	return c.Error
}

// String implements the interface Stringer for Carbon struct.
// 实现 Stringer 接口
func (c *Carbon) String() string {
	if c == nil || c.IsInvalid() || c.IsZero() {
		return ""
	}
	return c.Layout(c.layout, c.Timezone())
}

// GormDataType sets gorm data type for Carbon struct.
// 设置 gorm 数据类型
func (c *Carbon) GormDataType() string {
	return "time"
}
