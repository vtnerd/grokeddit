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

type Path string

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

	// this is added at the last stage
	queries := redditUrl.Query()
	queries.Del(redditBeforeModifier)
	queries.Del(redditAfterModifier)
	redditUrl.RawQuery = queries.Encode()

	// flatten once to (hopefully) reduce garbage
	flattenedPath := redditUrl.String()

	// this is the sketchiest part -- assume dangling ? and & are
	// acceptable, and assume they won't be present in these situations ...
	// unittests can at least catch unexpected string values.
	if len(redditUrl.Query()) == 0 {
		flattenedPath = flattenedPath + "?"
	} else {
		flattenedPath = flattenedPath + "&"
	}

	return Path(flattenedPath), nil
}

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
