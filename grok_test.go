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
        "before": "t5_2qh00"
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

	commentTree string = `
[
    {
        "kind": "Listing",
        "data": {
            "modhash": "73e9l7288vd5cbae00ea93c692511c4634a2d9ed81578c0390",
            "children": [
                {
                    "kind": "t3",
                    "data": {
                        "domain": "self.askscience",
                        "banned_by": null,
                        "media_embed": {},
                        "subreddit": "askscience",
                        "selftext_html": "<html>og link text</html>",
                        "selftext": "og link text",
                        "likes": null,
                        "secure_media": null,
                        "link_flair_text": "Biology",
                        "id": "214czs",
                        "gilded": 0,
                        "secure_media_embed": {},
                        "clicked": false,
                        "stickied": false,
                        "author": "drumersrule",
                        "media": null,
                        "score": 760,
                        "approved_by": null,
                        "over_18": false,
                        "hidden": false,
                        "thumbnail": "",
                        "subreddit_id": "t5_2qm4e",
                        "edited": false,
                        "link_flair_css_class": "bio",
                        "author_flair_css_class": null,
                        "downs": 373,
                        "saved": false,
                        "is_self": true,
                        "permalink": "/r/askscience/comments/214czs/do_offspring_ever_take_care_of_their_parents_in/",
                        "name": "t3_214czs",
                        "created": 1395570389,
                        "url": "http://www.reddit.com/r/askscience/comments/214czs/do_offspring_ever_take_care_of_their_parents_in/",
                        "author_flair_text": null,
                        "title": "Do offspring ever take care of their parents in other species?",
                        "created_utc": 1395541589,
                        "ups": 1133,
                        "num_comments": 137,
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
            "modhash": "73e9l7288vd5cbae00ea93c692511c4634a2d9ed81578c0390",
            "children": [
                {
                    "kind": "t1",
                    "data": {
                        "subreddit_id": "t5_2qm4e",
                        "banned_by": null,
                        "subreddit": "askscience",
                        "likes": null,
                        "replies": {
                            "kind": "Listing",
                            "data": {
                                "modhash": "73e9l7288vd5cbae00ea93c692511c4634a2d9ed81578c0390",
                                "children": [
                                    {
                                        "kind": "t1",
                                        "data": {
                                            "subreddit_id": "t5_2qm4e",
                                            "banned_by": null,
                                            "subreddit": "askscience",
                                            "likes": null,
                                            "replies": "",
                                            "saved": false,
                                            "id": "cg9qapl",
                                            "gilded": 0,
                                            "author": "PM_ME_YOUR_NIGHTMARE",
                                            "parent_id": "t1_cg9ptzf",
                                            "approved_by": null,
                                            "body": "first reply to comment",
                                            "edited": false,
                                            "author_flair_css_class": null,
                                            "downs": 0,
                                            "body_html": "<html>first reply to comment</html>",
                                            "link_id": "t3_214czs",
                                            "score_hidden": false,
                                            "name": "t1_cg9qapl",
                                            "created": 1395610381,
                                            "author_flair_text": null,
                                            "created_utc": 1395581581,
                                            "distinguished": null,
                                            "num_reports": null,
                                            "ups": 3
                                        }
                                    },
                                    {
                                        "kind": "t1",
                                        "data": {
                                            "subreddit_id": "t5_2qm4e",
                                            "banned_by": null,
                                            "subreddit": "askscience",
                                            "likes": null,
                                            "replies": {
                                                "kind": "Listing",
                                                "data": {
                                                    "modhash": "73e9l7288vd5cbae00ea93c692511c4634a2d9ed81578c0390",
                                                    "children": [
                                                        {
                                                            "kind": "t1",
                                                            "data": {
                                                                "subreddit_id": "t5_2qm4e",
                                                                "banned_by": null,
                                                                "subreddit": "askscience",
                                                                "likes": null,
                                                                "replies": "",
                                                                "saved": false,
                                                                "id": "cg9s0vd",
                                                                "gilded": 0,
                                                                "author": "inderstube",
                                                                "parent_id": "t1_cg9qv5w",
                                                                "approved_by": null,
                                                                "body": "first reply to second comment",
                                                                "edited": false,
                                                                "author_flair_css_class": null,
                                                                "downs": 0,
                                                                "body_html": "<html>first reply to second comment</html>",
                                                                "link_id": "t3_214czs",
                                                                "score_hidden": false,
                                                                "name": "t1_cg9s0vd",
                                                                "created": 1395616509,
                                                                "author_flair_text": null,
                                                                "created_utc": 1395587709,
                                                                "distinguished": null,
                                                                "num_reports": null,
                                                                "ups": 1
                                                            }
                                                        }
                                                    ],
                                                    "after": null,
                                                    "before": "t5_2qh04"
                                                }
                                            },
                                            "saved": false,
                                            "id": "cg9qv5w",
                                            "gilded": 0,
                                            "author": "Dave37",
                                            "parent_id": "t1_cg9ptzf",
                                            "approved_by": null,
                                            "body": "second reply to comment",
                                            "edited": false,
                                            "author_flair_css_class": null,
                                            "downs": 0,
                                            "body_html": "<html>second reply to comment</html>",
                                            "link_id": "t3_214czs",
                                            "score_hidden": false,
                                            "name": "t1_cg9qv5w",
                                            "created": 1395612719,
                                            "author_flair_text": null,
                                            "created_utc": 1395583919,
                                            "distinguished": null,
                                            "num_reports": null,
                                            "ups": 3
                                        }
                                    }
                                ],
                                "after": "t5_2qh03",
                                "before": null
                            }
                        },
                        "saved": false,
                        "id": "cg9ptzf",
                        "gilded": 0,
                        "author": "inderstube",
                        "parent_id": "t3_214czs",
                        "approved_by": null,
                        "body": "random comment in comment tree",
                        "edited": false,
                        "author_flair_css_class": null,
                        "downs": 6,
                        "body_html": "<html>random comment in comment tree</html>",
                        "link_id": "t3_214czs",
                        "score_hidden": false,
                        "name": "t1_cg9ptzf",
                        "created": 1395608112,
                        "author_flair_text": null,
                        "created_utc": 1395579312,
                        "distinguished": null,
                        "num_reports": null,
                        "ups": 3
                    }
                }
            ],
            "after": null,
            "before": null
        }
    }
]`

	commentListing string = `
{
    "kind": "Listing",
    "data": {
        "modhash": "9thsdknfa49c38eb55e7db72adf564df864c684bbb2a28efab",
        "children": [
            {
                "kind": "t1",
                "data": {
                    "subreddit_id": "t5_2qm4e",
                    "link_title": "Can the Casimir effect take place near a black hole?",
                    "banned_by": null,
                    "subreddit": "askscience",
                    "link_author": "Blocksy",
                    "likes": null,
                    "replies": null,
                    "saved": false,
                    "id": "cg9tfgc",
                    "gilded": 0,
                    "author": "babeltoothe",
                    "parent_id": "t1_cg9s6hv",
                    "approved_by": null,
                    "body": "comment 1",
                    "edited": false,
                    "author_flair_css_class": null,
                    "downs": 0,
                    "body_html": "<html>comment 1</html>",
                    "link_id": "t3_2156gs",
                    "score_hidden": false,
                    "name": "t1_cg9tfgc",
                    "created": 1395620331,
                    "author_flair_text": null,
                    "link_url": "http://www.reddit.com/r/askscience/comments/2156gs/can_the_casimir_effect_take_place_near_a_black/",
                    "created_utc": 1395591531,
                    "ups": 1,
                    "num_reports": null,
                    "distinguished": null
                }
            },
            {
                "kind": "t1",
                "data": {
                    "subreddit_id": "t5_2qm4e",
                    "link_title": "Is there any knowledge on how it has affected evolution that some genes are dominant, and some recessive?",
                    "banned_by": null,
                    "subreddit": "askscience",
                    "link_author": "throwaway774829",
                    "likes": null,
                    "replies": null,
                    "saved": false,
                    "id": "cg9te8o",
                    "gilded": 0,
                    "author": "Apiphilia",
                    "parent_id": "t3_21539t",
                    "approved_by": null,
                    "body": "comment 2",
                    "edited": false,
                    "author_flair_css_class": null,
                    "downs": 0,
                    "body_html": "<html>comment 2</html>",
                    "link_id": "t3_21539t",
                    "score_hidden": false,
                    "name": "t1_cg9te8o",
                    "created": 1395620242,
                    "author_flair_text": null,
                    "link_url": "http://www.reddit.com/r/askscience/comments/21539t/is_there_any_knowledge_on_how_it_has_affected/",
                    "created_utc": 1395591442,
                    "ups": 1,
                    "num_reports": null,
                    "distinguished": null
                }
            },
            {
                "kind": "t1",
                "data": {
                    "subreddit_id": "t5_2qm4e",
                    "link_title": "Transcendental numbers in a different base?",
                    "banned_by": null,
                    "subreddit": "askscience",
                    "link_author": "AlexxTheKid",
                    "likes": null,
                    "replies": null,
                    "saved": false,
                    "id": "cg9tdui",
                    "gilded": 0,
                    "author": "AlexxTheKid",
                    "parent_id": "t1_cg9mvtk",
                    "approved_by": null,
                    "body": "comment 3",
                    "edited": false,
                    "author_flair_css_class": null,
                    "downs": 0,
                    "body_html": "<html>comment 3</html>",
                    "link_id": "t3_2148x2",
                    "score_hidden": false,
                    "name": "t1_cg9tdui",
                    "created": 1395620214,
                    "author_flair_text": null,
                    "link_url": "http://www.reddit.com/r/askscience/comments/2148x2/transcendental_numbers_in_a_different_base/",
                    "created_utc": 1395591414,
                    "ups": 1,
                    "num_reports": null,
                    "distinguished": null
                }
            }
        ],
        "after": "t1_cg9tdui",
        "before": null
    }
}`
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
		expectEqual(t, expectedThing.Author, actual.Children[index].Author, "Author"+errorMessage)
		expectEqual(t, expectedThing.CreatedUtc, actual.Children[index].CreatedUtc, "Created timestamp"+errorMessage)
		expectEqual(t, expectedThing.Id, actual.Children[index].Id, "id"+errorMessage)
		expectEqual(t, expectedThing.LastUpdateUtc, actual.Children[index].LastUpdateUtc, "Last update timestamp"+errorMessage)
		expectEqual(t, expectedThing.LinkId, actual.Children[index].LinkId, "link id" + errorMessage)
		expectEqual(t, expectedThing.ParentId, actual.Children[index].ParentId, "parent id" + errorMessage)

		// recursively check replies
		verifyGroked(t, expectedThing.Replies, actual.Children[index].Replies)

		expectEqual(t, expectedThing.Subreddit, actual.Children[index].Subreddit, "Subreddit name"+errorMessage)
		expectEqual(t, expectedThing.SubredditId, actual.Children[index].SubredditId, "Subreddit id"+errorMessage)
		expectEqual(t, expectedThing.Text_html, actual.Children[index].Text_html, "Text html"+errorMessage)
		expectEqual(t, expectedThing.Title, actual.Children[index].Title, "Title"+errorMessage)
		expectEqual(t, expectedThing.Url, actual.Children[index].Url, "Url"+errorMessage)
	}
}

