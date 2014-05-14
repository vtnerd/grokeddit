package fetcheddit

import (
	"errors"
	"io"
	"net/http"
)

const (
	DefaultFetch = HttpFetch("http://www.reddit.com")
)

/* A Fetcher that uses http. String represents the domain appended to
fetched paths. */
type HttpFetch string

/* Fetches the path using an HTTP get request. The string value is used as the
domain (prepended to the path). The ReadCloser returns is to the body. An error
is returned if the page could not be retrieved, or if a non 200 status code is
returned. */
func (httpFetch HttpFetch) Fetch(relativePath string) (io.ReadCloser, error) {

	getUrl := string(httpFetch) + relativePath
	response, error := http.Get(getUrl)

	if error != nil {
		return nil, errors.New("Unable to retreive listing [" + getUrl + "]: " + error.Error())
	}

	if response.StatusCode != 200 {
		return nil, errors.New("Unable to retrieve listing [" + getUrl + "]: Expected response code 200")
	}

	return response.Body, nil
}
