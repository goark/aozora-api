package aozora

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	APIVersion = "api/v0.1"
)

//Client is http.Client for Aozora API Server
type Client struct {
	server *Server
	client *http.Client
}

//MakeSearchCommand returns URI for search command
func (c *Client) MakeSearchCommand(t Target, v url.Values) *url.URL {
	return &url.URL{
		Scheme:   c.server.scheme,
		Host:     c.server.name,
		Path:     fmt.Sprintf("/%v/%v", APIVersion, t),
		RawQuery: v.Encode(),
	}
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeLookupCommand(t Target, id string) *url.URL {
	return &url.URL{
		Scheme: c.server.scheme,
		Host:   c.server.name,
		Path:   fmt.Sprintf("/%v/%v/%v", APIVersion, t, id),
	}
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeCardCommand(id string) *url.URL {
	u := c.MakeLookupCommand(Books, id)
	u.Path = u.Path + "/card"
	return u
}

//MakeLookupCommand returns URI for lookup command
func (c *Client) MakeContentCommand(id string, f Format) *url.URL {
	u := c.MakeLookupCommand(Books, id)
	u.Path = u.Path + "/content"
	u.RawQuery = (url.Values{"format": {f.String()}}).Encode()
	return u
}

//MakeLookupCommand returns URI for lookup ranking info command
func (c *Client) MakeRankingCommand(tm time.Time) *url.URL {
	return &url.URL{
		Scheme: c.server.scheme,
		Host:   c.server.name,
		Path:   fmt.Sprintf("/%v/%v/%v/%v", APIVersion, Ranking, "xhtml", tm.Format("2006/01")),
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
