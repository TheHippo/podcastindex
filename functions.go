package podcastindex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

func decode(in []byte, out interface{}) error {
	decoder := json.NewDecoder(bytes.NewReader(in))
	return decoder.Decode(out)
}

// Search for podcasts, authors or owners
func (c *Client) Search(term string) ([]*Podcast, error) {
	return c.SearchC(term, false, 0)
}

// SearchC for searching with more options than Search
//
// - clean for non explicit feeds according to itunes:explicit
// - fullBody to return the more then 100 characters in the descriptions
// - max for the number of results, when set to 0 it uses the API default
func (c *Client) SearchC(term string, clean bool, max uint) ([]*Podcast, error) {
	url := fmt.Sprintf("search/byterm?q=\"%s\"&fulltext%s%s", term, addClean(clean), addMax(max))
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}
	result := &podcastsResult{}
	err = decode(res, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find a podcast for that term")
	}
	return result.Feeds, err
}

func (c *Client) getPodcast(url string, notFound error) (*Podcast, error) {
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}
	result := &podcastResult{}
	err = decode(res, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, notFound
	}
	return &result.Feed, err
}

// PodcastByFeedURL returns general information about a podcast by its
// feed URL
func (c *Client) PodcastByFeedURL(url string) (*Podcast, error) {
	u := fmt.Sprintf("podcasts/byfeedurl?url=%s&fulltext", url)
	return c.getPodcast(u, errors.New("Could not find a podcast for that feed URL"))
}

// PodcastByFeedID returns general information about a podcast by its id
func (c *Client) PodcastByFeedID(id uint) (*Podcast, error) {
	url := fmt.Sprintf("podcasts/byfeedid?id=%d&fulltext", id)
	return c.getPodcast(url, errors.New("Could not find a podcast for that id"))
}

// PodcastByITunesID returns general information about a podcast by its
// ITune id
func (c *Client) PodcastByITunesID(id uint) (*Podcast, error) {
	url := fmt.Sprintf("podcasts/byitunesid?id=%d&fulltext", id)
	return c.getPodcast(url, errors.New("Could not find a podcast for that iTunes id"))
}
