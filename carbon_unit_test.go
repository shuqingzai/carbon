package carbon

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewCarbon(t *testing.T) {
	loc, _ := time.LoadLocation(PRC)

	t1, _ := time.Parse(DateTimeLayout, "2020-08-05 13:14:15")
	t2, _ := time.ParseInLocation(DateTimeLayout, "2020-08-05 13:14:15", loc)

	t.Run("zero time", func(t *testing.T) {
		c1 := NewCarbon()
		assert.Equal(t, "0001-01-01 00:00:00 +0000 UTC", c1.ToString())
		assert.Equal(t, time.Time{}.String(), c1.ToString())

		c2 := NewCarbon().SetLocation(loc)
		assert.Equal(t, "0001-01-01 08:05:43 +0805 LMT", c2.ToString())
		assert.Equal(t, time.Time{}.In(loc).String(), c2.ToString())
	})

	t.Run("valid time", func(t *testing.T) {
		c1 := NewCarbon(t1)
		assert.Equal(t, "2020-08-05 13:14:15 +0000 UTC", c1.ToString())
		assert.Equal(t, t1.String(), c1.ToString())

		c2 := NewCarbon(t2)
		assert.Equal(t, "2020-08-05 13:14:15 +0800 CST", c2.ToString())
		assert.Equal(t, t2.String(), c2.ToString())
	})
}

func TestCarbon_Copy(t *testing.T) {
	t.Run("copy time", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", oldCarbon.ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())

		oldCarbon = oldCarbon.AddDay()
		assert.Equal(t, "2020-08-06 00:00:00 +0000 UTC", oldCarbon.ToString())
		assert.Equal(t, "2020-08-05 00:00:00 +0000 UTC", newCarbon.ToString())
	})

	t.Run("copy timezone", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, UTC, oldCarbon.Timezone())
		assert.Equal(t, UTC, newCarbon.Timezone())

		oldCarbon = oldCarbon.SetTimezone(PRC)
		assert.Equal(t, PRC, oldCarbon.Timezone())
		assert.Equal(t, UTC, newCarbon.Timezone())

		newCarbon = newCarbon.SetTimezone(Japan)
		assert.Equal(t, PRC, oldCarbon.Timezone())
		assert.Equal(t, Japan, newCarbon.Timezone())
	})

	t.Run("copy layout", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, "2006-01-02", oldCarbon.CurrentLayout())
		assert.Equal(t, "2006-01-02", newCarbon.CurrentLayout())

		oldCarbon = oldCarbon.SetLayout(DateTimeLayout)
		assert.Equal(t, DateTimeLayout, oldCarbon.CurrentLayout())
		assert.Equal(t, DateLayout, newCarbon.CurrentLayout())

		newCarbon = newCarbon.SetLayout(RFC1036Layout)
		assert.Equal(t, DateTimeLayout, oldCarbon.CurrentLayout())
		assert.Equal(t, RFC1036Layout, newCarbon.CurrentLayout())
	})

	t.Run("copy weekStartsAt", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, DefaultWeekStartsAt, oldCarbon.WeekStartsAt())
		assert.Equal(t, DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		oldCarbon = oldCarbon.SetWeekStartsAt(Monday)
		assert.Equal(t, Monday, oldCarbon.WeekStartsAt())
		assert.Equal(t, DefaultWeekStartsAt, newCarbon.WeekStartsAt())

		newCarbon = newCarbon.SetWeekStartsAt(Sunday)
		assert.Equal(t, Monday, oldCarbon.WeekStartsAt())
		assert.Equal(t, Sunday, newCarbon.WeekStartsAt())
	})

	t.Run("copy lang", func(t *testing.T) {
		oldCarbon := Parse("2020-08-05")
		newCarbon := oldCarbon.Copy()

		assert.Equal(t, August, oldCarbon.ToMonthString())
		assert.Equal(t, August, newCarbon.ToMonthString())

		oldCarbon.SetLocale("zh-CN")

		assert.Equal(t, "八月", oldCarbon.ToMonthString())
		assert.Equal(t, "August", newCarbon.ToMonthString())

		newCarbon = newCarbon.SetLocale("jp")

		assert.Equal(t, "八月", oldCarbon.ToMonthString())
		assert.Equal(t, "8月", newCarbon.ToMonthString())
	})

	t.Run("copy error", func(t *testing.T) {
		oldCarbon := Parse("xxx")
		newCarbon := oldCarbon.Copy()

		assert.True(t, oldCarbon.HasError())
		assert.True(t, newCarbon.HasError())

		newCarbon = newCarbon.SetLayout("xxx")
		assert.True(t, oldCarbon.HasError())
		assert.True(t, newCarbon.HasError())
	})
}
