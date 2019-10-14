package aozora

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

const (
	defaultAPIDir = "api/v0.1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client *http.Client
	ctx    context.Context
}

//SearchBooksParamsFunc is self-referential function for functional options pattern
type SearchBooksParamsFunc func(url.Values)

//SearchBooksRaw gets list of books (raw data)
func (c *Client) SearchBooksRaw(opts ...SearchBooksParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	b, err := c.get(c.makeSearchCommand(TargetBooks, params))
	return b, errs.Wrap(err, "")
}

//SearchBooks gets list of books (struct data)
func (c *Client) SearchBooks(opts ...SearchBooksParamsFunc) ([]Book, error) {
	b, err := c.SearchBooksRaw(opts...)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	books, err := DecodeBooks(b)
	return books, errs.Wrap(err, "")
}

//WithBookTitle returns function for setting Marketplace
func WithBookTitle(title string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(title) > 0 {
			params.Set("title", title)
		}
	}
}

//WithBookAuthor returns function for setting Marketplace
func WithBookAuthor(author string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(author) > 0 {
			params.Set("author", author)
		}
	}
}

//WithBookFields returns function for setting Marketplace
func WithBookFields(fields string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(fields) > 0 {
			params.Add("fields", fields)
		}
	}
}

//WithBookLimit returns function for setting Marketplace
func WithBookLimit(limit int) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && limit > 0 {
			params.Set("limit", strconv.Itoa(limit))
		}
	}
}

//WithBookSkip returns function for setting Marketplace
func WithBookSkip(skip int) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && skip > 0 {
			params.Set("skip", strconv.Itoa(skip))
		}
	}
}

//WithBookAfter returns function for setting Marketplace
func WithBookAfter(after time.Time) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && !after.IsZero() {
			params.Set("after", after.Format("2006-01-02"))
		}
	}
}

//SearchPersonsParamsFunc is self-referential function for functional options pattern
type SearchPersonsParamsFunc func(url.Values)

//SearchPersonsRaw gets list of persons (raw data)
func (c *Client) SearchPersonsRaw(opts ...SearchPersonsParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	b, err := c.get(c.makeSearchCommand(TargetPersons, params))
	return b, errs.Wrap(err, "")
}

//SearchPersons gets list of persons (struct data)
func (c *Client) SearchPersons(opts ...SearchPersonsParamsFunc) ([]Person, error) {
	b, err := c.SearchPersonsRaw(opts...)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	persons, err := DecodePersons(b)
	return persons, errs.Wrap(err, "")
}

//WithPersonName returns function for setting Marketplace
func WithPersonName(name string) SearchPersonsParamsFunc {
	return func(params url.Values) {
		if params != nil && len(name) > 0 {
			params.Set("name", name)
		}
	}
}

//SearchPersonsParamsFunc is self-referential function for functional options pattern
type SearchWorkersParamsFunc func(url.Values)

//SearchWorkersRaw gets list of workers (raw data)
func (c *Client) SearchWorkersRaw(opts ...SearchWorkersParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	b, err := c.get(c.makeSearchCommand(TargetWorkers, params))
	return b, errs.Wrap(err, "")
}

//SearchWorkers gets list of workers (struct data)
func (c *Client) SearchWorkers(opts ...SearchWorkersParamsFunc) ([]Worker, error) {
	b, err := c.SearchWorkersRaw(opts...)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	workers, err := DecodeWorkers(b)
	return workers, errs.Wrap(err, "")
}

//WithWorkerName returns function for setting Marketplace
func WithWorkerName(name string) SearchWorkersParamsFunc {
	return func(params url.Values) {
		if params != nil && len(name) > 0 {
			params.Set("name", name)
		}
	}
}

//LookupBookRaw gets book data (raw data)
func (c *Client) LookupBookRaw(id int) ([]byte, error) {
	b, err := c.get(c.makeLookupCommand(TargetBooks, id))
	return b, errs.Wrap(err, "")
}

