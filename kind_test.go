package grokeddit

import "testing"

func TestParsingInvalidKinds(t *testing.T) {
	invalidInputs := []string {
		"t2",
		"",
		"no such kind",
	}

	for _, invalidInput := range invalidInputs {
		_, error := ParseKind(invalidInput)

		if error == nil {
			t.Error("Expected error with input \"%s\"", invalidInput)
		}
	}
}

func TestParsingValidKinds(t *testing.T) {
	
	validInputs := []struct {
		external string
		internal KindType
	}{
		{"t1", Comment},
		{"t3", Link},
		{"t5", Subreddit},
	}

	for _, validInput := range validInputs {
		internalResult, error := ParseKind(validInput.external)

		if error != nil {
			t.Errorf("Unexpected error with input \"%s\"", validInput.external)
		}

		if validInput.internal != internalResult {
			t.Errorf("Expected %d but got %d", validInput.internal, internalResult)
		}

		externalResult := validInput.internal.String()

		if externalResult != validInput.external {
			t.Errorf("Expected \"%s\" but got \"%s\"", validInput.external, externalResult)
		}
	}
}
