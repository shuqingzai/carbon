package carbon_test

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dromara/carbon/v2"
)

func TestLayoutType_Scan(t *testing.T) {
	c := carbon.NewLayoutType[carbon.DateTime](carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(carbon.Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestLayoutType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := carbon.NewLayoutType[carbon.DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewLayoutType[carbon.DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestLayoutType_MarshalJSON(t *testing.T) {
	type User struct {
		Date      carbon.LayoutType[carbon.Date]      `json:"date"`
		DateMilli carbon.LayoutType[carbon.DateMilli] `json:"date_milli"`
		DateMicro carbon.LayoutType[carbon.DateMicro] `json:"date_micro"`
		DateNano  carbon.LayoutType[carbon.DateNano]  `json:"date_nano"`

		Time      carbon.LayoutType[carbon.Time]      `json:"time"`
		TimeMilli carbon.LayoutType[carbon.TimeMilli] `json:"time_milli"`
		TimeMicro carbon.LayoutType[carbon.TimeMicro] `json:"time_micro"`
		TimeNano  carbon.LayoutType[carbon.TimeNano]  `json:"time_nano"`

		DateTime      carbon.LayoutType[carbon.DateTime]      `json:"date_time"`
		DateTimeMilli carbon.LayoutType[carbon.DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro carbon.LayoutType[carbon.DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  carbon.LayoutType[carbon.DateTimeNano]  `json:"date_time_nano"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = carbon.NewLayoutType[carbon.Date](nil)
		user.DateMilli = carbon.NewLayoutType[carbon.DateMilli](nil)
		user.DateMicro = carbon.NewLayoutType[carbon.DateMicro](nil)
		user.DateNano = carbon.NewLayoutType[carbon.DateNano](nil)

		user.Time = carbon.NewLayoutType[carbon.Time](nil)
		user.TimeMilli = carbon.NewLayoutType[carbon.TimeMilli](nil)
		user.TimeMicro = carbon.NewLayoutType[carbon.TimeMicro](nil)
		user.TimeNano = carbon.NewLayoutType[carbon.TimeNano](nil)

		user.DateTime = carbon.NewLayoutType[carbon.DateTime](nil)
		user.DateTimeMilli = carbon.NewLayoutType[carbon.DateTimeMilli](nil)
		user.DateTimeMicro = carbon.NewLayoutType[carbon.DateTimeMicro](nil)
		user.DateTimeNano = carbon.NewLayoutType[carbon.DateTimeNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()

		user.Date = carbon.NewLayoutType[carbon.Date](c)
		user.DateMilli = carbon.NewLayoutType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewLayoutType[carbon.DateMicro](c)
		user.DateNano = carbon.NewLayoutType[carbon.DateNano](c)

		user.Time = carbon.NewLayoutType[carbon.Time](c)
		user.TimeMilli = carbon.NewLayoutType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewLayoutType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewLayoutType[carbon.TimeNano](c)

		user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewLayoutType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewLayoutType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewLayoutType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Date = carbon.NewLayoutType[carbon.Date](c)
		user.DateMilli = carbon.NewLayoutType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewLayoutType[carbon.DateMicro](c)
		user.DateNano = carbon.NewLayoutType[carbon.DateNano](c)

		user.Time = carbon.NewLayoutType[carbon.Time](c)
		user.TimeMilli = carbon.NewLayoutType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewLayoutType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewLayoutType[carbon.TimeNano](c)

		user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewLayoutType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewLayoutType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewLayoutType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")

		user.Date = carbon.NewLayoutType[carbon.Date](c)
		user.DateMilli = carbon.NewLayoutType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewLayoutType[carbon.DateMicro](c)
		user.DateNano = carbon.NewLayoutType[carbon.DateNano](c)

		user.Time = carbon.NewLayoutType[carbon.Time](c)
		user.TimeMilli = carbon.NewLayoutType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewLayoutType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewLayoutType[carbon.TimeNano](c)

		user.DateTime = carbon.NewLayoutType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewLayoutType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewLayoutType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewLayoutType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`, string(data))
	})
}

func TestLayoutType_UnmarshalJSON(t *testing.T) {
	type User struct {
		Date      carbon.LayoutType[carbon.Date]      `json:"date"`
		DateMilli carbon.LayoutType[carbon.DateMilli] `json:"date_milli"`
		DateMicro carbon.LayoutType[carbon.DateMicro] `json:"date_micro"`
		DateNano  carbon.LayoutType[carbon.DateNano]  `json:"date_nano"`

		Time      carbon.LayoutType[carbon.Time]      `json:"time"`
		TimeMilli carbon.LayoutType[carbon.TimeMilli] `json:"time_milli"`
		TimeMicro carbon.LayoutType[carbon.TimeMicro] `json:"time_micro"`
		TimeNano  carbon.LayoutType[carbon.TimeNano]  `json:"time_nano"`

		DateTime      carbon.LayoutType[carbon.DateTime]      `json:"date_time"`
		DateTimeMilli carbon.LayoutType[carbon.DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro carbon.LayoutType[carbon.DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  carbon.LayoutType[carbon.DateTimeNano]  `json:"date_time_nano"`
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
	c := carbon.Now()
	dataType := "time"

	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.Date](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateNano](c).GormDataType())

	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.Time](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.TimeMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.TimeMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.TimeNano](c).GormDataType())

	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateTime](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateTimeMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateTimeMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewLayoutType[carbon.DateTimeNano](c).GormDataType())
}

func TestFormatType_Scan(t *testing.T) {
	c := carbon.NewFormatType[carbon.DateTime](carbon.Now())

	t.Run("[]byte type", func(t *testing.T) {
		assert.Nil(t, c.Scan([]byte(carbon.Now().ToDateString())))
	})

	t.Run("string type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().ToDateString()))
	})

	t.Run("int64 type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().Timestamp()))
	})

	t.Run("time type", func(t *testing.T) {
		assert.Nil(t, c.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		assert.Error(t, c.Scan(nil))
	})
}

func TestFormatType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v, err := carbon.NewFormatType[carbon.DateTime](nil).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Nil(t, err)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Nil(t, v)
		assert.Error(t, err)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")
		v, err := carbon.NewFormatType[carbon.DateTime](c).Value()
		assert.Equal(t, c.StdTime(), v)
		assert.Nil(t, err)
	})
}

