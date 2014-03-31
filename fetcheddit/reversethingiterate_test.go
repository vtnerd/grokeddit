package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"reflect"
	"testing"
)

func TestReverseIterate(t *testing.T) {
	testData := []struct {
		input []grokeddit.Thing
	}{
		{nil},
		{[]grokeddit.Thing{}},
		{
			[]grokeddit.Thing{
				grokeddit.Thing{
					"example author",
					45455,
					grokeddit.GlobalId{45355, grokeddit.Comment},
					4534666,
					grokeddit.GlobalId{75674, grokeddit.Link},
					grokeddit.GlobalId{566654, grokeddit.Subreddit},
					grokeddit.Groked{},
					"the subreddit",
					grokeddit.GlobalId{645654, grokeddit.Subreddit},
					"text",
					"title title",
					"url url",
				},
			},
		},
		{
			[]grokeddit.Thing{
				grokeddit.Thing{
					"example author",
					45455,
					grokeddit.GlobalId{45355, grokeddit.Comment},
					4534666,
					grokeddit.GlobalId{75674, grokeddit.Link},
					grokeddit.GlobalId{566654, grokeddit.Subreddit},
					grokeddit.Groked{},
					"the subreddit",
					grokeddit.GlobalId{645654, grokeddit.Subreddit},
					"text",
					"title title",
					"url url",
				},
				grokeddit.Thing{
					"example author2",
					454545,
					grokeddit.GlobalId{345355, grokeddit.Comment},
					4534666,
					grokeddit.GlobalId{755674, grokeddit.Link},
					grokeddit.GlobalId{5686654, grokeddit.Subreddit},
					grokeddit.Groked{},
					"the subreddit2",
					grokeddit.GlobalId{6458654, grokeddit.Subreddit},
					"text2",
					"title title2",
					"url url2",
				},
			},
		},
	}

	for _, test := range testData {
		testObject := reverseThingIterate{}

		if testObject.hasNext() {
			t.Error("iterater empty - has next should return false")
		}

		if _, error := testObject.getNext(); error == nil {
			t.Error("iterater empty - get next should return error")
		}

		count := len(test.input) - 1
		testObject.setArray(test.input)
		for testObject.hasNext() {

			actual, error := testObject.getNext()
			if error != nil {
				t.Error("Unexpected error - hasNext() returned true")
			}

			if !reflect.DeepEqual(test.input[count], actual) {
				t.Error("Unexpected thing value")
			}

			count = count - 1			
		}

		if _, error := testObject.getNext(); error == nil {
			t.Error("iterater empty - get next should return error")
		}
	}
}
