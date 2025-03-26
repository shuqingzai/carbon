package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, Parse(""))
		assert.Empty(t, Parse("").ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, Parse("2020-08-05", "").HasError())
		assert.True(t, Parse("2020-08-05", "xxx").HasError())

		assert.Empty(t, Parse("2020-08-05", "").ToString())
		assert.Empty(t, Parse("2020-08-05", "xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, Parse("0").Error)
		assert.Error(t, Parse("xxx").Error)
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		assert.Equal(t, Now().Timestamp(), Parse("now").Timestamp())
		assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday").Timestamp())
		assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow").Timestamp())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-5").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-8-05").ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", Parse("2020-08-05").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-8-5 1:2:3").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:2:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 1:02:03").ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0000 UTC", Parse("2020-08-05 01:02:03").ToString())
	})

	t.Run("valid time with timezone", func(t *testing.T) {
		assert.Equal(t, Now().Timestamp(), Parse("now", PRC).Timestamp())
		assert.Equal(t, Yesterday().Timestamp(), Parse("yesterday", PRC).Timestamp())
		assert.Equal(t, Tomorrow().Timestamp(), Parse("tomorrow", PRC).Timestamp())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-5", PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-8-05", PRC).ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", Parse("2020-08-05", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-8-5 1:2:3", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:2:03", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 1:02:03", PRC).ToString())
		assert.Equal(t, "2020-08-05 01:02:03 +0800 CST", Parse("2020-08-05 01:02:03", PRC).ToString())
	})

	// https://github.com/dromara/carbon/issues/202
	t.Run("issue202", func(t *testing.T) {
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-1-8T09:02:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-08T9:2:48").ToString())
		assert.Equal(t, "2023-01-08 09:02:48 +0000 UTC", Parse("2023-01-8T9:2:48").ToString())
	})

	// https://github.com/dromara/carbon/issues/232
	t.Run("issue232", func(t *testing.T) {
		assert.Equal(t, "0000-01-01 00:00:00 +0000 UTC", Parse("0000-01-01 00:00:00").ToString())
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", Parse("0001-01-01 00:00:00").ToString())
		assert.Equal(t, "", Parse("0001-00-00 00:00:00").ToString())
	})
}

func TestParseByFormat(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseByFormat("", DateFormat))
		assert.Empty(t, ParseByFormat("", DateFormat).ToString())
	})

	t.Run("empty format", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", "").HasError())
		assert.Empty(t, ParseByFormat("2020-08-05", "").ToString())

	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", DateFormat, "").HasError())
		assert.True(t, ParseByFormat("2020-08-05", DateFormat, "xxx").HasError())

		assert.Empty(t, ParseByFormat("2020-08-05", DateFormat, "").ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", DateFormat, "xxx").ToString())
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		assert.True(t, ParseByFormat("2020-08-05", "U").HasError())
		assert.True(t, ParseByFormat("2020-08-05", "V").HasError())
		assert.True(t, ParseByFormat("2020-08-05", "X").HasError())
		assert.True(t, ParseByFormat("2020-08-05", "Z").HasError())

		assert.Empty(t, ParseByFormat("2020-08-05", "U").ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", "V").ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", "X").ToString())
		assert.Empty(t, ParseByFormat("2020-08-05", "Z").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseByFormat("0", DateFormat).Error)
		assert.Error(t, ParseByFormat("xxx", DateFormat).Error)
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByFormat("2020-08-05", DateFormat).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒").ToString())
	})

	t.Run("valid time with timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByFormat("2020-08-05", DateFormat, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020-08-05 13:14:15", DateTimeFormat, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("2020|08|05 13:14:15", "Y|m|d H:i:s", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("It is 2020-08-05 13:14:15", "\\I\\t \\i\\s Y-m-d H:i:s", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByFormat("今天是 2020年08月05日13时14分15秒", "今天是 Y年m月d日H时i分s秒", PRC).ToString())
	})

	// https://github.com/dromara/carbon/issues/206
	t.Run("issue206", func(t *testing.T) {
		assert.Equal(t, "2023-11-11 04:34:00 +0000 UTC", ParseByFormat("1699677240", "U").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666 +0000 UTC", ParseByFormat("1699677240666", "V").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666 +0000 UTC", ParseByFormat("1699677240666666", "X").ToString())
		assert.Equal(t, "2023-11-11 04:34:00.666666666 +0000 UTC", ParseByFormat("1699677240666666666", "Z").ToString())
	})
}

