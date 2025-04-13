package carbon

import (
	"encoding/json"
	"testing"
)

func BenchmarkLayoutType_Scan(b *testing.B) {
	c := NewLayoutType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_ = c.Scan(Now().StdTime())
	}
}

func BenchmarkLayoutType_Value(b *testing.B) {
	c := NewLayoutType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_, _ = c.Value()
	}
}

func BenchmarkLayoutType_MarshalJSON(b *testing.B) {
	type User struct {
		Date      LayoutType[Date]      `json:"date"`
		DateMilli LayoutType[DateMilli] `json:"date_milli"`
		DateMicro LayoutType[DateMicro] `json:"date_micro"`
		DateNano  LayoutType[DateNano]  `json:"date_nano"`

		Time      LayoutType[Time]      `json:"time"`
		TimeMilli LayoutType[TimeMilli] `json:"time_milli"`
		TimeMicro LayoutType[TimeMicro] `json:"time_micro"`
		TimeNano  LayoutType[TimeNano]  `json:"time_nano"`

		DateTime      LayoutType[DateTime]      `json:"date_time"`
		DateTimeMilli LayoutType[DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro LayoutType[DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  LayoutType[DateTimeNano]  `json:"date_time_nano"`
	}

	var user User

	c := Parse("2020-08-05 13:14:15.999999999")

	user.Date = NewLayoutType[Date](c)
	user.DateMilli = NewLayoutType[DateMilli](c)
	user.DateMicro = NewLayoutType[DateMicro](c)
	user.DateNano = NewLayoutType[DateNano](c)

	user.Time = NewLayoutType[Time](c)
	user.TimeMilli = NewLayoutType[TimeMilli](c)
	user.TimeMicro = NewLayoutType[TimeMicro](c)
	user.TimeNano = NewLayoutType[TimeNano](c)

	user.DateTime = NewLayoutType[DateTime](c)
	user.DateTimeMilli = NewLayoutType[DateTimeMilli](c)
	user.DateTimeMicro = NewLayoutType[DateTimeMicro](c)
	user.DateTimeNano = NewLayoutType[DateTimeNano](c)

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(&user)
	}
}

func BenchmarkLayoutType_UnmarshalJSON(b *testing.B) {
	type User struct {
		Date      LayoutType[Date]      `json:"date"`
		DateMilli LayoutType[DateMilli] `json:"date_milli"`
		DateMicro LayoutType[DateMicro] `json:"date_micro"`
		DateNano  LayoutType[DateNano]  `json:"date_nano"`

		Time      LayoutType[Time]      `json:"time"`
		TimeMilli LayoutType[TimeMilli] `json:"time_milli"`
		TimeMicro LayoutType[TimeMicro] `json:"time_micro"`
		TimeNano  LayoutType[TimeNano]  `json:"time_nano"`

		DateTime      LayoutType[DateTime]      `json:"date_time"`
		DateTimeMilli LayoutType[DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro LayoutType[DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  LayoutType[DateTimeNano]  `json:"date_time_nano"`
	}

	var user User

	value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`

	for n := 0; n < b.N; n++ {
		_ = json.Unmarshal([]byte(value), &user)
	}
}

func BenchmarkLayoutType_String(b *testing.B) {
	c := NewLayoutType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_ = c.String()
	}
}

func BenchmarkLayoutType_GormDataType(b *testing.B) {
	c := NewLayoutType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		c.GormDataType()
	}
}

func BenchmarkFormatType_Scan(b *testing.B) {
	c := NewFormatType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_ = c.Scan(Now().StdTime())
	}
}

func BenchmarkFormatType_Value(b *testing.B) {
	c := NewFormatType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_, _ = c.Value()
	}
}

func BenchmarkFormatType_MarshalJSON(b *testing.B) {
	type User struct {
		Date      FormatType[Date]      `json:"date"`
		DateMilli FormatType[DateMilli] `json:"date_milli"`
		DateMicro FormatType[DateMicro] `json:"date_micro"`
		DateNano  FormatType[DateNano]  `json:"date_nano"`

		Time      FormatType[Time]      `json:"time"`
		TimeMilli FormatType[TimeMilli] `json:"time_milli"`
		TimeMicro FormatType[TimeMicro] `json:"time_micro"`
		TimeNano  FormatType[TimeNano]  `json:"time_nano"`

		DateTime      FormatType[DateTime]      `json:"date_time"`
		DateTimeMilli FormatType[DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro FormatType[DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  FormatType[DateTimeNano]  `json:"date_time_nano"`
	}

	var user User

	c := Parse("2020-08-05 13:14:15.999999999")

	user.Date = NewFormatType[Date](c)
	user.DateMilli = NewFormatType[DateMilli](c)
	user.DateMicro = NewFormatType[DateMicro](c)
	user.DateNano = NewFormatType[DateNano](c)

	user.Time = NewFormatType[Time](c)
	user.TimeMilli = NewFormatType[TimeMilli](c)
	user.TimeMicro = NewFormatType[TimeMicro](c)
	user.TimeNano = NewFormatType[TimeNano](c)

	user.DateTime = NewFormatType[DateTime](c)
	user.DateTimeMilli = NewFormatType[DateTimeMilli](c)
	user.DateTimeMicro = NewFormatType[DateTimeMicro](c)
	user.DateTimeNano = NewFormatType[DateTimeNano](c)

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(&user)
	}
}

func BenchmarkFormatType_UnmarshalJSON(b *testing.B) {
	type User struct {
		Date      FormatType[Date]      `json:"date"`
		DateMilli FormatType[DateMilli] `json:"date_milli"`
		DateMicro FormatType[DateMicro] `json:"date_micro"`
		DateNano  FormatType[DateNano]  `json:"date_nano"`

		Time      FormatType[Time]      `json:"time"`
		TimeMilli FormatType[TimeMilli] `json:"time_milli"`
		TimeMicro FormatType[TimeMicro] `json:"time_micro"`
		TimeNano  FormatType[TimeNano]  `json:"time_nano"`

		DateTime      FormatType[DateTime]      `json:"date_time"`
		DateTimeMilli FormatType[DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro FormatType[DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  FormatType[DateTimeNano]  `json:"date_time_nano"`
	}

	var user User

	value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`

	for n := 0; n < b.N; n++ {
		_ = json.Unmarshal([]byte(value), &user)
	}
}

func BenchmarkFormatType_String(b *testing.B) {
	c := NewFormatType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		_ = c.String()
	}
}

func BenchmarkFormatType_GormDataType(b *testing.B) {
	c := NewFormatType[DateTime](Now())
	for n := 0; n < b.N; n++ {
		c.GormDataType()
	}
}

func BenchmarkTimestampType_Scan(b *testing.B) {
	c := NewTimestampType[Timestamp](Now())
	for n := 0; n < b.N; n++ {
		_ = c.Scan(Now().StdTime())
	}
}

func BenchmarkTimestampType_Value(b *testing.B) {
	c := NewTimestampType[Timestamp](Now())
	for n := 0; n < b.N; n++ {
		_, _ = c.Value()
	}
}

func BenchmarkTimestampType_MarshalJSON(b *testing.B) {
	type User struct {
		Timestamp      TimestampType[Timestamp]      `json:"timestamp"`
		TimestampMilli TimestampType[TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro TimestampType[TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  TimestampType[TimestampNano]  `json:"timestamp_nano"`
	}

	var user User

	now := Parse("2020-08-05 13:14:15")

	user.Timestamp = NewTimestampType[Timestamp](now)
	user.TimestampMilli = NewTimestampType[TimestampMilli](now.AddDay())
	user.TimestampMicro = NewTimestampType[TimestampMicro](now.SubDay())
	user.TimestampNano = NewTimestampType[TimestampNano](now.EndOfDay())

	for n := 0; n < b.N; n++ {
		_, _ = json.Marshal(&user)
	}
}

func BenchmarkTimestampType_UnmarshalJSON(b *testing.B) {
	type User struct {
		Timestamp      TimestampType[Timestamp]      `json:"timestamp"`
		TimestampMilli TimestampType[TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro TimestampType[TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  TimestampType[TimestampNano]  `json:"timestamp_nano"`
	}

	var user User

	value := `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`

	for n := 0; n < b.N; n++ {
		_ = json.Unmarshal([]byte(value), &user)
	}
}

func BenchmarkTimestampType_String(b *testing.B) {
	c := NewTimestampType[Timestamp](Now())
	for n := 0; n < b.N; n++ {
		_ = c.String()
	}
}

func BenchmarkTimestampType_GormDataType(b *testing.B) {
	c := NewTimestampType[Timestamp](Now())
	for n := 0; n < b.N; n++ {
		c.GormDataType()
	}
}
