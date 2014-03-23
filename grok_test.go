package grokeddit

import (
	"reflect"
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

func expectEqual(t *testing.T, expected interface{}, actual interface{}, failMessage string) bool {
	expectedType := reflect.ValueOf(expected)
	actualType := reflect.ValueOf(actual)

	if expectedType.Kind() != actualType.Kind() {
		t.Fatal(failMessage + ": Expected type \"" + expectedType.Kind().String() + "\" but got type \"" + actualType.Kind().String() + "\"")
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Error(failMessage + ": Expected \"" + expectedType.String() + "\" but got \"" + actualType.String() + "\"")
		return false
	}

	return true
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}, failMessage string) {
	if !expectEqual(t, expected, actual, failMessage) {
		t.Fatal("assertion requested, aborting test")
	}
}

func verifyGroked(t *testing.T, expected Groked, actual Groked) {

	expectEqual(t, expected.ListingPrev, actual.ListingPrev, "Previous listing error")
	expectEqual(t, expected.ListingNext, actual.ListingNext, "Next listing error")
	assertEqual(t, len(expected.Children), len(actual.Children), "Number of children error")
	
	for index, expectedThing := range expected.Children {
		errorMessage := " error at child index " + strconv.Itoa(index)
		expectEqual(t, expectedThing.Author, actual.Children[index].Author, "Author" +errorMessage)
		expectEqual(t, expectedThing.Created_utc, actual.Children[index].Created_utc, "Created timestamp" + errorMessage)
		expectEqual(t, expectedThing.ParentId, actual.Children[index].ParentId, "global id" + errorMessage)

		// recursively check replies
		verifyGroked(t, expectedThing.Replies, actual.Children[index].Replies)

		expectEqual(t, expectedThing.Subreddit, actual.Children[index].Subreddit, "Subreddit name" + errorMessage)
		expectEqual(t, expectedThing.SubredditId, actual.Children[index].SubredditId, "Subreddit id" + errorMessage)
		expectEqual(t, expectedThing.Text_html, actual.Children[index].Text_html, "Text html" + errorMessage)
		expectEqual(t, expectedThing.Title, actual.Children[index].Title, "Title" + errorMessage)
		expectEqual(t, expectedThing.Url, actual.Children[index].Url, "Url" + errorMessage)
	}
}

func TestRedditPage(t *testing.T) {
	actual, error := GrokListing(strings.NewReader(reddit3))
	if error != nil {
		t.Error("Failed to parse listing: " + error.Error())
	}

	expected := Groked{}

	expected.ListingPrev = "blah blah made up"
	expected.ListingNext = "t5_2qh03"
	expected.Children = make([]Thing, 3)

	expected.Children[0] = Thing{"", 1201221069, GlobalId{4594350, Subreddit}, 1201221069, GlobalId{}, Groked{}, "pics", GlobalId{4594350, Subreddit}, "", "/r/Pics", "/r/pics/"}
	expected.Children[1] = Thing{"", 1201242956, GlobalId{4594431, Subreddit}, 1201246556, GlobalId{}, Groked{}, "funny", GlobalId{4594431, Subreddit}, "", "funny", "/r/funny/"}
	expected.Children[2] = Thing{"", 1190054605, GlobalId{4594323, Subreddit}, 1190054605, GlobalId{}, Groked{}, "gaming", GlobalId{4594323, Subreddit}, "", "", "/r/gaming/"}

	verifyGroked(t, expected, actual)
}

func TestSubredditPage(t *testing.T) {
	actual, error := GrokListing(strings.NewReader(subreddit3Posts))
	if error != nil {
		t.Error("Failed to parse object: " + error.Error())
	}

	expected := Groked{}
	expected.ListingNext = "t3_20d5ol"
	expected.Children = make([]Thing, 3)

	expected.Children[0] = Thing{"reality_bugger", 1394831435, GlobalId{121660225, Link}, 1394831435, GlobalId{}, Groked{}, "redditdev", GlobalId{4596889, Subreddit}, "<html>1</html>", "Just a short thank you.", "http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/"}
	expected.Children[1] = Thing{"Grimzentide", 1394793933, GlobalId{121593987, Link}, 1394793933, GlobalId{}, Groked{}, "redditdev", GlobalId{4596889, Subreddit}, "<html>2</html>", "Is this the most efficient way to run this code?", "http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/"}
	expected.Children[2] = Thing{"amleszk", 1394758226, GlobalId{121546245, Link}, 1394758226, GlobalId{}, Groked{}, "redditdev", GlobalId{4596889, Subreddit}, "<html>3</html>", "api/friend.json has no effect", "http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/"}

	verifyGroked(t, expected, actual)
}

func TestCommentsPage(t *testing.T) {

	actual, error := GrokListingArray(strings.NewReader(comments1Reply))
	if error != nil {
		t.Error("Failed to parse array: " + error.Error())
	}

	assertEqual(t, len(actual), 2, "Groked array failure")

	expected := make([]Groked, 2)
	expected[0].Children = make([]Thing, 1)
	expected[0].Children[0] = Thing{"amleszk", 1394758226, GlobalId{121546245, Link}, 1394758226, GlobalId{}, Groked{}, "redditdev", GlobalId{4596889, Subreddit}, "<html>", "api/friend.json has no effect", "http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/"}

	expected[1].Children = make([]Thing, 1)
	expected[1].Children[0] = Thing{"bsimpson", 1394804227, GlobalId{27092900622, Comment}, 1394804227, GlobalId{121546245, Link}, Groked{}, "redditdev", GlobalId{4596889, Subreddit}, "<body>", "", ""}


	verifyGroked(t, expected[0], actual[0])
	verifyGroked(t, expected[1], actual[1])
}
