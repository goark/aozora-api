package aozora

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/goark/errs"
	"github.com/goark/fetch"
)

const (
	defaultAPIDir = "api/v0.1"
)

//Client is http.Client for Aozora API Server.
type Client struct {
	server *Server
	client fetch.Client
}

//SearchBooksParamsFunc is self-referential function for functional options pattern.
type SearchBooksParamsFunc func(url.Values)

//SearchBooksRaw gets list of books. (raw data)
func (c *Client) SearchBooksRaw(opts ...SearchBooksParamsFunc) ([]byte, error) {
	return c.SearchBooksRawContext(context.Background(), opts...)
}

//SearchBooks gets list of books. (struct data)
func (c *Client) SearchBooks(opts ...SearchBooksParamsFunc) ([]Book, error) {
	return c.SearchBooksContext(context.Background(), opts...)
}

//SearchBooksRawContext gets list of books with context.Context. (raw data)
func (c *Client) SearchBooksRawContext(ctx context.Context, opts ...SearchBooksParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	return c.get(ctx, c.makeSearchCommand(TargetBooks, params))
}

//SearchBooksContext gets list of books with context.Context. (struct data)
func (c *Client) SearchBooksContext(ctx context.Context, opts ...SearchBooksParamsFunc) ([]Book, error) {
	b, err := c.SearchBooksRawContext(ctx, opts...)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodeBooks(b)
}

//WithBookTitle returns function for setting Marketplace.
func WithBookTitle(title string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(title) > 0 {
			params.Set("title", title)
		}
	}
}

//WithBookAuthor returns function for setting Marketplace.
func WithBookAuthor(author string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(author) > 0 {
			params.Set("author", author)
		}
	}
}

//WithBookFields returns function for setting Marketplace.
func WithBookFields(fields string) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && len(fields) > 0 {
			params.Add("fields", fields)
		}
	}
}

//WithBookLimit returns function for setting Marketplace.
func WithBookLimit(limit int) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && limit > 0 {
			params.Set("limit", strconv.Itoa(limit))
		}
	}
}

//WithBookSkip returns function for setting Marketplace.
func WithBookSkip(skip int) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && skip > 0 {
			params.Set("skip", strconv.Itoa(skip))
		}
	}
}

//WithBookAfter returns function for setting Marketplace.
func WithBookAfter(after time.Time) SearchBooksParamsFunc {
	return func(params url.Values) {
		if params != nil && !after.IsZero() {
			params.Set("after", after.Format("2006-01-02"))
		}
	}
}

//SearchPersonsParamsFunc is self-referential function for functional options pattern.
type SearchPersonsParamsFunc func(url.Values)

//SearchPersonsRaw gets list of persons. (raw data)
func (c *Client) SearchPersonsRaw(opts ...SearchPersonsParamsFunc) ([]byte, error) {
	return c.SearchPersonsRawContext(context.Background(), opts...)
}

//SearchPersons gets list of persons. (struct data)
func (c *Client) SearchPersons(opts ...SearchPersonsParamsFunc) ([]Person, error) {
	return c.SearchPersonsContext(context.Background(), opts...)
}

//SearchPersonsRawContext gets list of persons with context.Context. (raw data)
func (c *Client) SearchPersonsRawContext(ctx context.Context, opts ...SearchPersonsParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	return c.get(ctx, c.makeSearchCommand(TargetPersons, params))
}

//SearchPersonsContext gets list of persons with context.Context. (struct data)
func (c *Client) SearchPersonsContext(ctx context.Context, opts ...SearchPersonsParamsFunc) ([]Person, error) {
	b, err := c.SearchPersonsRawContext(ctx, opts...)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodePersons(b)
}

//WithPersonName returns function for setting Marketplace.
func WithPersonName(name string) SearchPersonsParamsFunc {
	return func(params url.Values) {
		if params != nil && len(name) > 0 {
			params.Set("name", name)
		}
	}
}

//SearchPersonsParamsFunc is self-referential function for functional options pattern.
type SearchWorkersParamsFunc func(url.Values)

//SearchWorkersRaw gets list of workers (raw data)
func (c *Client) SearchWorkersRaw(opts ...SearchWorkersParamsFunc) ([]byte, error) {
	return c.SearchWorkersRawContext(context.Background(), opts...)
}

//SearchWorkers gets list of workers (struct data)
func (c *Client) SearchWorkers(opts ...SearchWorkersParamsFunc) ([]Worker, error) {
	return c.SearchWorkersContext(context.Background(), opts...)
}

//SearchWorkersRawContext gets list of workers with context.Context. (raw data)
func (c *Client) SearchWorkersRawContext(ctx context.Context, opts ...SearchWorkersParamsFunc) ([]byte, error) {
	params := url.Values{}
	for _, opt := range opts {
		opt(params)
	}
	return c.get(ctx, c.makeSearchCommand(TargetWorkers, params))
}

//SearchWorkersContext gets list of workers with context.Context. (struct data)
func (c *Client) SearchWorkersContext(ctx context.Context, opts ...SearchWorkersParamsFunc) ([]Worker, error) {
	b, err := c.SearchWorkersRawContext(ctx, opts...)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodeWorkers(b)
}

//WithWorkerName returns function for setting Marketplace.
func WithWorkerName(name string) SearchWorkersParamsFunc {
	return func(params url.Values) {
		if params != nil && len(name) > 0 {
			params.Set("name", name)
		}
	}
}