//LookupBook gets books data (struct data)
func (c *Client) LookupBook(id int) (*Book, error) {
	b, err := c.LookupBookRaw(id)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	book, err := DecodeBook(b)
	return book, errs.Wrap(err, "")
}

//LookupBookCardRaw gets book card info (HTML page data)
func (c *Client) LookupBookCardRaw(id int) ([]byte, error) {
	b, err := c.get(c.makeCardCommand(id))
	return b, errs.Wrap(err, "")
}

//LookupBookContentRaw gets book content (plain or HTML formatted text data)
func (c *Client) LookupBookContentRaw(id int, f Format) ([]byte, error) {
	b, err := c.get(c.makeContentCommand(id, f))
	return b, errs.Wrap(err, "")
}

//LookupPersonRaw gets person data (raw data)
func (c *Client) LookupPersonRaw(id int) ([]byte, error) {
	b, err := c.get(c.makeLookupCommand(TargetPersons, id))
	return b, errs.Wrap(err, "")
}

//LookupPerson gets person data (struct data)
func (c *Client) LookupPerson(id int) (*Person, error) {
	b, err := c.LookupPersonRaw(id)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	person, err := DecodePerson(b)
	return person, errs.Wrap(err, "")
}

//LookupWorker gets worker data (raw data)
func (c *Client) LookupWorkerRaw(id int) ([]byte, error) {
	b, err := c.get(c.makeLookupCommand(TargetWorkers, id))
	return b, errs.Wrap(err, "")
}

//LookupWorkerRaw gets worker data (struct data)
func (c *Client) LookupWorker(id int) (*Worker, error) {
	b, err := c.LookupWorkerRaw(id)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	worker, err := DecodeWorker(b)
	return worker, errs.Wrap(err, "")
}

//RankingRaw gets ranking data (raw data)
func (c *Client) RankingRaw(tm time.Time) ([]byte, error) {
	b, err := c.get(c.makeRankingCommand(tm))
	return b, errs.Wrap(err, "")
}

//Ranking gets ranking data (struct data)
func (c *Client) Ranking(tm time.Time) (Ranking, error) {
	b, err := c.RankingRaw(tm)
	if err != nil {
		return nil, errs.Wrap(err, "")
	}
	ranking, err := DecodeRanking(b)
	return ranking, errs.Wrap(err, "")
}

func (c *Client) makeSearchCommand(t Target, v url.Values) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v", c.apiDir(), t)
	u.RawQuery = v.Encode()
	return u
}

func (c *Client) makeLookupCommand(t Target, id int) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v/%v", c.apiDir(), t, strconv.Itoa(id))
	return u
}

func (c *Client) makeCardCommand(id int) *url.URL {
	u := c.makeLookupCommand(TargetBooks, id)
	u.Path = u.Path + "/card"
	return u
}

func (c *Client) makeContentCommand(id int, f Format) *url.URL {
	u := c.makeLookupCommand(TargetBooks, id)
	u.Path = u.Path + "/content"
	u.RawQuery = (url.Values{"format": {f.String()}}).Encode()
	return u
}

func (c *Client) makeRankingCommand(tm time.Time) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v/%v/%v", c.apiDir(), TargetRanking, "xhtml", tm.Format("2006/01"))
	return u
}

func (c *Client) apiDir() string {
	return defaultAPIDir
}

func (c *Client) get(u *url.URL) ([]byte, error) {
	req, err := http.NewRequestWithContext(c.ctx, "GET", u.String(), nil)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("url", u.String()))
	}
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithContext("url", u.String()))
	}
	defer resp.Body.Close()

	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return nil, errs.Wrap(ErrHTTPStatus, "", errs.WithContext("url", u.String()), errs.WithContext("status", resp.Status))
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errs.Wrap(err, "", errs.WithContext("url", u.String()))
	}
	return body, nil
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
