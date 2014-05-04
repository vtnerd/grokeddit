package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
	"net/url"
	"strings"
)

const (
	redditJsonSuffix     = ".json"
	redditBeforeModifier = "before"
	redditAfterModifier  = "after"
)

// Represents a resource on Reddit (page of links, comments, etc.)
type Path string

/* Create a new path object. The redditPath should be a relative path to the
desired resource. The path will automatically be "cleaned" up,
references to "before" and "after" in the query portion are removed. A ? or &
will automatically be added to the end so that a "before" or "after" can be
added to the Query portion without constructing another url.URL object. */
func CreatePath(redditPath string) (Path, error) {

	redditUrl, error := url.Parse(redditPath)
	if error != nil {
		return "", errors.New("Bad reddit path: " + error.Error())
	}

	redditUrl.Scheme = ""
	redditUrl.Host = ""

	if !strings.HasSuffix(redditUrl.Path, redditJsonSuffix) {
		redditUrl.Path = redditUrl.Path + redditJsonSuffix
	}

	// before/after is added when path is fetched
	if len(redditUrl.RawQuery) != 0 {
		queries := redditUrl.Query()
		queries.Del(redditBeforeModifier)
		queries.Del(redditAfterModifier)
		redditUrl.RawQuery = queries.Encode()
	}

	// flatten once to (hopefully) reduce garbage
	flattenedPath := redditUrl.String()

	// this is the sketchiest part -- assume dangling ? and & are
	// acceptable, and assume they won't be present in these situations ...
	// unittests can at least catch unexpected string values.
	if len(redditUrl.RawQuery) == 0 { // len can change after removal above
		flattenedPath = flattenedPath + "?"
	} else {
		flattenedPath = flattenedPath + "&"
	}

	return Path(flattenedPath), nil
}

/* Retrieves the path, at the specified anchor point (page), using the provided
 Fetcher.  */
func (path Path) FetchGrokedListing(contentFetcher Fetcher, anchor *AnchorPoint) (grokeddit.Groked, error) {
	retrievePath := string(path)

	if anchor != nil {
		retrievePath = retrievePath + anchor.String()
	}

	fetched, error := contentFetcher.Fetch(retrievePath)

	if error != nil {
		return grokeddit.Groked{}, errors.New("Unable to retrieve listing [" + retrievePath + "]: " + error.Error())
	}

	defer fetched.Close()
	return grokeddit.GrokListing(fetched)
}