func TestRedditPage(t *testing.T) {
	actual, error := GrokListing(strings.NewReader(reddit3))
	if error != nil {
		t.Error("Failed to parse listing: " + error.Error())
	}

	expected := Groked{
		&GlobalId{4594320, Subreddit},
		&GlobalId{4594323, Subreddit},
		[]Thing{
			Thing{
				"",
				1201221069,
				GlobalId{4594350, Subreddit},
				1201221069,
				GlobalId{},
				GlobalId{},
				Groked{},
				"pics",
				GlobalId{4594350, Subreddit},
				"",
				"/r/Pics",
				"/r/pics/",
			},
			Thing{
				"",
				1201242956,
				GlobalId{4594431, Subreddit},
				1201242956,
				GlobalId{},
				GlobalId{},
				Groked{},
				"funny",
				GlobalId{4594431, Subreddit},
				"",
				"funny",
				"/r/funny/",
			},
			Thing{
				"",
				1190054605,
				GlobalId{4594323, Subreddit},
				1190054605,
				GlobalId{},
				GlobalId{},
				Groked{},
				"gaming",
				GlobalId{4594323, Subreddit},
				"",
				"",
				"/r/gaming/",
			},
		},
	}

	verifyGroked(t, expected, actual)
}

func TestSubredditPage(t *testing.T) {
	actual, error := GrokListing(strings.NewReader(subreddit3Posts))
	if error != nil {
		t.Error("Failed to parse object: " + error.Error())
	}

	expected := Groked{
		nil,
		&GlobalId{121546245, Link},
		[]Thing{
			Thing{
				"reality_bugger",
				1394831435,
				GlobalId{121660225, Link},
				1394831435,
				GlobalId{},
				GlobalId{},
				Groked{},
				"redditdev",
				GlobalId{4596889, Subreddit},
				"<html>1</html>",
				"Just a short thank you.",
				"http://www.reddit.com/r/redditdev/comments/20flmp/just_a_short_thank_you/",
			},
			Thing{
				"Grimzentide",
				1394793933,
				GlobalId{121593987, Link}, 1394793933,
				GlobalId{},
				GlobalId{},
				Groked{},
				"redditdev",
				GlobalId{4596889, Subreddit},
				"<html>2</html>",
				"Is this the most efficient way to run this code?",
				"http://www.reddit.com/r/redditdev/comments/20e6ir/is_this_the_most_efficient_way_to_run_this_code/",
			},
			Thing{
				"amleszk",
				1394758226,
				GlobalId{121546245, Link},
				1394758226,
				GlobalId{},
				GlobalId{},
				Groked{},
				"redditdev",
				GlobalId{4596889, Subreddit},
				"<html>3</html>",
				"api/friend.json has no effect", "http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
			},
		},
	}

	verifyGroked(t, expected, actual)
}