func TestParseByLayout(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseByLayout("", DateLayout))
		assert.Empty(t, ParseByLayout("", DateLayout).ToString())
	})

	t.Run("empty layout", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", "").HasError())
		assert.Empty(t, ParseByLayout("2020-08-05", "").ToString())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", DateLayout, "").HasError())
		assert.True(t, ParseByLayout("2020-08-05", DateLayout, "xxx").HasError())

		assert.Empty(t, ParseByLayout("2020-08-05", DateLayout, "").ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", DateLayout, "xxx").ToString())
	})

	t.Run("invalid timestamp", func(t *testing.T) {
		assert.True(t, ParseByLayout("2020-08-05", "timestamp").HasError())
		assert.True(t, ParseByLayout("2020-08-05", "timestampMilli").HasError())
		assert.True(t, ParseByLayout("2020-08-05", "timestampMicro").HasError())
		assert.True(t, ParseByLayout("2020-08-05", "timestampNano").HasError())

		assert.Empty(t, ParseByLayout("2020-08-05", "timestamp").ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", "timestampMilli").ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", "timestampMicro").ToString())
		assert.Empty(t, ParseByLayout("2020-08-05", "timestampNano").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseByLayout("0", DateLayout).Error)
		assert.Error(t, ParseByLayout("xxx", DateLayout).Error)
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", ParseByLayout("2020-08-05", DateLayout).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05").ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒").ToString())
	})

	t.Run("valid time with timezone", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 00:00:00 +0800 CST", ParseByLayout("2020-08-05", DateLayout, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020-08-05 13:14:15", DateTimeLayout, PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("2020|08|05 13:14:15", "2006|01|02 15:04:05", PRC).ToString())
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", ParseByLayout("今天是 2020年08月05日13时14分15秒", "今天是 2006年01月02日15时04分05秒", PRC).ToString())
	})
}

func TestParseWithLayouts(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseWithLayouts("", []string{DateTimeLayout}))
		assert.Empty(t, ParseWithLayouts("", []string{DateTimeLayout}).ToString())
	})

	t.Run("empty layouts", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseWithLayouts("2020-08-05 13:14:15", []string{}).ToString())
		assert.Equal(t, "2006-01-02 15:04:05", ParseWithLayouts("2020-08-05 13:14:15", []string{}).CurrentLayout())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "").HasError())
		assert.True(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "xxx").HasError())

		assert.Empty(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "").ToString())
		assert.Empty(t, ParseWithLayouts("2020-08-05 13:14:15", []string{DateLayout}, "xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseWithLayouts("0", []string{DateTimeLayout}).Error)
		assert.Error(t, ParseWithLayouts("xxx", []string{DateTimeLayout}).Error)
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		c := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"})
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})

	t.Run("valid time with timezone", func(t *testing.T) {
		c := ParseWithLayouts("2020|08|05 13|14|15", []string{"2006|01|02 15|04|05", "2006|1|2 3|4|5"}, PRC)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})
}

func TestParseWithFormats(t *testing.T) {
	t.Run("empty time", func(t *testing.T) {
		assert.Nil(t, ParseWithFormats("", []string{DateTimeFormat}))
		assert.Empty(t, ParseWithFormats("", []string{DateTimeFormat}).ToString())
	})

	t.Run("empty layouts", func(t *testing.T) {
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", ParseWithFormats("2020-08-05 13:14:15", []string{}).ToString())
		assert.Equal(t, "2006-01-02 15:04:05", ParseWithFormats("2020-08-05 13:14:15", []string{}).CurrentLayout())
	})

	t.Run("invalid timezone", func(t *testing.T) {
		assert.True(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "").HasError())
		assert.True(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "xxx").HasError())

		assert.Empty(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "").ToString())
		assert.Empty(t, ParseWithFormats("2020-08-05 13:14:15", []string{DateFormat}, "xxx").ToString())
	})

	t.Run("invalid time", func(t *testing.T) {
		assert.Error(t, ParseWithFormats("0", []string{DateFormat}).Error)
		assert.Error(t, ParseWithFormats("xxx", []string{DateFormat}).Error)
	})

	t.Run("valid time without timezone", func(t *testing.T) {
		c := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"})
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})

	t.Run("valid time with timezone", func(t *testing.T) {
		c := ParseWithFormats("2020|08|05 13|14|15", []string{"Y|m|d H|i|s", "y|m|d h|i|s"}, PRC)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c.ToString())
		assert.Equal(t, "2006|01|02 15|04|05", c.CurrentLayout())
	})
}