//LookupBookRaw gets book data. (raw data)
func (c *Client) LookupBookRaw(id int) ([]byte, error) {
	return c.LookupBookRawContext(context.Background(), id)
}

//LookupBook gets books data. (struct data)
func (c *Client) LookupBook(id int) (*Book, error) {
	return c.LookupBookContext(context.Background(), id)
}

//LookupBookCardRaw gets book card info (HTML page data)
func (c *Client) LookupBookCardRaw(id int) ([]byte, error) {
	return c.LookupBookCardRawContext(context.Background(), id)
}

//LookupBookContentRaw gets book content (plain or HTML formatted text data)
func (c *Client) LookupBookContentRaw(id int, f Format) ([]byte, error) {
	return c.LookupBookContentRawContext(context.Background(), id, f)
}

//LookupBookRawContext gets book data with context.Context. (raw data)
func (c *Client) LookupBookRawContext(ctx context.Context, id int) ([]byte, error) {
	return c.get(ctx, c.makeLookupCommand(TargetBooks, id))
}

//LookupBookContext gets books data with context.Context. (struct data)
func (c *Client) LookupBookContext(ctx context.Context, id int) (*Book, error) {
	b, err := c.LookupBookRawContext(ctx, id)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodeBook(b)
}

//LookupBookCardRawContext gets book card info with context.Context. (HTML page data)
func (c *Client) LookupBookCardRawContext(ctx context.Context, id int) ([]byte, error) {
	return c.get(ctx, c.makeCardCommand(id))
}

//LookupBookContentRawContext gets book content with context.Context. (plain or HTML formatted text data)
func (c *Client) LookupBookContentRawContext(ctx context.Context, id int, f Format) ([]byte, error) {
	return c.get(ctx, c.makeContentCommand(id, f))
}

//LookupPersonRaw gets person data. (raw data)
func (c *Client) LookupPersonRaw(id int) ([]byte, error) {
	return c.LookupPersonRawContext(context.Background(), id)
}

//LookupPerson gets person data. (struct data)
func (c *Client) LookupPerson(id int) (*Person, error) {
	return c.LookupPersonContext(context.Background(), id)
}

//LookupPersonRawContext gets person data with context.Context. (raw data)
func (c *Client) LookupPersonRawContext(ctx context.Context, id int) ([]byte, error) {
	return c.get(ctx, c.makeLookupCommand(TargetPersons, id))
}

//LookupPersonContext gets person data with context.Context. (struct data)
func (c *Client) LookupPersonContext(ctx context.Context, id int) (*Person, error) {
	b, err := c.LookupPersonRawContext(ctx, id)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodePerson(b)
}

//LookupWorkerRaw gets worker data. (raw data)
func (c *Client) LookupWorkerRaw(id int) ([]byte, error) {
	return c.LookupWorkerRawContext(context.Background(), id)
}

//LookupWorker gets worker data. (struct data)
func (c *Client) LookupWorker(id int) (*Worker, error) {
	return c.LookupWorkerContext(context.Background(), id)
}

//LookupWorkerRawContext gets worker data with context.Context. (raw data)
func (c *Client) LookupWorkerRawContext(ctx context.Context, id int) ([]byte, error) {
	return c.get(ctx, c.makeLookupCommand(TargetWorkers, id))
}

//LookupWorkerContext gets worker data with context.Context. (struct data)
func (c *Client) LookupWorkerContext(ctx context.Context, id int) (*Worker, error) {
	b, err := c.LookupWorkerRawContext(ctx, id)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodeWorker(b)
}

//RankingRaw gets ranking data (raw data)
func (c *Client) RankingRaw(tm time.Time) ([]byte, error) {
	return c.RankingRawContext(context.Background(), tm)
}

//Ranking gets ranking data (struct data)
func (c *Client) Ranking(tm time.Time) (Ranking, error) {
	return c.RankingContext(context.Background(), tm)
}

//RankingRawContext gets ranking data (raw data)
func (c *Client) RankingRawContext(ctx context.Context, tm time.Time) ([]byte, error) {
	return c.get(ctx, c.makeRankingCommand(tm))
}

//RankingContext gets ranking data (struct data)
func (c *Client) RankingContext(ctx context.Context, tm time.Time) (Ranking, error) {
	b, err := c.RankingRawContext(ctx, tm)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return DecodeRanking(b)
}

func (c *Client) makeSearchCommand(t Target, v url.Values) *url.URL {
	var u *url.URL
	if c == nil {
		u = New().URL()
	} else {
		u = c.server.URL()
	}
	u.Path = fmt.Sprintf("/%v/%v", c.apiDir(), t)
	u.RawQuery = v.Encode()
	return u
}

func (c *Client) makeLookupCommand(t Target, id int) *url.URL {
	var u *url.URL
	if c == nil {
		u = New().URL()
	} else {
		u = c.server.URL()
	}
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
	var u *url.URL
	if c == nil {
		u = New().URL()
	} else {
		u = c.server.URL()
	}
	u.Path = fmt.Sprintf("/%v/%v/%v/%v", c.apiDir(), TargetRanking, "xhtml", tm.Format("2006/01"))
	return u
}

func (c *Client) apiDir() string {
	return defaultAPIDir
}

func (c *Client) get(ctx context.Context, u *url.URL) ([]byte, error) {
	if c == nil {
		return nil, errs.Wrap(ErrNullPointer)
	}
	resp, err := c.client.Get(u, fetch.WithContext(ctx))
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("url", u.String()))
	}
	return resp.DumpBodyAndClose()
}

/* Copyright 2019-2021 Spiegel
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
