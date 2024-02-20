package main

import (
	"log"
	"testing"

	"github.com/song940/feedparser-go/feed"
)

func TestRss(t *testing.T) {
	feed, err := feed.FetchRss("https://blog.lsong.org/feed.xml")
	if err != nil {
		t.Error(err)
	}
	log.Println(feed.Title)
}

func TestAtom(t *testing.T) {
	feed, err := feed.FetchAtom("https://qust.me/atom.xml")
	if err != nil {
		t.Error(err)
	}
	log.Println(feed.Title.Data)
}

func TestRssFailed(t *testing.T) {
	_, err := feed.FetchRss("https://qust.me/atom.xml")
	// should error
	log.Println(err)
}

func TestAtomFailed(t *testing.T) {
	_, err := feed.FetchAtom("https://blog.lsong.org/feed.xml")
	// should error
	log.Println(err)
}
