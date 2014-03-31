package fetcheddit

import "testing"

func TestAnchorDirection(t *testing.T) {
	testData := []struct {
		input    AnchorDirection
		expected string
	}{
		{Previous, "before"},
		{Next, "after"},
	}

	for _, test := range testData {
		actual := test.input.String()
		if actual != test.expected {
			t.Errorf("Expected \"%s\" but got \"%s\"", test.expected, actual)
		}
	}
}
