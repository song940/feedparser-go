package feed

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type RssGuid struct {
	Value       string `xml:",chardata"`
	IsPermaLink bool   `xml:"isPermaLink,attr,omitempty"`
}

type RssItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description string  `xml:"description"`
	PubDate     string  `xml:"pubDate"`
	Guid        RssGuid `xml:"guid"`
}

type RssFeed struct {
	XMLName     xml.Name  `xml:"rss"`
	Title       string    `xml:"channel>title"`
	Description string    `xml:"channel>description"`
	Link        string    `xml:"channel>link"`
	Items       []RssItem `xml:"channel>item"`
}

func ParseRss(data []byte) (feed *RssFeed, err error) {
	data = bytes.Map(func(r rune) rune {
		if r == '\u0008' {
			return -1
		}
		return r
	}, data)
	err = xml.Unmarshal(data, &feed)
	return
}

func FetchRss(url string) (feed *RssFeed, err error) {
	res, err := http.Get(url)
	if err != nil {
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return
	}
	feed, err = ParseRss(data)
	return
}