func TestFormatType_MarshalJSON(t *testing.T) {
	type User struct {
		Date      carbon.FormatType[carbon.Date]      `json:"date"`
		DateMilli carbon.FormatType[carbon.DateMilli] `json:"date_milli"`
		DateMicro carbon.FormatType[carbon.DateMicro] `json:"date_micro"`
		DateNano  carbon.FormatType[carbon.DateNano]  `json:"date_nano"`

		Time      carbon.FormatType[carbon.Time]      `json:"time"`
		TimeMilli carbon.FormatType[carbon.TimeMilli] `json:"time_milli"`
		TimeMicro carbon.FormatType[carbon.TimeMicro] `json:"time_micro"`
		TimeNano  carbon.FormatType[carbon.TimeNano]  `json:"time_nano"`

		DateTime      carbon.FormatType[carbon.DateTime]      `json:"date_time"`
		DateTimeMilli carbon.FormatType[carbon.DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro carbon.FormatType[carbon.DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  carbon.FormatType[carbon.DateTimeNano]  `json:"date_time_nano"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Date = carbon.NewFormatType[carbon.Date](nil)
		user.DateMilli = carbon.NewFormatType[carbon.DateMilli](nil)
		user.DateMicro = carbon.NewFormatType[carbon.DateMicro](nil)
		user.DateNano = carbon.NewFormatType[carbon.DateNano](nil)

		user.Time = carbon.NewFormatType[carbon.Time](nil)
		user.TimeMilli = carbon.NewFormatType[carbon.TimeMilli](nil)
		user.TimeMicro = carbon.NewFormatType[carbon.TimeMicro](nil)
		user.TimeNano = carbon.NewFormatType[carbon.TimeNano](nil)

		user.DateTime = carbon.NewFormatType[carbon.DateTime](nil)
		user.DateTimeMilli = carbon.NewFormatType[carbon.DateTimeMilli](nil)
		user.DateTimeMicro = carbon.NewFormatType[carbon.DateTimeMicro](nil)
		user.DateTimeNano = carbon.NewFormatType[carbon.DateTimeNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()

		user.Date = carbon.NewFormatType[carbon.Date](c)
		user.DateMilli = carbon.NewFormatType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewFormatType[carbon.DateMicro](c)
		user.DateNano = carbon.NewFormatType[carbon.DateNano](c)

		user.Time = carbon.NewFormatType[carbon.Time](c)
		user.TimeMilli = carbon.NewFormatType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewFormatType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewFormatType[carbon.TimeNano](c)

		user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewFormatType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewFormatType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewFormatType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"","date_milli":"","date_micro":"","date_nano":"","time":"","time_milli":"","time_micro":"","time_nano":"","date_time":"","date_time_milli":"","date_time_micro":"","date_time_nano":""}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Date = carbon.NewFormatType[carbon.Date](c)
		user.DateMilli = carbon.NewFormatType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewFormatType[carbon.DateMicro](c)
		user.DateNano = carbon.NewFormatType[carbon.DateNano](c)

		user.Time = carbon.NewFormatType[carbon.Time](c)
		user.TimeMilli = carbon.NewFormatType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewFormatType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewFormatType[carbon.TimeNano](c)

		user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewFormatType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewFormatType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewFormatType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("2020-08-05 13:14:15.999999999")

		user.Date = carbon.NewFormatType[carbon.Date](c)
		user.DateMilli = carbon.NewFormatType[carbon.DateMilli](c)
		user.DateMicro = carbon.NewFormatType[carbon.DateMicro](c)
		user.DateNano = carbon.NewFormatType[carbon.DateNano](c)

		user.Time = carbon.NewFormatType[carbon.Time](c)
		user.TimeMilli = carbon.NewFormatType[carbon.TimeMilli](c)
		user.TimeMicro = carbon.NewFormatType[carbon.TimeMicro](c)
		user.TimeNano = carbon.NewFormatType[carbon.TimeNano](c)

		user.DateTime = carbon.NewFormatType[carbon.DateTime](c)
		user.DateTimeMilli = carbon.NewFormatType[carbon.DateTimeMilli](c)
		user.DateTimeMicro = carbon.NewFormatType[carbon.DateTimeMicro](c)
		user.DateTimeNano = carbon.NewFormatType[carbon.DateTimeNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"date":"2020-08-05","date_milli":"2020-08-05.999","date_micro":"2020-08-05.999999","date_nano":"2020-08-05.999999999","time":"13:14:15","time_milli":"13:14:15.999","time_micro":"13:14:15.999999","time_nano":"13:14:15.999999999","date_time":"2020-08-05 13:14:15","date_time_milli":"2020-08-05 13:14:15.999","date_time_micro":"2020-08-05 13:14:15.999999","date_time_nano":"2020-08-05 13:14:15.999999999"}`, string(data))
	})
}

