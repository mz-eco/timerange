package timerange

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func Test_Weekdays(t *testing.T) {

	var (
		tr = RangeTo(
			Month.Date(2018, 10), 7*Day)
	)

	for index, block := range tr.Split(Day) {

		assert.Equal(
			t,
			weekday(block.B),
			index,
			fmt.Sprintf("time %s is not weekday %d", block.B, index),
		)
	}
}