func TestCommentsPageWith1Comment(t *testing.T) {

	actual, error := GrokListingArray(strings.NewReader(comments1Reply))
	if error != nil {
		t.Error("Failed to parse array: " + error.Error())
	}

	assertEqual(t, len(actual), 2, "Groked array failure")

	expected := []Groked{
		Groked{
			nil,
			nil,
			[]Thing{
				Thing{
					"amleszk",
					1394758226,
					GlobalId{121546245, Link},
					1394758226,
					GlobalId{},
					GlobalId{},
					Groked{},
					"redditdev",
					GlobalId{4596889, Subreddit},
					"<html>",
					"api/friend.json has no effect",
					"http://www.reddit.com/r/redditdev/comments/20d5ol/apifriendjson_has_no_effect/",
				},
			},
		},
		Groked{
			nil,
			nil,
			[]Thing{
				Thing{
					"bsimpson",
					1394804227,
					GlobalId{27092900622, Comment},
					1394804227,
					GlobalId{121546245, Link},
					GlobalId{121546245, Link},
					Groked{},
					"redditdev",
					GlobalId{4596889, Subreddit},
					"<body>",
					"",
					"",
				},
			},
		},
	}

	verifyGroked(t, expected[0], actual[0])
	verifyGroked(t, expected[1], actual[1])
}

