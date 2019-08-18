package aozora

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/spiegel-im-spiegel/errs"
)

const (
	APIVersion = "api/v0.1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client *http.Client
}

//SearchBooksParamsFunc is self-referential function for functional options pattern
type SearchBooksParamsFunc func(url.Values)

//SearchBooksRaw gets list of books (raw data)
func (c *Client) SearchBooksRaw(opts ...SearchBooksParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	return c.get(c.MakeSearchCommand(TargetBooks, params))
}

//SearchBooks gets list of books (struct data)
func (c *Client) SearchBooks(opts ...SearchBooksParamsFunc) ([]Book, error) {
	b, err := c.SearchBooksRaw(opts...)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.SearchBooks() function")
	}
	return DecodeBooks(b)
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
	return c.get(c.MakeSearchCommand(TargetPersons, params))
}

//SearchPersons gets list of persons (struct data)
func (c *Client) SearchPersons(opts ...SearchPersonsParamsFunc) ([]Person, error) {
	b, err := c.SearchPersonsRaw(opts...)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.SearchPersons() function")
	}
	return DecodePersons(b)
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
	return c.get(c.MakeSearchCommand(TargetWorkers, params))
}

//SearchWorkers gets list of workers (struct data)
func (c *Client) SearchWorkers(opts ...SearchWorkersParamsFunc) ([]Worker, error) {
	b, err := c.SearchWorkersRaw(opts...)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.SearchWorkers() function")
	}
	return DecodeWorkers(b)
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
	return c.get(c.MakeLookupCommand(TargetBooks, id))
}

//LookupBook gets books data (struct data)
func (c *Client) LookupBook(id int) (*Book, error) {
	b, err := c.LookupBookRaw(id)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.LookupBookRaw() function")
	}
	return DecodeBook(b)
}

//LookupBookCardRaw gets book card info (HTML page data)
func (c *Client) LookupBookCardRaw(id int) ([]byte, error) {
	return c.get(c.MakeCardCommand(id))
}

//LookupBookContentRaw gets book content (plain or HTML formatted text data)
func (c *Client) LookupBookContentRaw(id int, f Format) ([]byte, error) {
	return c.get(c.MakeContentCommand(id, f))
}

//LookupPersonRaw gets person data (raw data)
func (c *Client) LookupPersonRaw(id int) ([]byte, error) {
	return c.get(c.MakeLookupCommand(TargetPersons, id))
}

//LookupPerson gets person data (struct data)
func (c *Client) LookupPerson(id int) (*Person, error) {
	b, err := c.LookupPersonRaw(id)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.LookupPerson() function")
	}
	return DecodePerson(b)
}

//LookupWorker gets worker data (raw data)
func (c *Client) LookupWorkerRaw(id int) ([]byte, error) {
	return c.get(c.MakeLookupCommand(TargetWorkers, id))
}

//LookupWorkerRaw gets worker data (struct data)
func (c *Client) LookupWorker(id int) (*Worker, error) {
	b, err := c.LookupWorkerRaw(id)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.LookupWorker() function")
	}
	return DecodeWorker(b)
}

//RankingRaw gets ranking data (raw data)
func (c *Client) RankingRaw(tm time.Time) ([]byte, error) {
	return c.get(c.MakeRankingCommand(tm))
}

//Ranking gets ranking data (struct data)
func (c *Client) Ranking(tm time.Time) (Ranking, error) {
	b, err := c.RankingRaw(tm)
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.Ranking() function")
	}
	return DecodeRanking(b)
}

//MakeSearchCommand returns URI for search command
func (c *Client) MakeSearchCommand(t Target, v url.Values) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v", APIVersion, t)
	u.RawQuery = v.Encode()
	return u
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeLookupCommand(t Target, id int) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v/%v", APIVersion, t, strconv.Itoa(id))
	return u
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeCardCommand(id int) *url.URL {
	u := c.MakeLookupCommand(TargetBooks, id)
	u.Path = u.Path + "/card"
	return u
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeContentCommand(id int, f Format) *url.URL {
	u := c.MakeLookupCommand(TargetBooks, id)
	u.Path = u.Path + "/content"
	u.RawQuery = (url.Values{"format": {f.String()}}).Encode()
	return u
}

//MakeLookupCommand returns URI for lookup ranking info command
func (c *Client) MakeRankingCommand(tm time.Time) *url.URL {
	u := c.server.URL()
	u.Path = fmt.Sprintf("/%v/%v/%v/%v", APIVersion, TargetRanking, "xhtml", tm.Format("2006/01"))
	return u
}

func (c *Client) get(u *url.URL) ([]byte, error) {
	resp, err := c.client.Get(u.String())
	if err != nil {
		return nil, errs.Wrapf(err, "error in Client.get(%v) function", u)
	}
	defer resp.Body.Close()

	if !(resp.StatusCode != 0 && resp.StatusCode < http.StatusBadRequest) {
		return nil, errs.Wrapf(ErrHTTPStatus, "%v (in %v)", resp.Status, u)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, errs.Wrapf(err, "error in Client.get(%v) function", u)
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
