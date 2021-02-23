package podcastindex

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
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
//
// - fullBody to return the more then 100 characters in the descriptions
//
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

func (c *Client) getEpisodes(url string, notFound error) ([]*Episode, error) {
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}
	result := &episodesResponse{}
	err = decode(res, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, notFound
	}
	return result.Items, nil
}

// EpisodesByFeedID returns episodes for a podcast by its id
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByFeedID(id uint, max uint, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byfeedid?id=%d&fulltext%s%s", id, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by feed id"))
}

// EpisodesByFeedURL returns episodes for a podcast by its feed URL
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByFeedURL(feedURL string, max uint, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byfeedurl?url=\"%s\"&fulltext%s%s", feedURL, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by feed URL"))
}

// EpisodesByITunesID returns episodes for a podcast by its iTunes id
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned
//
// - since = only return episodes since that time. Set time to zero to not filter
// by time
func (c *Client) EpisodesByITunesID(id uint, max uint, since time.Time) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/byitunesid?id=%d&fulltext%s%s", id, addMax(max), addTime(since))
	return c.getEpisodes(url, errors.New("Could not get episodes by iTunes id"))
}

// EpisodeByID return a single episode by its id
func (c *Client) EpisodeByID(id uint) (*Episode, error) {
	url := fmt.Sprintf("episodes/byid?id=%d&fulltext", id)
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}
	result := &episodeResponse{}
	err = decode(res, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not find episode")
	}
	return result.Episode, nil
}

// RandomEpisodes returns a random episode
//
// categories and notCategories can be combined
//
// - languages = the languages the episodes should be in. "unknown" for when the language is not known.
// Leave empty if languages does not matter
//
// - categories = name of the category or categories the episodes should be in.
// Leave empty if categories do not matter
//
// - notCategories = name of the category or categories the episodes should not be in.
// Leave empty if categories do not matter
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned, the default is 1
func (c *Client) RandomEpisodes(languages, categories, notCategories []string, max uint) ([]*Episode, error) {
	url := fmt.Sprintf("episodes/random?fulltext%s%s%s%s", addMax(max), addFilter("lang", languages), addFilter("cat", categories), addFilter("notcat", notCategories))
	res, err := c.request(url)
	if err != nil {
		return nil, err
	}
	result := &randomEpisodesResponse{}
	err = decode(res, result)
	if err != nil {
		return nil, err
	}
	if result.Status == "false" {
		return nil, errors.New("Could not get random episodes")
	}
	return result.Items, nil
}

// RecentEpisodes returns the last episodes across the entire database
//
// - before = only return episodes that are older than the episode with this id. set to zero
// to ignore
//
// - excludeString = exclude episodes with this string in title or url. Leave empty for no
// filter
//
// - max = number of episodes to return, if max is 0 the default number of episodes will be
// returned, the default is 10
func (c *Client) RecentEpisodes(before uint, max uint, exclude string) ([]*Episode, error) {
	url := fmt.Sprintf("recent/episodes?fulltext%s%s%s", addMax(max), addExclude(exclude), addBefore(before))
	return c.getEpisodes(url, errors.New("Could not get recent episodes"))
}
