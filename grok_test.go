package grokeddit

import (
	"encoding/json"
	"strconv"
	"strings"
	"testing"
)

const (
	reddit3 string = `
{
    "kind": "Listing",
    "data": {
        "modhash": "pf7ljp9fjea06f2f446c845c5b93487f0a098bf996a0f3e82b",
        "children": [
            {
                "kind": "t5",
                "data": {
                    "submit_text_html": null,
                    "user_is_banned": false,
                    "id": "2qh0u",
                    "submit_text": "",
                    "display_name": "pics",
                    "header_img": "http://f.thumbs.redditmedia.com/5j2D-mwj6zafK81e.png",
                    "description_html": "<html>description1</html>",
                    "title": "/r/Pics",
                    "over18": false,
                    "user_is_moderator": false,
                    "header_title": "Logo by corvuskorax",
                    "description": "description1",
                    "submit_link_label": null,
                    "accounts_active": null,
                    "public_traffic": false,
                    "header_size": [
                        160,
                        64
                    ],
                    "subscribers": 5460468,
                    "submit_text_label": null,
                    "name": "t5_2qh0u",
                    "created": 1201224669,
                    "url": "/r/pics/",
                    "created_utc": 1201221069,
                    "user_is_contributor": false,
                    "public_description": "A place to share photographs and pictures.",
                    "comment_score_hide_mins": 60,
                    "subreddit_type": "public",
                    "submission_type": "link",
                    "user_is_subscriber": false
                }
            },
            {
                "kind": "t5",
                "data": {
                    "submit_text_html": null,
                    "user_is_banned": false,
                    "id": "2qh33",
                    "submit_text": "",
                    "display_name": "funny",
                    "header_img": "http://e.thumbs.redditmedia.com/g2Xn0gAOiibrx1j4.png",
                    "description_html": "<html>description2</html>",
                    "title": "funny",
                    "over18": false,
                    "user_is_moderator": false,
                    "header_title": "Logo by corvuskorax",
                    "description": "description2",
                    "submit_link_label": null,
                    "accounts_active": null,
                    "public_traffic": true,
                    "header_size": [
                        160,
                        64
                    ],
                    "subscribers": 5540358,
                    "submit_text_label": null,
                    "name": "t5_2qh33",
                    "created": 1201246556,
                    "url": "/r/funny/",
                    "created_utc": 1201242956,
                    "user_is_contributor": false,
                    "public_description": "",
                    "comment_score_hide_mins": 0,
                    "subreddit_type": "public",
                    "submission_type": "any",
                    "user_is_subscriber": false
                }
            },
            {
                "kind": "t5",
                "data": {
                    "submit_text_html": "<html>submit</html>",
                    "user_is_banned": false,
                    "id": "2qh03",
                    "submit_text": "submit",
                    "display_name": "gaming",
                    "header_img": "http://a.thumbs.redditmedia.com/vPDRq3ESKlnjqDff.png",
                    "description_html": "<html>description3</html>",
                    "over18": false,
                    "user_is_moderator": false,
                    "header_title": "High score!",
                    "description": "description3",
                    "submit_link_label": null,
                    "accounts_active": null,
                    "public_traffic": false,
                    "header_size": [
                        148,
                        120
                    ],
                    "subscribers": 4750358,
                    "submit_text_label": null,
                    "name": "t5_2qh03",
                    "created": 1190058205,
                    "url": "/r/gaming/",
                    "created_utc": 1190054605,
                    "user_is_contributor": false,
                    "public_description": "A subreddit for (almost) anything related to games - video games, board games, card games, etc. (but not [sports](http://www.reddit.com/r/sports)).\n\nFor more informative gaming content such as news and articles, please visit /r/Games.",
                    "comment_score_hide_mins": 0,
                    "subreddit_type": "public",
                    "submission_type": "any",
                    "user_is_subscriber": false
                }
            }
        ],
        "after": "t5_2qh03",
        "before": "blah blah made up"
    }
}
`

	subreddit3Posts string = `
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
            },
            {
                "kind": "t3",
                "data": {
                    "domain": "self.redditdev",
                    "banned_by": null,
                    "media_embed": {},
                    "subreddit": "redditdev",
                    "selftext_html": "<html>3</html>",
                    "likes": null,
                    "secure_media": null,
                    "link_flair_text": null,
                    "id": "20d5ol",
                    "gilded": 0,
                    "secure_media_embed": {},
                    "clicked": false,
                    "stickied": false,
                    "author": "amleszk",
                    "media": null,
                    "score": 3,
                    "approved_by": null,
                    "over_18": false,
                    "hidden": false,
                    "thumbnail": "",
                    "subreddit_id": "t5_2qizd",
                    "edited": false,
                    "link_flair_css_class": null,
                    "author_flair_css_class": null,
                    "downs": 0,
                    "saved": false,
                    "is_self": true,
                    "permalink": "/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
                    "name": "t3_20d5ol",
                    "created": 1394787026,
                    "url": "http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
                    "author_flair_text": null,
                    "title": "api/friend.json has no effect",
                    "created_utc": 1394758226,
                    "ups": 3,
                    "num_comments": 1,
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

	comments1Reply string = `
[
    {
        "kind": "Listing",
        "data": {
            "modhash": "r5n6qwsqvi32853d16880f9b42076c7b78a255e8b7f75a3f72",
            "children": [
                {
                    "kind": "t3",
                    "data": {
                        "domain": "self.redditdev",
                        "banned_by": null,
                        "media_embed": {},
                        "subreddit": "redditdev",
                        "selftext_html": "<html>",
                        "likes": null,
                        "secure_media": null,
                        "link_flair_text": null,
                        "id": "20d5ol",
                        "gilded": 0,
                        "secure_media_embed": {},
                        "clicked": false,
                        "stickied": false,
                        "author": "amleszk",
                        "media": null,
                        "score": 3,
                        "approved_by": null,
                        "over_18": false,
                        "hidden": false,
                        "thumbnail": "",
                        "subreddit_id": "t5_2qizd",
                        "edited": false,
                        "link_flair_css_class": null,
                        "author_flair_css_class": null,
                        "downs": 0,
                        "saved": false,
                        "is_self": true,
                        "permalink": "/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
                        "name": "t3_20d5ol",
                        "created": 1394787026,
                        "url": "http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
                        "author_flair_text": null,
                        "title": "api/friend.json has no effect",
                        "created_utc": 1394758226,
                        "ups": 3,
                        "num_comments": 1,
                        "visited": false,
                        "num_reports": null,
                        "distinguished": null
                    }
                }
            ],
            "after": null,
            "before": null
        }
    },
    {
        "kind": "Listing",
        "data": {
            "modhash": "r5n6qwsqvi32853d16880f9b42076c7b78a255e8b7f75a3f72",
            "children": [
                {
                    "kind": "t1",
                    "data": {
                        "subreddit_id": "t5_2qizd",
                        "banned_by": null,
                        "subreddit": "redditdev",
                        "likes": null,
                        "replies": "",
                        "saved": false,
                        "id": "cg2evwu",
                        "gilded": 0,
                        "author": "bsimpson",
                        "parent_id": "t3_20d5ol",
                        "approved_by": null,
                        "body": "that body",
                        "edited": false,
                        "author_flair_css_class": null,
                        "downs": 0,
                        "body_html": "<body>",
                        "link_id": "t3_20d5ol",
                        "score_hidden": false,
                        "name": "t1_cg2evwu",
                        "created": 1394833027,
                        "author_flair_text": null,
                        "created_utc": 1394804227,
                        "distinguished": null,
                        "num_reports": null,
                        "ups": 1
                    }
                }
            ],
            "after": null,
            "before": null
        }
    }
]`
)

type expectedThing struct {
	expectedType       ThingType
	expectedFieldCount uint
}

func verifyGroked(t *testing.T, actual *Groked, expectedPrev string, expectedNext string, expectedThings []expectedThing) {

	if actual == nil {
		t.Fatal("Actual should not be nil")
	}

	if actual.ListingPrev != expectedPrev {
		t.Error("Expected previous listing \"" + expectedPrev + "\" but got \"" + actual.ListingPrev + "\"")
	}

	if actual.ListingNext != expectedNext {
		t.Error("Expected next listing \"" + expectedNext + "\" but got \"" + actual.ListingNext + "\"")
	}

	if len(actual.Children) != len(expectedThings) {
		t.Fatal("Expected " + strconv.Itoa(len(expectedThings)) + " but got " + strconv.Itoa(len(actual.Children)))
	}

	for index, expectedThing := range expectedThings {
		if expectedThing.expectedType != actual.Children[index].Type {
			t.Error("Incorrect kind at index " + strconv.Itoa(index))
		}

		elements := make(map[string]interface{})
		if error := json.Unmarshal(actual.Children[index].data, &elements); error != nil {
			t.Error("Error unmarshalling data at index " + strconv.Itoa(index))
		}

		if expectedThing.expectedFieldCount != uint(len(elements)) {
			t.Error("Expected " + strconv.FormatUint(uint64(expectedThing.expectedFieldCount), 10) + " fields but got " + strconv.Itoa(len(elements)) + " at index " + strconv.Itoa(index))
		}
	}
}

func TestRedditPage(t *testing.T) {
	groked, error := GrokObject(strings.NewReader(reddit3))
	if error != nil {
		t.Error("Failed to parse object: " + error.Error())
	}

	expectedThings := make([]expectedThing, 0, 3)
	expectedThings = append(expectedThings, expectedThing{SubredditType, 28})
	expectedThings = append(expectedThings, expectedThing{SubredditType, 28})
	expectedThings = append(expectedThings, expectedThing{SubredditType, 27})

	verifyGroked(t, groked, "blah blah made up", "t5_2qh03", expectedThings)
}

func TestSubredditPage(t *testing.T) {
	groked, error := GrokObject(strings.NewReader(subreddit3Posts))
	if error != nil {
		t.Error("Failed to parse object: " + error.Error())
	}

	expectedThings := make([]expectedThing, 0, 3)
	expectedThings = append(expectedThings, expectedThing{LinkType, 40})
	expectedThings = append(expectedThings, expectedThing{LinkType, 39})
	expectedThings = append(expectedThings, expectedThing{LinkType, 39})

	verifyGroked(t, groked, "", "t3_20d5ol", expectedThings)
}

func TestCommentsPage(t *testing.T) {

	groked, error := GrokArray(strings.NewReader(comments1Reply))
	if error != nil {
		t.Error("Failed to parse array: " + error.Error())
	}

	if len(groked) != 2 {
		t.Fatal("Expected 2 groked items")
	}

	expectedThings := make([]expectedThing, 0, 1)
	expectedThings = append(expectedThings, expectedThing{LinkType, 39})

	verifyGroked(t, &groked[0], "", "", expectedThings)

	expectedThings[0].expectedType = CommentType
	expectedThings[0].expectedFieldCount = 25
	verifyGroked(t, &groked[1], "", "", expectedThings)
}
