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

	type Expected struct {
		output    []*grokeddit.Thing
		requested []string
	}

	type TestData struct {
		input    Input
		expected Expected
	}

	// not sure how much how like this ... but I want to keep the two above
	// structs hidden to this file, not sure of any other way in Go
	validateThingList := func(
		t *testing.T, expected []*grokeddit.Thing, actual thingList) {

		for actual.hasNext() {

			result, error := actual.getNext()

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

					if !reflect.DeepEqual(result, *(expected[0])) {
						t.Error("Mis-match on expected data")
					}
				}

				expected = expected[1:len(expected)]
			}
		}

		if len(expected) != 0 {
			t.Errorf("Expected %d more", len(expected))
		}

		_, error := actual.getNext()
		if error == nil {
			t.Error("Expected non-nil error, reached EOF")
		}
	}

	validateFetched := func(t *testing.T, expected []string, actual *TestFetch) {
		for _, expectedPath := range expected {
			actualPath, error := actual.GetNextFetchLocation()
			if error != nil {
				t.Error("Expected another path")
			} else {
				if expectedPath != actualPath {
					t.Errorf(
						"Expected retrieval path \"%s\" but got \"%s\"",
						expectedPath,
						actualPath)
				}
			}
		}

		_, error := actual.Fetch("blah")
		if error == nil {
			t.Errorf("Did not fetch all paths")
		}
		
		leftoverPath, error := actual.GetNextFetchLocation()
		if error == nil {
			t.Errorf("Expected another path retrieval \"%s\"", leftoverPath)
		}
	}

	tests := []TestData{
		// Test errors on first retrieval
		{Input{[]string{}, nil}, Expected{[]*grokeddit.Thing{}, []string{}}},
		{
			Input{
				[]string{},
				&AnchorPoint{grokeddit.GlobalId{}, Previous},
			},
			Expected{[]*grokeddit.Thing{}, []string{}},
		},
		{
			Input{
				[]string{},
				&AnchorPoint{grokeddit.GlobalId{}, Next},
			},
			Expected{[]*grokeddit.Thing{}, []string{}},
		},
		// Test errors on second retrieval
		{
			Input{[]string{listingForward}, nil},
			Expected{
				[]*grokeddit.Thing{
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					nil,
				},
				[]string{"test_path.json?"},
			},
		},
		{
			Input{
				[]string{listingReverse},
				&AnchorPoint{grokeddit.GlobalId{}, Previous},
			},
			Expected{
				[]*grokeddit.Thing{
					&listingOutputReverse.Children[1],
					&listingOutputReverse.Children[0],
					nil,
				},
				[]string{"test_path.json?before=t1_0"},
			},
		},
		{
			Input{
				[]string{listingForward},
				&AnchorPoint{grokeddit.GlobalId{}, Next},
			},
			Expected{
				[]*grokeddit.Thing{
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					nil,
				},
				[]string{"test_path.json?after=t1_0"},
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
			Expected{
				[]*grokeddit.Thing{
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					&listingOutputReverse.Children[0],
					&listingOutputReverse.Children[1],
				},
				[]string{
					"test_path.json?",
					"test_path.json?after=t3_20d5ol",
					"test_path.json?after=t3_20d5ol",
				},
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
			Expected{
				[]*grokeddit.Thing{
					&listingOutputReverse.Children[1],
					&listingOutputReverse.Children[0],
					&listingOutputReverse.Children[1],
					&listingOutputReverse.Children[0],
					&listingOutputForward.Children[1],
					&listingOutputForward.Children[0],
				},
				[]string{
					"test_path.json?before=t1_0",
					"test_path.json?before=t3_20d5ol",
					"test_path.json?before=t3_20d5ol",
				},
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
			Expected{
				[]*grokeddit.Thing{
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					&listingOutputForward.Children[0],
					&listingOutputForward.Children[1],
					&listingOutputReverse.Children[0],
					&listingOutputReverse.Children[1],
				},
				[]string{
					"test_path.json?after=t1_0",
					"test_path.json?after=t3_20d5ol",
					"test_path.json?after=t3_20d5ol",
				},
			},
		},
	}

	for _, test := range tests {

		path, error := CreatePath("test_path")
		if error != nil {
			t.Error("Couldn't create path!")
		}

		testFetch := CreateTestFetch(test.input.list)
		thingList, error := fetchThingList(
			path.FetchGrokedListing,
			testFetch,
			test.input.initialAnchor)

		if error != nil {
			if len(test.expected.output) != 0 {
				t.Error("Unexpected error when creating thing list")
			}
		} else { // no error retrieving first block
			validateThingList(t, test.expected.output, thingList)
		}

		validateFetched(t, test.expected.requested, testFetch)
	}
}
