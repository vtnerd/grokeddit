package fetcheddit

import "code.leeclagett.com/grokeddit"

// Represents a single reddit link.
type Link struct {
	Author        string             // name of the poster
	CreatedUtc    int64              // utc of creation time
	Id            grokeddit.GlobalId // uniquely indentifies the thing
	LastUpdateUtc int64              // utc of last update
	replies       grokeddit.Groked   // replies to a comment
	Subreddit     string             // name of the subreddit associated with the thing
	SubredditId   grokeddit.GlobalId // uniquely identifies the subreddit associated with the thing
	Text_html     string             // html from post
	Title         string             // title of the post
	Url           string             // url of the post
}
