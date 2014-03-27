package fetcheddit

import "code.leeclagett.com/grokeddit"

// Represents a single reddit comment.
type Comment struct {
	Author        string             // name of the poster
	CreatedUtc    int64              // utc of creation time
	Id            grokeddit.GlobalId // uniquely indentifies the thing
	LastUpdateUtc int64              // utc of last update
	ParentId      grokeddit.GlobalId // uniquely identifies the parent to this comment
	replies       grokeddit.Groked   // replies to a comment
	Subreddit     string             // name of the subreddit associated with the thing
	SubredditId   grokeddit.GlobalId // uniquely identifies the subreddit associated with the thing
	TextHtml      string             // html from post
}
