package grokeddit

import "testing"

func TestParseInvalidGobalId(t *testing.T) {

	invalidInputs := []string {
		"",
		"t2_blah", // t2 unsupported
		"t3blah", 
		"t3_blah ",
	}
	
	for _, invalidInput := range invalidInputs {
		_, error := ParseGlobalId(invalidInput)

		if error == nil {
			t.Errorf("Expected error when parsing \"%s\"", invalidInput)
		}
	}
}

func TestParseValidGlobalId(t *testing.T) {

	validInputs := []struct {
		input string
		expected GlobalId
	}{
		{"t1_blah", GlobalId{540809, Comment}},
		{"t3_blbh", GlobalId{540845, Link}},
		{"t5_blai", GlobalId{540810, Subreddit}},
	}

	for _, validInput := range validInputs {
		result, error := ParseGlobalId(validInput.input)

		if error != nil {
			t.Fatalf("Did not expect error with input \"%s\"", validInput)
		}

		if validInput.expected.Kind != result.Kind {
			t.Errorf("Expected type %d but got %d", validInput.expected.Kind, result.Kind)
		}

		if validInput.expected.Id != result.Id {
			t.Errorf("Expected id %d but got %d", validInput.expected.Id, result.Id)
		}
	}
}
