package www2json

import (
	"errors"

	"github.com/gilliek/go-feedsfinder/feeds"
)

// FeedFromLink return feed link if found
func FeedFromLink(link string) (string, error) {
	feeds, err := feeds.FindFromURL(link)
	if err != nil {
		return "", err
	}
	for _, feed := range feeds {
		return feed.URL, nil
	}
	return "", errors.New("Error: feed not found")
}
