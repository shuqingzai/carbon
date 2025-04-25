package carbon

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLayoutType_Scan(t *testing.T) {
	c := NewLayoutType[DateTime](Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestLayoutType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := NewLayoutType[DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()
		v, err := NewLayoutType[DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")
		v, err := NewLayoutType[DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")
		v, err := NewLayoutType[DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestLayoutType_MarshalJSON(t *testing.T) {
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

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = NewLayoutType[Date](nil)
		user.DateMilli = NewLayoutType[DateMilli](nil)
		user.DateMicro = NewLayoutType[DateMicro](nil)
		user.DateNano = NewLayoutType[DateNano](nil)

		user.Time = NewLayoutType[Time](nil)
		user.TimeMilli = NewLayoutType[TimeMilli](nil)
		user.TimeMicro = NewLayoutType[TimeMicro](nil)
		user.TimeNano = NewLayoutType[TimeNano](nil)

		user.DateTime = NewLayoutType[DateTime](nil)
		user.DateTimeMilli = NewLayoutType[DateTimeMilli](nil)
		user.DateTimeMicro = NewLayoutType[DateTimeMicro](nil)
		user.DateTimeNano = NewLayoutType[DateTimeNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := NewCarbon()

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

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := Parse("xxx")

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

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
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

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`, string(data))
	})
}

func TestLayoutType_UnmarshalJSON(t *testing.T) {
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

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date":"null","date_milli":"null","date_micro":"null","date_nano":"null","time":"null","time_milli":"null","time_micro":"null","time_nano":"null","date_time":"null","date_time_milli":"null","date_time_micro":"null","date_time_nano":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date":"0","date_milli":"0","date_micro":"0","date_nano":"0","time":"0","time_milli":"0","time_micro":"0","time_nano":"0","date_time":"0","date_time_milli":"0","date_time_micro":"0","date_time_nano":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "2020-08-05", user.Date.String())
		assert.Equal(t, "2020-08-05.999", user.DateMilli.String())
		assert.Equal(t, "2020-08-05.999999", user.DateMicro.String())
		assert.Equal(t, "2020-08-05.999999999", user.DateNano.String())

		assert.Equal(t, "13:14:15", user.Time.String())
		assert.Equal(t, "13:14:15.999", user.TimeMilli.String())
		assert.Equal(t, "13:14:15.999999", user.TimeMicro.String())
		assert.Equal(t, "13:14:15.999999999", user.TimeNano.String())

		assert.Equal(t, "2020-08-05 13:14:15", user.DateTime.String())
		assert.Equal(t, "2020-08-05 13:14:15.999", user.DateTimeMilli.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999", user.DateTimeMicro.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999999", user.DateTimeNano.String())
	})
}

func TestLayoutType_GormDataType(t *testing.T) {
	c := Now()
	dataType := "time"

	assert.Equal(t, dataType, NewLayoutType[Date](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateMilli](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateMicro](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateNano](c).GormDataType())

	assert.Equal(t, dataType, NewLayoutType[Time](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[TimeMilli](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[TimeMicro](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[TimeNano](c).GormDataType())

	assert.Equal(t, dataType, NewLayoutType[DateTime](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateTimeMilli](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateTimeMicro](c).GormDataType())
	assert.Equal(t, dataType, NewLayoutType[DateTimeNano](c).GormDataType())
}

func TestFormatType_Scan(t *testing.T) {
	c := NewFormatType[DateTime](Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestFormatType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := NewFormatType[DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()
		v, err := NewFormatType[DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")
		v, err := NewFormatType[DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")
		v, err := NewFormatType[DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestFormatType_MarshalJSON(t *testing.T) {
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

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = NewFormatType[Date](nil)
		user.DateMilli = NewFormatType[DateMilli](nil)
		user.DateMicro = NewFormatType[DateMicro](nil)
		user.DateNano = NewFormatType[DateNano](nil)

		user.Time = NewFormatType[Time](nil)
		user.TimeMilli = NewFormatType[TimeMilli](nil)
		user.TimeMicro = NewFormatType[TimeMicro](nil)
		user.TimeNano = NewFormatType[TimeNano](nil)

		user.DateTime = NewFormatType[DateTime](nil)
		user.DateTimeMilli = NewFormatType[DateTimeMilli](nil)
		user.DateTimeMicro = NewFormatType[DateTimeMicro](nil)
		user.DateTimeNano = NewFormatType[DateTimeNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := NewCarbon()

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

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := Parse("xxx")

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

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
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

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`, string(data))
	})
}

func TestFormatType_UnmarshalJSON(t *testing.T) {
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

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"date":"null","date_milli":"null","date_micro":"null","date_nano":"null","time":"null","time_milli":"null","time_micro":"null","time_nano":"null","date_time":"null","date_time_milli":"null","date_time_micro":"null","date_time_nano":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"date":"0","date_milli":"0","date_micro":"0","date_nano":"0","time":"0","time_milli":"0","time_micro":"0","time_nano":"0","date_time":"0","date_time_milli":"0","date_time_micro":"0","date_time_nano":"0"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Empty(t, user.Date.String())
		assert.Empty(t, user.DateMilli.String())
		assert.Empty(t, user.DateMicro.String())
		assert.Empty(t, user.DateNano.String())

		assert.Empty(t, user.Time.String())
		assert.Empty(t, user.TimeMilli.String())
		assert.Empty(t, user.TimeMicro.String())
		assert.Empty(t, user.TimeNano.String())

		assert.Empty(t, user.DateTime.String())
		assert.Empty(t, user.DateTimeMilli.String())
		assert.Empty(t, user.DateTimeMicro.String())
		assert.Empty(t, user.DateTimeNano.String())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "2020-08-05", user.Date.String())
		assert.Equal(t, "2020-08-05.999", user.DateMilli.String())
		assert.Equal(t, "2020-08-05.999999", user.DateMicro.String())
		assert.Equal(t, "2020-08-05.999999999", user.DateNano.String())

		assert.Equal(t, "13:14:15", user.Time.String())
		assert.Equal(t, "13:14:15.999", user.TimeMilli.String())
		assert.Equal(t, "13:14:15.999999", user.TimeMicro.String())
		assert.Equal(t, "13:14:15.999999999", user.TimeNano.String())

		assert.Equal(t, "2020-08-05 13:14:15", user.DateTime.String())
		assert.Equal(t, "2020-08-05 13:14:15.999", user.DateTimeMilli.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999", user.DateTimeMicro.String())
		assert.Equal(t, "2020-08-05 13:14:15.999999999", user.DateTimeNano.String())
	})
}

func TestFormatType_GormDataType(t *testing.T) {
	c := Now()
	dataType := "time"

	assert.Equal(t, dataType, NewFormatType[Date](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateMilli](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateMicro](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateNano](c).GormDataType())

	assert.Equal(t, dataType, NewFormatType[Time](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[TimeMilli](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[TimeMicro](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[TimeNano](c).GormDataType())

	assert.Equal(t, dataType, NewFormatType[DateTime](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateTimeMilli](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateTimeMicro](c).GormDataType())
	assert.Equal(t, dataType, NewFormatType[DateTimeNano](c).GormDataType())
}

func TestTimestampType_Scan(t *testing.T) {
	t.Run("[]byte type", func(t *testing.T) {
		ts1 := NewTimestampType[Timestamp](Now())
		assert.Error(t, ts1.Scan([]byte("xxx")))
		assert.Nil(t, ts1.Scan([]byte(strconv.Itoa(int(ts1.Timestamp())))))

		ts2 := NewTimestampType[TimestampMilli](Now())
		assert.Error(t, ts2.Scan([]byte("xxx")))
		assert.Nil(t, ts2.Scan([]byte(strconv.Itoa(int(ts2.Timestamp())))))

		ts3 := NewTimestampType[TimestampMicro](Now())
		assert.Error(t, ts3.Scan([]byte("xxx")))
		assert.Nil(t, ts3.Scan([]byte(strconv.Itoa(int(ts3.Timestamp())))))

		ts4 := NewTimestampType[TimestampNano](Now())
		assert.Error(t, ts4.Scan([]byte("xxx")))
		assert.Nil(t, ts4.Scan([]byte(strconv.Itoa(int(ts4.Timestamp())))))
	})

	t.Run("string type", func(t *testing.T) {
		ts1 := NewTimestampType[Timestamp](Now())
		assert.Error(t, ts1.Scan("xxx"))
		assert.Nil(t, ts1.Scan(strconv.Itoa(int(ts1.Timestamp()))))

		ts2 := NewTimestampType[TimestampMilli](Now())
		assert.Error(t, ts2.Scan("xxx"))
		assert.Nil(t, ts2.Scan(strconv.Itoa(int(ts2.Timestamp()))))

		ts3 := NewTimestampType[TimestampMicro](Now())
		assert.Error(t, ts3.Scan("xxx"))
		assert.Nil(t, ts3.Scan(strconv.Itoa(int(ts3.Timestamp()))))

		ts4 := NewTimestampType[TimestampNano](Now())
		assert.Error(t, ts4.Scan("xxx"))
		assert.Nil(t, ts4.Scan(strconv.Itoa(int(ts4.Timestamp()))))
	})

	t.Run("int64 type", func(t *testing.T) {
		ts1 := NewTimestampType[Timestamp](Now())
		assert.Nil(t, ts1.Scan(Now().Timestamp()))

		ts2 := NewTimestampType[TimestampMilli](Now())
		assert.Nil(t, ts2.Scan(Now().TimestampMilli()))

		ts3 := NewTimestampType[TimestampMicro](Now())
		assert.Nil(t, ts3.Scan(Now().TimestampMicro()))

		ts4 := NewTimestampType[TimestampNano](Now())
		assert.Nil(t, ts4.Scan(Now().TimestampNano()))
	})

	t.Run("time type", func(t *testing.T) {
		ts1 := NewTimestampType[Timestamp](Now())
		assert.Nil(t, ts1.Scan(Now().StdTime()))

		ts2 := NewTimestampType[TimestampMilli](Now())
		assert.Nil(t, ts2.Scan(Now().StdTime()))

		ts3 := NewTimestampType[TimestampMicro](Now())
		assert.Nil(t, ts3.Scan(Now().StdTime()))

		ts4 := NewTimestampType[TimestampNano](Now())
		assert.Nil(t, ts4.Scan(Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		c := Now()

		ts1 := NewTimestampType[Timestamp](c)
		assert.Error(t, ts1.Scan(nil))

		ts2 := NewTimestampType[TimestampMilli](c)
		assert.Error(t, ts2.Scan(nil))

		ts3 := NewTimestampType[TimestampMicro](c)
		assert.Error(t, ts3.Scan(nil))

		ts4 := NewTimestampType[TimestampNano](c)
		assert.Error(t, ts4.Scan(nil))
	})
}

func TestTimestampType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v1, e1 := NewTimestampType[Timestamp](nil).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := NewTimestampType[TimestampMilli](nil).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := NewTimestampType[TimestampMicro](nil).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := NewTimestampType[TimestampNano](nil).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("zero time", func(t *testing.T) {
		c := NewCarbon()

		v1, e1 := NewTimestampType[Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := NewTimestampType[TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := NewTimestampType[TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := NewTimestampType[TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := Parse("xxx")

		v1, e1 := NewTimestampType[Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Error(t, e1)

		v2, e2 := NewTimestampType[TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Error(t, e2)

		v3, e3 := NewTimestampType[TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Error(t, e3)

		v4, e4 := NewTimestampType[TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Error(t, e4)
	})

	t.Run("valid time", func(t *testing.T) {
		c := Parse("2020-08-05")

		v1, e1 := NewTimestampType[Timestamp](c).Value()
		assert.Equal(t, c.Timestamp(), v1)
		assert.Nil(t, e1)

		v2, e2 := NewTimestampType[TimestampMilli](c).Value()
		assert.Equal(t, c.TimestampMilli(), v2)
		assert.Nil(t, e2)

		v3, e3 := NewTimestampType[TimestampMicro](c).Value()
		assert.Equal(t, c.TimestampMicro(), v3)
		assert.Nil(t, e3)

		v4, e4 := NewTimestampType[TimestampNano](c).Value()
		assert.Equal(t, c.TimestampNano(), v4)
		assert.Nil(t, e4)
	})
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      TimestampType[Timestamp]      `json:"timestamp"`
		TimestampMilli TimestampType[TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro TimestampType[TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  TimestampType[TimestampNano]  `json:"timestamp_nano"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Timestamp = NewTimestampType[Timestamp](nil)
		user.TimestampMilli = NewTimestampType[TimestampMilli](nil)
		user.TimestampMicro = NewTimestampType[TimestampMicro](nil)
		user.TimestampNano = NewTimestampType[TimestampNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := NewCarbon()

		user.Timestamp = NewTimestampType[Timestamp](c)
		user.TimestampMilli = NewTimestampType[TimestampMilli](c)
		user.TimestampMicro = NewTimestampType[TimestampMicro](c)
		user.TimestampNano = NewTimestampType[TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := Parse("xxx")

		user.Timestamp = NewTimestampType[Timestamp](c)
		user.TimestampMilli = NewTimestampType[TimestampMilli](c)
		user.TimestampMicro = NewTimestampType[TimestampMicro](c)
		user.TimestampNano = NewTimestampType[TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		now := Parse("2020-08-05 13:14:15")

		user.Timestamp = NewTimestampType[Timestamp](now)
		user.TimestampMilli = NewTimestampType[TimestampMilli](now.AddDay())
		user.TimestampMicro = NewTimestampType[TimestampMicro](now.SubDay())
		user.TimestampNano = NewTimestampType[TimestampNano](now.EndOfDay())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`, string(data))
	})
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      TimestampType[Timestamp]      `json:"timestamp"`
		TimestampMilli TimestampType[TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro TimestampType[TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  TimestampType[TimestampNano]  `json:"timestamp_nano"`
	}

	t.Run("empty value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"","timestamp_milli":"","timestamp_micro":"","timestamp_nano":""}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("null value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"null","timestamp_milli":"null","timestamp_micro":"null","timestamp_nano":"null"}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("zero value", func(t *testing.T) {
		var user User

		value := `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("invalid value", func(t *testing.T) {
		var user User

		value := `{"timestamp":"xxx","timestamp_milli":"xxx","timestamp_micro":"xxx","timestamp_nano":"xxx"}`
		assert.Error(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "0", user.Timestamp.String())
		assert.Equal(t, "0", user.TimestampMilli.String())
		assert.Equal(t, "0", user.TimestampMicro.String())
		assert.Equal(t, "0", user.TimestampNano.String())

		assert.Zero(t, user.Timestamp.Int64())
		assert.Zero(t, user.TimestampMilli.Int64())
		assert.Zero(t, user.TimestampMicro.Int64())
		assert.Zero(t, user.TimestampNano.Int64())
	})

	t.Run("valid value", func(t *testing.T) {
		var user User

		value := `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`
		assert.NoError(t, json.Unmarshal([]byte(value), &user))

		assert.Equal(t, "1596633255", user.Timestamp.String())
		assert.Equal(t, "1596633255000", user.TimestampMilli.String())
		assert.Equal(t, "1596633255000000", user.TimestampMicro.String())
		assert.Equal(t, "1596671999999999999", user.TimestampNano.String())

		assert.Equal(t, int64(1596633255), user.Timestamp.Int64())
		assert.Equal(t, int64(1596633255000), user.TimestampMilli.Int64())
		assert.Equal(t, int64(1596633255000000), user.TimestampMicro.Int64())
		assert.Equal(t, int64(1596671999999999999), user.TimestampNano.Int64())
	})
}

func TestTimestamp_GormDataType(t *testing.T) {
	c := Now()
	dataType := "time"

	assert.Equal(t, dataType, NewTimestampType[Timestamp](c).GormDataType())
	assert.Equal(t, dataType, NewTimestampType[TimestampMilli](c).GormDataType())
	assert.Equal(t, dataType, NewTimestampType[TimestampMicro](c).GormDataType())
	assert.Equal(t, dataType, NewTimestampType[TimestampNano](c).GormDataType())
}

type RFC3339Type string

func (t RFC3339Type) LayoutTemplate() string {
	return RFC3339Layout
}

func TestLayoutType_Customer(t *testing.T) {
	type User struct {
		Customer LayoutType[RFC3339Type] `json:"customer"`
	}

	var user User

	c := Parse("2020-08-05 13:14:15")
	user.Customer = NewLayoutType[RFC3339Type](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15Z"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15Z", person.Customer.String())
}

type ISO8601Type string

func (t ISO8601Type) FormatTemplate() string {
	return ISO8601Format
}

func TestFormatType_Customer(t *testing.T) {
	type User struct {
		Customer FormatType[ISO8601Type] `json:"customer"`
	}

	var user User

	c := Parse("2020-08-05 13:14:15")
	user.Customer = NewFormatType[ISO8601Type](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15+00:00"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15+00:00", person.Customer.String())
}
