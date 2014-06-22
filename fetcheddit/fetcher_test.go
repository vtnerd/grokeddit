package fetcheddit

import (
	"code.leeclagett.com/grokeddit"
	"errors"
	"io"
	"io/ioutil"
	"strings"
)

const (
	listingForward string = `
{
    "kind": "Listing",
    "data": {
        "modhash": "vaygtcstccf58f59494458317e0f140807769e7dcbe40daf12",
        "children": [
            {
                "kind": "t3",
                "data": {
                    "domain": "self.redditdev",
                    "banned_by": null,
                    "media_embed": {},
                    "subreddit": "redditdev",
                    "selftext_html": "<html>1</html>",
                    "selftext": "1",
                    "likes": null,
                    "secure_media": null,
                    "link_flair_text": null,
                    "id": "20flmp",
                    "gilded": 0,
                    "secure_media_embed": {},
                    "clicked": false,
                    "stickied": false,
                    "author": "reality_bugger",
                    "media": null,
                    "score": 7,
                    "approved_by": null,
                    "over_18": false,
                    "hidden": false,
                    "thumbnail": "",
                    "subreddit_id": "t5_2qizd",
                    "edited": false,
                    "link_flair_css_class": null,
                    "author_flair_css_class": null,
                    "downs": 4,
                    "saved": false,
                    "is_self": true,
                    "permalink": "/r/redditdev/comments/20flmp/just_a_short_thank_you/",
                    "name": "t3_20flmp",
                    "created": 1394860235,
                    "url": "http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/",
                    "author_flair_text": null,
                    "title": "Just a short thank you.",
                    "created_utc": 1394831435,
                    "ups": 11,
                    "num_comments": 2,
                    "visited": false,
                    "num_reports": null,
                    "distinguished": null
                }
            },
            {
                "kind": "t3",
                "data": {
                    "domain": "self.redditdev",
                    "banned_by": null,
                    "media_embed": {},
                    "subreddit": "redditdev",
                    "selftext_html": "<html>2</html>",
                    "likes": null,
                    "secure_media": null,
                    "link_flair_text": null,
                    "id": "20e6ir",
                    "gilded": 0,
                    "secure_media_embed": {},
                    "clicked": false,
                    "stickied": false,
                    "author": "Grimzentide",
                    "media": null,
                    "score": 1,
                    "approved_by": null,
                    "over_18": false,
                    "hidden": false,
                    "thumbnail": "",
                    "subreddit_id": "t5_2qizd",
                    "edited": false,
                    "link_flair_css_class": null,
                    "author_flair_css_class": null,
                    "downs": 1,
                    "saved": false,
                    "is_self": true,
                    "permalink": "/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
                    "name": "t3_20e6ir",
                    "created": 1394822733,
                    "url": "http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
                    "author_flair_text": null,
                    "title": "Is this the most efficient way to run this code?",
                    "created_utc": 1394793933,
                    "ups": 2,
                    "num_comments": 3,
                    "visited": false,
                    "num_reports": null,
                    "distinguished": null
                }
            }
        ],
        "after": "t3_20d5ol",
        "before": null
    }
}
`

	listingReverse string = `
{
    "kind": "Listing",
    "data": {
        "modhash": "vaygtcstccf58f59494458317e0f140807769e7dcbe40daf12",
        "children": [
            {
                "kind": "t3",
                "data": {
                    "domain": "self.redditdev",
                    "banned_by": null,
                    "media_embed": {},
                    "subreddit": "redditdev",
                    "selftext_html": "<html>2</html>",
                    "likes": null,
                    "secure_media": null,
                    "link_flair_text": null,
                    "id": "20e6ir",
                    "gilded": 0,
                    "secure_media_embed": {},
                    "clicked": false,
                    "stickied": false,
                    "author": "Grimzentide",
                    "media": null,
                    "score": 1,
                    "approved_by": null,
                    "over_18": false,
                    "hidden": false,
                    "thumbnail": "",
                    "subreddit_id": "t5_2qizd",
                    "edited": false,
                    "link_flair_css_class": null,
                    "author_flair_css_class": null,
                    "downs": 1,
                    "saved": false,
                    "is_self": true,
                    "permalink": "/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
                    "name": "t3_20e6ir",
                    "created": 1394822733,
                    "url": "http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
                    "author_flair_text": null,
                    "title": "Is this the most efficient way to run this code?",
                    "created_utc": 1394793933,
                    "ups": 2,
                    "num_comments": 3,
                    "visited": false,
                    "num_reports": null,
                    "distinguished": null
                }
            },
            {
                "kind": "t3",
                "data": {
                    "domain": "self.redditdev",
                    "banned_by": null,
                    "media_embed": {},
                    "subreddit": "redditdev",
                    "selftext_html": "<html>1</html>",
                    "selftext": "1",
                    "likes": null,
                    "secure_media": null,
                    "link_flair_text": null,
                    "id": "20flmp",
                    "gilded": 0,
                    "secure_media_embed": {},
                    "clicked": false,
                    "stickied": false,
                    "author": "reality_bugger",
                    "media": null,
                    "score": 7,
                    "approved_by": null,
                    "over_18": false,
                    "hidden": false,
                    "thumbnail": "",
                    "subreddit_id": "t5_2qizd",
                    "edited": false,
                    "link_flair_css_class": null,
                    "author_flair_css_class": null,
                    "downs": 4,
                    "saved": false,
                    "is_self": true,
                    "permalink": "/r/redditdev/comments/20flmp/just_a_short_thank_you/",
                    "name": "t3_20flmp",
                    "created": 1394860235,
                    "url": "http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/",
                    "author_flair_text": null,
                    "title": "Just a short thank you.",
                    "created_utc": 1394831435,
                    "ups": 11,
                    "num_comments": 2,
                    "visited": false,
                    "num_reports": null,
                    "distinguished": null
                }
            }
        ],
        "before": "t3_20d5ol",
        "after": null
    }
}
`
)

