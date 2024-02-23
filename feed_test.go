package main

import (
	"log"
	"testing"

	"github.com/song940/feedparser-go/feed"
)

func TestRss(t *testing.T) {
	feed, err := feed.FetchRss("https://www.luqijian.com/feed/")
	if err != nil {
		t.Error(err)
	}
	log.Println(feed.Title)
	for _, item := range feed.Items {
		log.Println(item.ContentEncoded)
	}
}