func TestFormatType_UnmarshalJSON(t *testing.T) {
	type User struct {
		Date      carbon.FormatType[carbon.Date]      `json:"date"`
		DateMilli carbon.FormatType[carbon.DateMilli] `json:"date_milli"`
		DateMicro carbon.FormatType[carbon.DateMicro] `json:"date_micro"`
		DateNano  carbon.FormatType[carbon.DateNano]  `json:"date_nano"`

		Time      carbon.FormatType[carbon.Time]      `json:"time"`
		TimeMilli carbon.FormatType[carbon.TimeMilli] `json:"time_milli"`
		TimeMicro carbon.FormatType[carbon.TimeMicro] `json:"time_micro"`
		TimeNano  carbon.FormatType[carbon.TimeNano]  `json:"time_nano"`

		DateTime      carbon.FormatType[carbon.DateTime]      `json:"date_time"`
		DateTimeMilli carbon.FormatType[carbon.DateTimeMilli] `json:"date_time_milli"`
		DateTimeMicro carbon.FormatType[carbon.DateTimeMicro] `json:"date_time_micro"`
		DateTimeNano  carbon.FormatType[carbon.DateTimeNano]  `json:"date_time_nano"`
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
	c := carbon.Now()
	dataType := "time"

	assert.Equal(t, dataType, carbon.NewFormatType[carbon.Date](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateNano](c).GormDataType())

	assert.Equal(t, dataType, carbon.NewFormatType[carbon.Time](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.TimeMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.TimeMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.TimeNano](c).GormDataType())

	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateTime](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateTimeMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateTimeMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewFormatType[carbon.DateTimeNano](c).GormDataType())
}

