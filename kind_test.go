package grokeddit

import "testing"

func TestParsingInvalidKinds(t *testing.T) {
	_, error := ParseKind("t2")
	if error == nil {
		t.Error("t2 is not a valid kind (yet)")
	}

	_, error = ParseKind("")
	if error == nil {
		t.Error("The empty should is never a valid kind")
	}

	_, error = ParseKind("no such kind")
	if error == nil {
		t.Error("A random string is never a valid kind")
	}
}

func TestParsingValidKinds(t *testing.T) {
	kind, error := ParseKind("t1")
	if error != nil {
		t.Error("Unexpected error when parsing kind t1: " + error.Error())
	}

	if kind != Comment {
		t.Error("Expected to kind type to be comment")
	}

	kind, error = ParseKind("t3")
	if error != nil {
		t.Error("Unexpected error when parsing kind t3: " + error.Error())
	}

	if kind != Link {
		t.Error("Expected to kind type to be comment")
	}

	kind, error = ParseKind("t5")
	if error != nil {
		t.Error("Unexpected error when parsing kind t5: " + error.Error())
	}

	if kind != Subreddit {
		t.Error("Expected to kind type to be comment")
	}
}

func TestKindStringConversion(t *testing.T) {
	testString := Comment.String()
	if testString != "t1" {
		t.Error("Expected \"t1\" but got \"" + testString + "\"")
	}

	testString = Link.String()
	if testString != "t3" {
		t.Error("Expected \"t3\" but got \"" + testString + "\"")
	}

	testString = Subreddit.String()
	if testString != "t5" {
		t.Error("Expected \"t5\" but got \"" + testString + "\"")
	}
}
