package fetcheddit

import "io"

/* Interface for fetching reddit content. The interface allows for non-HTTP
sources, or non reddit.com domains to be used to retrieve content. */
type Fetcher interface {
	/* Implementations must fetch the resource requested via the string
	   method (it will be a relative path), and return a ReadCloser to that
	   resource. */
	Fetch(string) (io.ReadCloser, error)
}
