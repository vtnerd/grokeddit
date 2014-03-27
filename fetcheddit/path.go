package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

const (
	redditRequestScheme = "http"
	redditHost = "www.reddit.com"
	redditJsonSuffix = ".json"
	redditBeforeModifier = "before"
	redditAfterModifier = "after"
)

type Path string 

func CreatePath(redditPath *url.URL) Path {
	if redditPath == nil {
		redditPath = new(url.URL)
	}

	redditPath.Scheme = redditRequestScheme
	redditPath.Host = redditHost

	if !strings.HasSuffix(redditPath.Path, redditJsonSuffix) {
		redditPath.Path = redditPath.Path + redditJsonSuffix
	}

	// this is added at the last stage
	redditPath.Query().Del(redditBeforeModifier)
	redditPath.Query().Del(redditAfterModifier)

	// flatten once to (hopefully) reduce garbage
	flattenedPath := redditPath.String()

	// this is the sketchiest part -- assume dangling ? and & are 
	// acceptable, and assume they won't be present in these situations ...
	// unittests can at least catch unexpected string values.
	if len(redditPath.Query()) == 0 {
		flattenedPath = flattenedPath + "?"
	} else {
		flattenedPath = flattenedPath + "&"
	}

	return Path(flattenedPath)
}

func (path Path) FetchGrokedListing(anchor *AnchorPoint) (grokeddit.Groked, error) {
	retrieveLocation := string(path)

	if anchor != nil {
		retrieveLocation = retrieveLocation + anchor.String()
	}

	response, error := http.Get(retrieveLocation)
	if error != nil {
		return grokeddit.Groked{}, errors.New("Unable to retrieve listing [" + retrieveLocation + "]: " + error.Error())
	}

	if response.StatusCode != 200 {
		return grokeddit.Groked{}, errors.New("Unable to retrieve listing [" + retrieveLocation + "]: Expected response code 200")
	}

	defer response.Body.Close()
	return grokeddit.GrokListing(response.Body)
}
