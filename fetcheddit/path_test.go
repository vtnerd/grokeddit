package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"reflect"
	"testing"
)

/*
 I wish I had some tests that caused CreatePath to error. I haven't found any
inputs to url.Parse that fail.
*/


func TestPath(t *testing.T) {

	type inputData struct {
		redditPath  string
		fetchedData string
		anchor      *AnchorPoint
		fetchError  bool
	}

	type expectedResults struct {
		fetchPath string
		result    grokeddit.Groked
	}

	noChildren := make([]grokeddit.Thing, 0, 0)

	tests := []struct {
		input    inputData
		expected expectedResults
	}{
		{inputData{"", "", nil, true}, expectedResults{"", grokeddit.Groked{}}},
		{
			inputData{"", "", nil, false},
			expectedResults{".json?", grokeddit.Groked{Children: noChildren}},
		},
		{
			inputData{
				"",
				"",
				&AnchorPoint{grokeddit.GlobalId{455, grokeddit.Comment}, Next},
				false,
			},
			expectedResults{".json?after=t1_cn", grokeddit.Groked{Children: noChildren}},
		},
		{
			inputData{
				"",
				"",
				&AnchorPoint{grokeddit.GlobalId{454, grokeddit.Link}, Previous},
				false,
			},
			expectedResults{".json?before=t3_cm", grokeddit.Groked{Children: noChildren}},
		},
		{
			inputData{"/r/all", "", nil, false},
			expectedResults{"/r/all.json?", grokeddit.Groked{Children: noChildren}},
		},
		{
			inputData{"/r/all?limit=100", "", nil, false},
			expectedResults{"/r/all.json?limit=100&", grokeddit.Groked{Children: noChildren}},
		},
		{
			inputData{"/r/all?limit=100&before=t3_r", "", nil, false},
			expectedResults{"/r/all.json?limit=100&", grokeddit.Groked{Children: noChildren}},
		},
	}

	for _, test := range tests {
		path, error := CreatePath(test.input.redditPath)

		if error != nil {
			t.Fatalf("Unexpected error with input path \"%s\"", test.input.redditPath)
		}

		testFetcher := &TestFetch{NextReturn: test.input.fetchedData, FetchError: test.input.fetchError}
		groked, error := path.FetchGrokedListing(testFetcher, test.input.anchor)

		if test.input.fetchError {
			if error == nil {
				t.Error("Expected error on Fetch call")
			}
		} else {
			if error != nil {
				t.Error("Unexpected error on Fetch call")
			}

			if testFetcher.LastFetchPath != test.expected.fetchPath {
				t.Errorf("Expected fetch path \"%s\" but got fetch path \"%s\"",
					test.expected.fetchPath,
					testFetcher.LastFetchPath)
			}

			if !reflect.DeepEqual(groked, test.expected.result) {
				t.Error("Unexpected struct data")
			}
		}
	}
}
