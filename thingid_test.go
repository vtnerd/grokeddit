package grokeddit

import "testing"

func TestInvalidThingIdParsing(t *testing.T) {
	_, error := ParseId("")
	if error == nil {
		t.Error("Empty string is not a valid id")
	}

	_, error = ParseId("blah ")
	if error == nil {
		t.Error("blah is not a valid id")
	}
}

func TestValidThingIdParsing(t *testing.T) {
	testId, error := ParseId("20ko77")
	if error != nil {
		t.Error("unexpected error when parsing id")
	}

	if testId != 121896835 {
		t.Error("Improper base36 decoding")
	}

	testId, error = ParseId("0")
	if error != nil {
		t.Error("unexpected error when parsing id")
	}

	if testId != 0 {
		t.Error("Improper base36 decoding")
	}

	testId, error = ParseId("3W5E11264SGSF")
	if error != nil {
		t.Error("unexpected error when parsing id")
	}

	if testId != ThingId(^uint64(0)) {
		t.Error("Improper base36 decoding")
	}
}

func TestThingIdToString(t *testing.T) {
	testString := ThingId(121896835).String()
	if testString != "20ko77" {
		t.Error("Expected 20ko77 but got " + testString)
	}

	testString = ThingId(0).String()
	if testString != "0" {
		t.Error("Expected 0 but got " + testString)
	}

	testString = ThingId(^uint64(0)).String()
	if testString != "3w5e11264sgsf" {
		t.Error("Expected 3w5E11264sgsf but got " + testString)
	}
}