func TestCommentTree(t *testing.T) {
	actual, error := GrokListingArray(strings.NewReader(commentTree))
	if error != nil {
		t.Error("Failed to parse array: " + error.Error())
	}

	assertEqual(t, len(actual), 2, "Groked array failure")

	expected := []Groked{
		Groked{
			nil,
			nil,
			[]Thing{
				Thing{
					"drumersrule",
					1395541589,
					GlobalId{122815432, Link},
					1395541589,
					GlobalId{},
					GlobalId{},
					Groked{},
					"askscience",
					GlobalId{4600958, Subreddit},
					"<html>og link text</html>",
					"Do offspring ever take care of their parents in other species?",
					"http://www.reddit.com/r/askscience/comments/214czs/do_offspring_ever_take_care_of_their_parents_in/",
				},
			},
		},
		Groked{
			nil,
			nil,
			[]Thing{
				Thing{
					"inderstube",
					1395579312,
					GlobalId{27105168651, Comment},
					1395579312,
					GlobalId{122815432, Link},
					GlobalId{122815432, Link},
					Groked{
						nil,
						&GlobalId{4594323, Subreddit},
						[]Thing{
							Thing{
								"PM_ME_YOUR_NIGHTMARE",
								1395581581,
								GlobalId{27105190329, Comment},
								1395581581,
								GlobalId{122815432, Link},
								GlobalId{27105168651, Comment},
								Groked{},
								"askscience",
								GlobalId{4600958, Subreddit},
								"<html>first reply to comment</html>",
								"",
								"",
							},
							Thing{
								"Dave37",
								1395583919,
								GlobalId{27105216836, Comment},
								1395583919,
								GlobalId{122815432, Link},
								GlobalId{27105168651, Comment},
								Groked{
									&GlobalId{4594324, Subreddit},
									nil,
									[]Thing{
										Thing{
											"inderstube",
											1395587709,
											GlobalId{27105270889, Comment},
											1395587709,
											GlobalId{122815432, Link},
											GlobalId{27105216836, Comment},
											Groked{},
											"askscience",
											GlobalId{4600958, Subreddit},
											"<html>first reply to second comment</html>",
											"",
											"",
										},
									},
								},
								"askscience",
								GlobalId{4600958, Subreddit},
								"<html>second reply to comment</html>",
								"",
								"",
							},
						},
					},
					"askscience",
					GlobalId{4600958, Subreddit},
					"<html>random comment in comment tree</html>",
					"",
					"",
				},
			},
		},
	}

	verifyGroked(t, expected[0], actual[0])
	verifyGroked(t, expected[1], actual[1])
}

