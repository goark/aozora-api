package aozora

import (
	"net/url"
	"testing"
	"time"
)

// func TestSearchBooks(t *testing.T) {
// 	testCases := []struct {
// 		title  string
// 		author string
// 		field1 string
// 		field2 string
// 		limit  int
// 		skip   int
// 		after  time.Time
// 		str    string
// 	}{
// 		{title: "title", author: "author", field1: "field1", field2: "field2", limit: 100, skip: 200, after: time.Now(), str: ""},
// 	}
//
// 	for _, tc := range testCases {
// 		s := DefaultClient().SearchBooksRaw(
// 			WithBookTitle(tc.title),
// 			WithBookAuthor(tc.author),
// 			WithBookFields(tc.field1),
// 			WithBookFields(tc.field2),
// 			WithBookLimit(tc.limit),
// 			WithBookSkip(tc.skip),
// 			WithBookAfter(tc.after),
// 		)
// 		if s != tc.str {
// 			t.Errorf("Client.MakeSearchCommand() is \"%v\", want \"%v\"", s, tc.str)
// 		}
// 	}
// }

func TestMakeSearchCommand(t *testing.T) {
	testCases := []struct {
		t   Target
		v   url.Values
		str string
	}{
		{t: Target(0), v: nil, str: "http://www.aozorahack.net/api/v0.1/unknown"},
		{t: TargetBooks, v: nil, str: "http://www.aozorahack.net/api/v0.1/books"},
		{t: TargetBooks, v: url.Values{}, str: "http://www.aozorahack.net/api/v0.1/books"},
		{t: TargetBooks, v: url.Values{"title": {"/foo/"}, "author": {"foo bar"}}, str: "http://www.aozorahack.net/api/v0.1/books?author=foo+bar&title=%2Ffoo%2F"},
		{t: TargetPersons, v: nil, str: "http://www.aozorahack.net/api/v0.1/persons"},
		{t: TargetPersons, v: url.Values{}, str: "http://www.aozorahack.net/api/v0.1/persons"},
		{t: TargetPersons, v: url.Values{"name": {"foo bar"}}, str: "http://www.aozorahack.net/api/v0.1/persons?name=foo+bar"},
		{t: TargetWorkers, v: nil, str: "http://www.aozorahack.net/api/v0.1/workers"},
		{t: TargetWorkers, v: url.Values{}, str: "http://www.aozorahack.net/api/v0.1/workers"},
		{t: TargetWorkers, v: url.Values{"name": {"foo bar"}}, str: "http://www.aozorahack.net/api/v0.1/workers?name=foo+bar"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeSearchCommand(tc.t, tc.v)
		if u.String() != tc.str {
			t.Errorf("Client.MakeSearchCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
}

func TestMakeLookupCommand(t *testing.T) {
	testCases := []struct {
		t   Target
		id  string
		str string
	}{
		{t: TargetBooks, id: "1234", str: "http://www.aozorahack.net/api/v0.1/books/1234"},
		{t: TargetPersons, id: "1234", str: "http://www.aozorahack.net/api/v0.1/persons/1234"},
		{t: TargetWorkers, id: "1234", str: "http://www.aozorahack.net/api/v0.1/workers/1234"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeLookupCommand(tc.t, tc.id)
		if u.String() != tc.str {
			t.Errorf("Client.MakeLookupCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
}

func TestMakeCardCommand(t *testing.T) {
	testCases := []struct {
		id  string
		str string
	}{
		{id: "1234", str: "http://www.aozorahack.net/api/v0.1/books/1234/card"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeCardCommand(tc.id)
		if u.String() != tc.str {
			t.Errorf("Client.MakeCardCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
}

func TestMakeContentCommand(t *testing.T) {
	testCases := []struct {
		id  string
		f   Format
		str string
	}{
		{id: "1234", f: Format(0), str: "http://www.aozorahack.net/api/v0.1/books/1234/content?format=txt"},
		{id: "1234", f: Text, str: "http://www.aozorahack.net/api/v0.1/books/1234/content?format=txt"},
		{id: "1234", f: HTML, str: "http://www.aozorahack.net/api/v0.1/books/1234/content?format=html"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeContentCommand(tc.id, tc.f)
		if u.String() != tc.str {
			t.Errorf("Client.MakeContentCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
}

func TestMakeRankingCommand(t *testing.T) {
	testCases := []struct {
		tm  time.Time
		str string
	}{
		{tm: time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC), str: "http://www.aozorahack.net/api/v0.1/ranking/xhtml/2019/01"},
	}

	for _, tc := range testCases {
		u := DefaultClient().MakeRankingCommand(tc.tm)
		if u.String() != tc.str {
			t.Errorf("Client.MakeRankingCommand() is \"%v\", want \"%v\"", u.String(), tc.str)
		}
	}
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
