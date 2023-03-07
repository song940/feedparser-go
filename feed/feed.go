// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//
// This is a slightly modified version of 'encoding/xml/read_test.go'.
// Copyright 2009 The Go Authors. All rights reserved. Use of this
// source code is governed by a BSD-style license that can be found in
// the LICENSE file.

package feed

import (
	"bytes"
	"encoding/xml"
	"io"
	"time"
)

// Feed represents a generic feed.
type Feed struct {
	// Title for the feed.
	Title string

	// Feed type (either atom or rss).
	Type string

	// URL to the website.
	Link string

	// Description or subtitle for the feed.
	Description string

	// Categories the feed belongs to.
	Categories []string

	// Email address of the feed author.
	Author string

	// Last time the feed was updated.
	Updated time.Time

	// URL to image for the feed.
	Image string

	// Software used to generate the feed.
	Generator string

	// Information about rights, for example copyrights.
	Rights string

	// Feed Items
	Items []Item
}

// Item represents a generic feed item.
type Item struct {
	// Universally unique item ID.
	ID string

	// Title of the item.
	Title string

	// URL for the item.
	Link string

	// Content of the item.
	Content string

	// Email address of the item author.
	Author string

	// Categories the item belongs to.
	Categories []string

	// Time the item was published.
	PubDate time.Time

	// URL to media attachment.
	Attachment string
}

func replaceIllegalXML(b []byte) []byte {
	return bytes.Map(func(r rune) rune {
		if r == '\u0008' {
			return -1
		}
		return r
	}, b)
}

func UnmarshalXML(r io.Reader, v any) (err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		panic(err)
	}
	err = xml.Unmarshal(replaceIllegalXML(data), v)
	return err
}
