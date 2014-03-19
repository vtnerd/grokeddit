package grokeddit

import "encoding/json"

type Thing struct {
	Author       string          // name of the poster
	Created_utc  int64           // utc of creation time
	Id           string          // unique indentifier for the thing
	LastUpdate   int64           // utc of last update
	Parent_id    string          // parent id of a comment
	replies      json.RawMessage // unparsed replies to a comment
	Subreddit    string          // name of the subreddit associated with the thing
	Subreddit_id string          // id of the subreddit associated with the thing
	Kind         KindType
	Text_html    string // html from post
	Title        string // title of the post
	Url          string // url of the post
}
