package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"reflect"
	"testing"
)

func TestThingList(t *testing.T) {
	type Input struct {
		list          []string
		initialAnchor *AnchorPoint
	}

	type TestData struct {
		input  Input
		output []*grokeddit.Thing
	}

	// not sure how much how like this ... but I want to keep the two above
	// structs hidden to this file, not sure of any other way in Go
	validateThingList := func(
		t *testing.T, expected []*grokeddit.Thing, input thingList) {

		for input.hasNext() {

			result, error := input.getNext()

			if len(expected) == 0 {
				t.Error("Expected less")
			} else { // have another expected element

				if expected[0] == nil {
					if error == nil {
						t.Error("Expected non-nil error")
					}
				} else { // expected element
					if error != nil {
						t.Error("Expected non-nil error, has next!")
					}

					if !reflect.DeepEqual(result, *expected[0]) {
						t.Error("Mis-match on expected data")
					}
				}

				expected = expected[1:len(expected)]
			}
		}

		if len(expected) != 0 {
			t.Errorf("Expected %d more", len(expected))
		}

		_, error := input.getNext()
		if error == nil {
			t.Error("Expected non-nil error, reached EOF")
		}
	}

	tests := []TestData{
		// Test errors on first retrieval 
		{Input{[]string{}, nil}, []*grokeddit.Thing{}},
		{
			Input{
				[]string{},
				&AnchorPoint{grokeddit.GlobalId{}, Previous},
			},
			[]*grokeddit.Thing{},
		},
		{
			Input{
				[]string{},
				&AnchorPoint{grokeddit.GlobalId{}, Next},
			},
			[]*grokeddit.Thing{},
		},
		// Test errors on second retrieval
		{
			Input{[]string{listingForward}, nil},
			[]*grokeddit.Thing{
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				nil,
			},
		},
		{
			Input{
				[]string{listingReverse},
				&AnchorPoint{grokeddit.GlobalId{}, Previous},
			},
			[]*grokeddit.Thing{
				&listingOutputReverse.Children[1],
				&listingOutputReverse.Children[0],
				nil,
			},
		},
		{
			Input{
				[]string{listingForward},
				&AnchorPoint{grokeddit.GlobalId{}, Next},
			},
			[]*grokeddit.Thing{
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				nil,
			},
		},
		// Test complete retrieval
		{
			Input{
				[]string{
					listingForward,
					listingForward,
					listingReverse,
				},
				nil,
			},
			[]*grokeddit.Thing{
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				&listingOutputReverse.Children[0],
				&listingOutputReverse.Children[1],
			},
		},
		{
			Input{
				[]string{
					listingReverse,
					listingReverse,
					listingForward,
				},
				&AnchorPoint{grokeddit.GlobalId{}, Previous},
			},
			[]*grokeddit.Thing{
				&listingOutputReverse.Children[1],
				&listingOutputReverse.Children[0],
				&listingOutputReverse.Children[1],
				&listingOutputReverse.Children[0],
				&listingOutputForward.Children[1],
				&listingOutputForward.Children[0],
			},
		},
		{
			Input{
				[]string{
					listingForward,
					listingForward,
					listingReverse,
				},
				&AnchorPoint{grokeddit.GlobalId{}, Next},
				},
			[]*grokeddit.Thing{
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				&listingOutputForward.Children[0],
				&listingOutputForward.Children[1],
				&listingOutputReverse.Children[0],
				&listingOutputReverse.Children[1],
			},
		},
	}

	for _, test := range tests {
		path, error := CreatePath("test_path")
		if error != nil {
			t.Error("Couldn't create path!")
		}

		thingList, error := fetchThingList(
			path.FetchGrokedListing,
			CreateTestFetch(test.input.list),
			test.input.initialAnchor)

		expectedList := test.output
		if error != nil {
			if len(expectedList) != 0 {
				t.Error("Unexpected error when creating thing list")
			}
		} else { // no error retrieving first block
			validateThingList(t, expectedList, thingList)
		}
	}
}
