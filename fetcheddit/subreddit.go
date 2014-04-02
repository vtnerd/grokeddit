package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
	"strings"
)

const (
	subredditSeparator = "+"
	subredditIndicator = "/r/"
	allSubreddits = "all"
)

type Subreddit struct {
	path           string
	contentFetcher Fetcher
}

func validateSubredditName(subredditName string) error {
	if strings.Contains(subredditName, subredditSeparator) {
		return errors.New("Subreddit names cannot contain a '+' character")
	}

	return nil
}

func (subreddit Subreddit) getPath(filter string) (Path, error) {
	return CreatePath(subredditIndicator + subreddit.path + "/" + filter)
}

func (subreddit Subreddit) fetchThingList(filter string, anchor *AnchorPoint) (thingList, error) {
	path, error := subreddit.getPath(filter)
	if error != nil {
		return thingList{}, error
	}

	// fetchThingList is in thinglist.go
	return fetchThingList(path.FetchGrokedListing, subreddit.contentFetcher, anchor)
}

func (subreddit Subreddit) fetchLinks(filter string, anchor *AnchorPoint) (Links, error) {
	if anchor != nil {
		if anchor.Anchor.Kind != grokeddit.Link {
			return Links{}, errors.New("Anchor point must be a link")
		}
	}

	thingList, error := subreddit.fetchThingList(filter, anchor)
	if error != nil {
		return Links{}, error
	}

	// fetchLinks is in links.go
	return fetchLinks(thingList), nil
}

func FetchSubreddit(subredditName string, contentFetcher Fetcher) (Subreddit, error) {
	if contentFetcher == nil {
		return Subreddit{}, errors.New("Cannot provide nil Fetcher")
	}

	if len(subredditName) == 0 {
		return Subreddit{allSubreddits, contentFetcher}, nil
	}

	if error := validateSubredditName(subredditName); error != nil {
		return Subreddit{}, error
	}

	return Subreddit{subredditName, contentFetcher}, nil
}

func FetchSubreddits(subreddits []string, contentFetcher Fetcher) (Subreddit, error) {
	if contentFetcher == nil {
		return Subreddit{}, errors.New("Cannot provide nil Fetcher")
	}

	if subreddits == nil || len(subreddits) == 0 {
		return Subreddit{allSubreddits, contentFetcher}, nil
	}

	var combinedSubredditPath string

	for _, redditName := range subreddits {
		if error := validateSubredditName(redditName); error != nil {
			return Subreddit{}, error
		}

		// reddit appears to support a leading +, so lets
		// cheat and do that
		combinedSubredditPath = combinedSubredditPath + subredditSeparator + redditName
	}

	return Subreddit{combinedSubredditPath, contentFetcher}, nil
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

func (subreddit Subreddit) FetchComments(anchor *AnchorPoint) (Comments, error) {
	if anchor != nil {
		if anchor.Anchor.Kind != grokeddit.Comment {
			return Comments{}, errors.New("Anchor point must be a comment")
		}
	}

	thingList, error := subreddit.fetchThingList("comments", anchor)
	if error != nil {
		return Comments{}, error
	}

	// fetchComments is in comments.go
	return fetchComments(thingList), nil
}
