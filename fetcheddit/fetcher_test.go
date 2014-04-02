package fetcheddit

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

type TestFetch struct {
	LastFetchPath string // The last path provided in a Fetch call
	NextReturn    string // String that will be returned in the next Fetch call
	FetchError    bool   // True indicates the fetch call should error
}

func (testFetch *TestFetch) Fetch(path string) (io.ReadCloser, error) {
	testFetch.LastFetchPath = path

	if testFetch.FetchError {
		return nil, errors.New("Expected error return")
	}

	return ioutil.NopCloser(strings.NewReader(testFetch.NextReturn)), nil
}
