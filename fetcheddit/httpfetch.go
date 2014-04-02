package fetcheddit

import (
	"errors"
	"io"
	"net/http"
//	"net/url"
)

const (
	DefaultFetch = HttpFetch("http://www.reddit.com")
)

type HttpFetch string

func CreateHttpFetch(domain string) HttpFetch {
	return HttpFetch("http://" + domain)
}

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