func TestTimestampType_Scan(t *testing.T) {
	t.Run("[]byte type", func(t *testing.T) {
		ts1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		assert.Error(t, ts1.Scan([]byte("xxx")))
		assert.Nil(t, ts1.Scan([]byte(strconv.Itoa(int(ts1.Timestamp())))))

		ts2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		assert.Error(t, ts2.Scan([]byte("xxx")))
		assert.Nil(t, ts2.Scan([]byte(strconv.Itoa(int(ts2.Timestamp())))))

		ts3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		assert.Error(t, ts3.Scan([]byte("xxx")))
		assert.Nil(t, ts3.Scan([]byte(strconv.Itoa(int(ts3.Timestamp())))))

		ts4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		assert.Error(t, ts4.Scan([]byte("xxx")))
		assert.Nil(t, ts4.Scan([]byte(strconv.Itoa(int(ts4.Timestamp())))))
	})

	t.Run("string type", func(t *testing.T) {
		ts1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		assert.Error(t, ts1.Scan("xxx"))
		assert.Nil(t, ts1.Scan(strconv.Itoa(int(ts1.Timestamp()))))

		ts2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		assert.Error(t, ts2.Scan("xxx"))
		assert.Nil(t, ts2.Scan(strconv.Itoa(int(ts2.Timestamp()))))

		ts3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		assert.Error(t, ts3.Scan("xxx"))
		assert.Nil(t, ts3.Scan(strconv.Itoa(int(ts3.Timestamp()))))

		ts4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		assert.Error(t, ts4.Scan("xxx"))
		assert.Nil(t, ts4.Scan(strconv.Itoa(int(ts4.Timestamp()))))
	})

	t.Run("int64 type", func(t *testing.T) {
		ts1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		assert.Nil(t, ts1.Scan(carbon.Now().Timestamp()))

		ts2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		assert.Nil(t, ts2.Scan(carbon.Now().TimestampMilli()))

		ts3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		assert.Nil(t, ts3.Scan(carbon.Now().TimestampMicro()))

		ts4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		assert.Nil(t, ts4.Scan(carbon.Now().TimestampNano()))
	})

	t.Run("time type", func(t *testing.T) {
		ts1 := carbon.NewTimestampType[carbon.Timestamp](carbon.Now())
		assert.Nil(t, ts1.Scan(carbon.Now().StdTime()))

		ts2 := carbon.NewTimestampType[carbon.TimestampMilli](carbon.Now())
		assert.Nil(t, ts2.Scan(carbon.Now().StdTime()))

		ts3 := carbon.NewTimestampType[carbon.TimestampMicro](carbon.Now())
		assert.Nil(t, ts3.Scan(carbon.Now().StdTime()))

		ts4 := carbon.NewTimestampType[carbon.TimestampNano](carbon.Now())
		assert.Nil(t, ts4.Scan(carbon.Now().StdTime()))
	})

	t.Run("unsupported type", func(t *testing.T) {
		c := carbon.Now()

		ts1 := carbon.NewTimestampType[carbon.Timestamp](c)
		assert.Error(t, ts1.Scan(nil))

		ts2 := carbon.NewTimestampType[carbon.TimestampMilli](c)
		assert.Error(t, ts2.Scan(nil))

		ts3 := carbon.NewTimestampType[carbon.TimestampMicro](c)
		assert.Error(t, ts3.Scan(nil))

		ts4 := carbon.NewTimestampType[carbon.TimestampNano](c)
		assert.Error(t, ts4.Scan(nil))
	})
}