func TestCommentListing(t *testing.T) {

	actual, error := GrokListing(strings.NewReader(commentListing))
	if error != nil {
		t.Error("Failed to parse array: " + error.Error())
	}

	expected := Groked{
		nil,
		&GlobalId{27105334362, Comment},
		[]Thing{
			Thing{
				"babeltoothe",
				1395591531,
				GlobalId{27105336444, Comment},
				1395591531,
				GlobalId{122853628, Link},
				GlobalId{27105278179, Comment},
				Groked{},
				"askscience",
				GlobalId{4600958, Subreddit},
				"<html>comment 1</html>",
				"",
				"",
			},
			Thing{
				"Apiphilia",
				1395591442,
				GlobalId{27105334872, Comment},
				1395591442,
				GlobalId{122849489, Link},
				GlobalId{122849489, Link},
				Groked{},
				"askscience",
				GlobalId{4600958, Subreddit},
				"<html>comment 2</html>",
				"",
				"",
			},
			Thing{
				"AlexxTheKid",
				1395591414,
				GlobalId{27105334362, Comment},
				1395591414,
				GlobalId{122810150, Link},
				GlobalId{27105031064, Comment},
				Groked{},
				"askscience",
				GlobalId{4600958, Subreddit},
				"<html>comment 3</html>",
				"",
				"",
			},
		},
	}

	verifyGroked(t, expected, actual)
}