var (
	listingOutputForward = grokeddit.Groked{
		nil,
		&grokeddit.GlobalId{121546245, grokeddit.Link},
		[]grokeddit.Thing{
			grokeddit.Thing{
				"reality_bugger",
				1394831435,
				grokeddit.GlobalId{121660225, grokeddit.Link},
				1394831435,
				grokeddit.GlobalId{},
				grokeddit.GlobalId{},
				grokeddit.Groked{},
				"redditdev",
				grokeddit.GlobalId{4596889, grokeddit.Subreddit},
				"<html>1</html>",
				"Just a short thank you.",
				"http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/",
			},
			grokeddit.Thing{
				"Grimzentide",
				1394793933,
				grokeddit.GlobalId{121593987, grokeddit.Link}, 1394793933,
				grokeddit.GlobalId{},
				grokeddit.GlobalId{},
				grokeddit.Groked{},
				"redditdev",
				grokeddit.GlobalId{4596889, grokeddit.Subreddit},
				"<html>2</html>",
				"Is this the most efficient way to run this code?",
				"http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
			},
		},
	}

	listingOutputReverse = grokeddit.Groked{
		&grokeddit.GlobalId{121546245, grokeddit.Link},
		nil,
		[]grokeddit.Thing{
			grokeddit.Thing{
				"Grimzentide",
				1394793933,
				grokeddit.GlobalId{121593987, grokeddit.Link}, 1394793933,
				grokeddit.GlobalId{},
				grokeddit.GlobalId{},
				grokeddit.Groked{},
				"redditdev",
				grokeddit.GlobalId{4596889, grokeddit.Subreddit},
				"<html>2</html>",
				"Is this the most efficient way to run this code?",
				"http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
			},
			grokeddit.Thing{
				"reality_bugger",
				1394831435,
				grokeddit.GlobalId{121660225, grokeddit.Link},
				1394831435,
				grokeddit.GlobalId{},
				grokeddit.GlobalId{},
				grokeddit.Groked{},
				"redditdev",
				grokeddit.GlobalId{4596889, grokeddit.Subreddit},
				"<html>1</html>",
				"Just a short thank you.",
				"http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/",
			},
		},
	}
)

type TestFetch struct {
	nextValue     chan string
	lastFetchPath chan string
}



func CreateTestFetch(values []string) *TestFetch {
	newTestFetch := TestFetch{
		make(chan string, len(values)),
		make(chan string, len(values)),
	}

	defer close(newTestFetch.nextValue)
	for _, value := range values {
		newTestFetch.nextValue <- value
	}

	return &newTestFetch
}

func (testFetch *TestFetch) Fetch(path string) (io.ReadCloser, error) {
	nextValue, ok := <-testFetch.nextValue
	if !ok {
		close(testFetch.lastFetchPath)
		return nil, errors.New("No more values to fetch")
	}

	testFetch.lastFetchPath <- path
	return ioutil.NopCloser(strings.NewReader(nextValue)), nil
}

func (testFetch *TestFetch) GetNextFetchLocation() (string, error) {
	nextValue, ok := <-testFetch.lastFetchPath
	if !ok {
		return "", errors.New("No more fetch locations")
	}

	return nextValue, nil
}
