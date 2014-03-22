package grokeddit

import "testing"

func TestParseInvalidGobalId(t *testing.T) {

	_, error := ParseGlobalId("")
	if error == nil {
		t.Error("Expected error when parsing \"\" for global id")
	}

	_, error = ParseGlobalId("t2_blah") // t2 unsupported
	if error == nil {
		t.Error("Expected error when parsing \"t2_blah\" global id")
	}

	_, error = ParseGlobalId("t3blah")
	if error == nil {
		t.Error("Expected error when parsing \"t2blag\" global id")
	}

	_, error = ParseGlobalId("t3_blah ") // space unsupported
	if error == nil {
		t.Error("Expected error when parsing \"t3_blah \" global id")
	}
}

func TestParseValidGlobalId(t *testing.T) {

	testId, error := ParseGlobalId("t1_blah")
	if error != nil {
		t.Error("Unexpected error when parsing global id \"t1_blah\": " + error.Error())
	}

	if testId.Id != 540809 {
		t.Error("Expected 540809 for the id value")
	}

	if testId.Kind != Comment {
		t.Error("Expected kind to be comment")
	}

	testId, error = ParseGlobalId("t3_blah")
	if error != nil {
		t.Error("Unexpected error when parsing global id \"t3_blah\": " + error.Error())
	}

	if testId.Id != 540809 {
		t.Error("Expected 540809 for the id value")
	}

	if testId.Kind != Link {
		t.Error("Expected kind to be link")
	}

	testId, error = ParseGlobalId("t5_blah")
	if error != nil {
		t.Error("Unexpected error when parsing global id \"t5_blah\": " + error.Error())
	}

	if testId.Id != 540809 {
		t.Error("Expected 540809 for the id value")
	}

	if testId.Kind != Subreddit {
		t.Error("Expected kind to be subreddit")
	}
}