func TestTimestampType_Value(t *testing.T) {
	t.Run("nil time", func(t *testing.T) {
		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](nil).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](nil).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](nil).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](nil).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("zero time", func(t *testing.T) {
		c := carbon.NewCarbon()

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Nil(t, e4)
	})

	t.Run("invalid time", func(t *testing.T) {
		c := carbon.Parse("xxx")

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Nil(t, v1)
		assert.Error(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Nil(t, v2)
		assert.Error(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Nil(t, v3)
		assert.Error(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Nil(t, v4)
		assert.Error(t, e4)
	})

	t.Run("valid time", func(t *testing.T) {
		c := carbon.Parse("2020-08-05")

		v1, e1 := carbon.NewTimestampType[carbon.Timestamp](c).Value()
		assert.Equal(t, c.Timestamp(), v1)
		assert.Nil(t, e1)

		v2, e2 := carbon.NewTimestampType[carbon.TimestampMilli](c).Value()
		assert.Equal(t, c.TimestampMilli(), v2)
		assert.Nil(t, e2)

		v3, e3 := carbon.NewTimestampType[carbon.TimestampMicro](c).Value()
		assert.Equal(t, c.TimestampMicro(), v3)
		assert.Nil(t, e3)

		v4, e4 := carbon.NewTimestampType[carbon.TimestampNano](c).Value()
		assert.Equal(t, c.TimestampNano(), v4)
		assert.Nil(t, e4)
	})
}

func TestTimestamp_MarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
	}

	t.Run("nil time", func(t *testing.T) {
		var user User

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](nil)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](nil)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](nil)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](nil)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("zero time", func(t *testing.T) {
		var user User

		c := carbon.NewCarbon()

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":0,"timestamp_milli":0,"timestamp_micro":0,"timestamp_nano":0}`, string(data))
	})

	t.Run("invalid time", func(t *testing.T) {
		var user User

		c := carbon.Parse("xxx")

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](c)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](c)
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](c)
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](c)

		data, err := json.Marshal(&user)
		assert.Error(t, err)
		assert.Empty(t, string(data))
	})

	t.Run("valid time", func(t *testing.T) {
		var user User

		now := carbon.Parse("2020-08-05 13:14:15")

		user.Timestamp = carbon.NewTimestampType[carbon.Timestamp](now)
		user.TimestampMilli = carbon.NewTimestampType[carbon.TimestampMilli](now.AddDay())
		user.TimestampMicro = carbon.NewTimestampType[carbon.TimestampMicro](now.SubDay())
		user.TimestampNano = carbon.NewTimestampType[carbon.TimestampNano](now.EndOfDay())

		data, err := json.Marshal(&user)
		assert.NoError(t, err)
		assert.Equal(t, `{"timestamp":1596633255,"timestamp_milli":1596633255000,"timestamp_micro":1596633255000000,"timestamp_nano":1596671999999999999}`, string(data))
	})
}

func TestTimestamp_UnmarshalJSON(t *testing.T) {
	type User struct {
		Timestamp      carbon.TimestampType[carbon.Timestamp]      `json:"timestamp"`
		TimestampMilli carbon.TimestampType[carbon.TimestampMilli] `json:"timestamp_milli"`
		TimestampMicro carbon.TimestampType[carbon.TimestampMicro] `json:"timestamp_micro"`
		TimestampNano  carbon.TimestampType[carbon.TimestampNano]  `json:"timestamp_nano"`
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
	c := carbon.Now()
	dataType := "time"

	assert.Equal(t, dataType, carbon.NewTimestampType[carbon.Timestamp](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewTimestampType[carbon.TimestampMilli](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewTimestampType[carbon.TimestampMicro](c).GormDataType())
	assert.Equal(t, dataType, carbon.NewTimestampType[carbon.TimestampNano](c).GormDataType())
}

type RFC3339Layout string

func (t RFC3339Layout) SetLayout() string {
	return carbon.RFC3339Layout
}

func TestLayoutType_Customer(t *testing.T) {
	type User struct {
		Customer carbon.LayoutType[RFC3339Layout] `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")
	user.Customer = carbon.NewLayoutType[RFC3339Layout](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15Z"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15Z", person.Customer.String())
}

type ISO8601Format string

func (t ISO8601Format) SetFormat() string {
	return carbon.ISO8601Format
}

func TestFormatType_Customer(t *testing.T) {
	type User struct {
		Customer carbon.FormatType[ISO8601Format] `json:"customer"`
	}

	var user User

	c := carbon.Parse("2020-08-05 13:14:15")
	user.Customer = carbon.NewFormatType[ISO8601Format](c)

	data, err := json.Marshal(&user)
	assert.NoError(t, err)
	assert.Equal(t, `{"customer":"2020-08-05T13:14:15+00:00"}`, string(data))

	var person User
	err = json.Unmarshal(data, &person)

	assert.NoError(t, err)
	assert.Equal(t, "2020-08-05T13:14:15+00:00", person.Customer.String())
}
