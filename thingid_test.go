package grokeddit

import (
	"strings"
	"testing"
)

func TestInvalidThingIdParsing(t *testing.T) {
	invalidInputs := []string {
		"",
		"blah ",
	}
	
	for _, invalidInput := range invalidInputs {
		_, error := ParseId(invalidInput)

		if error == nil {
			t.Errorf("Expected error for input \"%s\"", invalidInput)
		}
	}
}

func TestValidThingIdParsing(t *testing.T) {

	validInputs := []struct {
		externalValue string
		internalValue ThingId
	}{
		{"20ko77", ThingId(121896835)},
		{"0", ThingId(0)},
		{"3W5E11264SGSF", ThingId(^uint64(0))},
	}

	for _, validInput := range validInputs {
		internalResult, error := ParseId(validInput.externalValue)

		if error != nil {
			t.Fatalf("Unexpected error for input \"%s\"", validInput.externalValue)
		}

		if internalResult != validInput.internalValue {
			t.Errorf("Expected value %d but got %d", validInput.internalValue, internalResult)
		}

		externalResult := validInput.internalValue.String()
		if externalResult != strings.ToLower(validInput.externalValue) {
			t.Errorf("Expected value \"%s\" but got \"%s\"", 
				strings.ToLower(validInput.externalValue), 
				externalResult)
		}
	}
}

