package fetcheddit

import (
	"errors"
	"net/url"
	"strings"
)

const (
	subredditSeparator = "+"
	pathPrefix = "http://www.reddit.com/r/"
	redditPostfix = "/"
	filterPostfix = ".json?limit=100&"
)

type Subreddit string

func validateSubredditName(subredditName string) error {
	if strings.Contains(subredditName, subredditSeparator) {
		return errors.New("Subreddit names cannot contain a '+' character")
	}

	return nil
}

func (subreddit Subreddit) getPath(filter string) (Path, error) {
	relativeUrl, error := url.Parse("/r/" + string(subreddit) + "/" + filter)
	if error != nil {
		return "", errors.New("Internal error creating subreddit path: " + error.Error())
	}

	return CreatePath(relativeUrl), nil
}

func (subreddit Subreddit) fetchLinks(filter string, anchor *AnchorPoint) (Links, error) {

	path, error := subreddit.getPath(filter)
	if error != nil {
		return Links{}, error
	}

	thingList, error := fetchThingList(path.FetchGrokedListing, anchor)
	if error != nil {
		return Links{}, error
	}
	
	// fetchLinks is in links.go
	return fetchLinks(thingList), nil
}

func FetchSubreddit(subredditName string) (Subreddit, error) {
	if error := validateSubredditName(subredditName); error != nil {
		return "", error
	}

	return Subreddit(subredditName), nil
}

func FetchSubreddits(subreddits []string) (Subreddit, error) {
	var combinedSubredditPath string

	for _, redditName := range subreddits {
		if error := validateSubredditName(redditName); error != nil {
			return "", error
		}

		// reddit appears to support a leading +, so lets 
		// cheat and do that
		combinedSubredditPath = combinedSubredditPath + subredditSeparator + redditName
	}

	return Subreddit(combinedSubredditPath), nil
}

func (subreddit Subreddit) FetchLinks(anchor *AnchorPoint) (Links, error) {
	return subreddit.fetchLinks("", anchor)
}

func (subreddit Subreddit) FetchNewLinks(anchor *AnchorPoint) (Links, error) {
	return subreddit.fetchLinks("new", anchor)
}

func (subreddit Subreddit) FetchTopLinks(anchor *AnchorPoint) (Links, error) {
	return subreddit.fetchLinks("top", anchor)
}

func (subreddit Subreddit) FetchControversialLinks(anchor *AnchorPoint) (Links, error) {
	return subreddit.fetchLinks("controversial", anchor)
}

