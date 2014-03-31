package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"testing"
)

func TestAnchorPoint(t *testing.T) {
	tests := []struct {
		test     AnchorPoint
		expected string
	}{
		{
			AnchorPoint{grokeddit.GlobalId{547846, grokeddit.Comment}, Previous}, 
			"before=t1_bqpy",
		},
		{
			AnchorPoint{grokeddit.GlobalId{547846, grokeddit.Comment}, Next}, 
			"after=t1_bqpy",
		},
		{
			AnchorPoint{grokeddit.GlobalId{54784642, grokeddit.Link}, Previous}, 
			"before=t3_wm83m",
		},
		{
			AnchorPoint{grokeddit.GlobalId{54784642, grokeddit.Link}, Next}, 
			"after=t3_wm83m",
		},
		{
			AnchorPoint{grokeddit.GlobalId{345353456, grokeddit.Subreddit}, Previous}, 
			"before=t5_5pm4fk",
		},
		{
			AnchorPoint{grokeddit.GlobalId{345353456, grokeddit.Subreddit}, Next}, 
			"after=t5_5pm4fk",
		},
	}

	for _, test := range tests {
		actual := test.test.String()
		if actual != test.expected {
			t.Errorf("Expected \"%s\" but got \"%s\"", test.expected, actual)
		}
	}
}
