package carbon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetDefault(t *testing.T) {
	defer ResetDefault()

	SetDefault(Default{
		Layout:       DateTimeLayout,
		Timezone:     PRC,
		Locale:       "zh-CN",
		WeekStartsAt: Monday,
		WeekendDays: []Weekday{
			Saturday, Sunday,
		},
	})

	assert.Equal(t, DateTimeLayout, DefaultLayout)
	assert.Equal(t, PRC, DefaultTimezone)
	assert.Equal(t, "zh-CN", DefaultLocale)
	assert.Equal(t, Monday, DefaultWeekStartsAt)
	assert.Equal(t, []Weekday{
		Saturday, Sunday,
	}, DefaultWeekendDays)
}
