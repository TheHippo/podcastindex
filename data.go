package podcastindex

import (
	"fmt"
	"strconv"
	"time"
)

// Time is a crutch to get the timestamp parsed correctly
// there is for obvious reasons no information on the timezone
type Time time.Time

// MarshalJSON is used to convert the timestamp to JSON
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *Time) UnmarshalJSON(s []byte) (err error) {
	u, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(u, 0)
	return nil
}

// String returns t as a formatted string
func (t Time) String() string {
	return time.Time(t).UTC().String()
}

// Duration is a crutch to get time.Duration parsed from the API
// results
type Duration time.Duration

// MarshalJSON is used to convert the duration to JSON
func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(int64(time.Duration(d).Seconds()), 10)), nil
}

// UnmarshalJSON is used to convert the duration from JSON
func (d *Duration) UnmarshalJSON(s []byte) (err error) {
	p, err := time.ParseDuration(fmt.Sprintf("%ss", s))
	if err != nil {
		return err
	}
	*(*time.Duration)(d) = p
	return nil
}

// String returns d as a formated string
func (d Duration) String() string {
	return time.Duration(d).String()
}

type podcastsResult struct {
	Status      string     `json:"status"`
	Feeds       []*Podcast `json:"feeds"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type podcastResult struct {
	Status string `json:"status"`
	Query  struct {
		URL string `json:"url"`
	} `json:"query"`
	Feed        Podcast `json:"feed"`
	Description string  `json:"description"`
}

// Podcast contains all informations about a podcast returned from the podcastindex API
type Podcast struct {
	ID                     uint            `json:"id"`
	Title                  string          `json:"title"`
	URL                    string          `json:"url"`
	OriginalURL            string          `json:"originalUrl"`
	Link                   string          `json:"link"`
	Description            string          `json:"description"`
	Author                 string          `json:"author"`
	OwnerName              string          `json:"ownerName"`
	Image                  string          `json:"image"`
	Artwork                string          `json:"artwork"`
	LastUpdateTime         Time            `json:"lastUpdateTime"`
	LastCrawlTime          Time            `json:"lastCrawlTime"`
	LastParseTime          Time            `json:"lastParseTime"`
	LastGoodHTTPStatusTime Time            `json:"lastGoodHttpStatusTime"`
	LastHTTPStatus         int             `json:"lastHttpStatus"`
	ContentType            string          `json:"contentType"`
	ItunesID               int             `json:"itunesId"`
	Generator              string          `json:"generator"`
	Language               string          `json:"language"`
	Type                   int             `json:"type"`
	Dead                   int             `json:"dead"`
	CrawlErrors            int             `json:"crawlErrors"`
	ParseErrors            int             `json:"parseErrors"`
	Categories             map[uint]string `json:"categories"`
}

type episodesResponse struct {
	Status      string     `json:"status"`
	Items       []*Episode `json:"items"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type randomEpisodesResponse struct {
	Status      string     `json:"status"`
	Items       []*Episode `json:"episodes"`
	Count       int        `json:"count"`
	Query       string     `json:"query"`
	Description string     `json:"description"`
}

type episodeResponse struct {
	Status      string   `json:"status"`
	ID          string   `json:"id"`
	Episode     *Episode `json:"episode"`
	Description string   `json:"description"`
}

// Episode contains all information about a single podcast episode returned from
// the podcastindex API
type Episode struct {
	ID              int      `json:"id"`
	Title           string   `json:"title"`
	Link            string   `json:"link"`
	Description     string   `json:"description"`
	GUID            string   `json:"guid"`
	DatePublished   Time     `json:"datePublished"`
	DateCrawled     Time     `json:"dateCrawled"`
	EnclosureURL    string   `json:"enclosureUrl"`
	EnclosureType   string   `json:"enclosureType"`
	EnclosureLength int      `json:"enclosureLength"`
	Duration        Duration `json:"duration"`
	Explicit        int      `json:"explicit"`
	Episode         int      `json:"episode"`
	EpisodeType     string   `json:"episodeType"`
	Season          int      `json:"season"`
	Image           string   `json:"image"`
	FeedItunesID    int      `json:"feedItunesId"`
	FeedImage       string   `json:"feedImage"`
	FeedID          int      `json:"feedId"`
	FeedLanguage    string   `json:"feedLanguage"`
	ChaptersURL     string   `json:"chaptersUrl"`
	TranscriptURL   string   `json:"transcriptUrl"`
}
