package aozora

import (
	"context"
	"net/http"
	"net/url"

	"github.com/spiegel-im-spiegel/fetch"
)

const (
	defaultScheme = "http"
	defaultHost   = "www.aozorahack.net"
)

//Server is informations of Aozora API
type Server struct {
	scheme string
	name   string //Aozora API server name
}

//ServerOptFunc is self-referential function for functional options pattern
type ServerOptFunc func(*Server)

//New returns new Server instance
func New(opts ...ServerOptFunc) *Server {
	server := &Server{scheme: defaultScheme, name: defaultHost}
	for _, opt := range opts {
		opt(server)
	}
	return server
}

//WithScheme returns function for setting scheme
func WithScheme(scheme string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.scheme = scheme
		}
	}
}

//WithServerName returns function for setting hostname
func WithServerName(host string) ServerOptFunc {
	return func(s *Server) {
		if s != nil {
			s.name = host
		}
	}
}

//URL returns url.URL instance
func (s *Server) URL() *url.URL {
	if s == nil {
		s = New()
	}
	return &url.URL{Scheme: s.scheme, Host: s.name}
}

//ClientOptFunc is self-referential function for functional options pattern
type ClientOptFunc func(*Client)

//CreateClient returns new Client instance
func (s *Server) CreateClient(opts ...ClientOptFunc) *Client {
	if s == nil {
		s = New()
	}
	cli := &Client{server: s}
	for _, opt := range opts {
		opt(cli)
	}
	if cli.client == nil {
		cli.client = fetch.New()
	}
	return cli
}

//WithContext is dummy function. Because this function is deprecated.
func WithContext(ctx context.Context) ClientOptFunc {
	return func(c *Client) {}
}

//WithHttpClient returns function for setting http.Client
func WithHttpClient(client *http.Client) ClientOptFunc {
	return func(c *Client) {
		if c != nil {
			c.client = fetch.New(fetch.WithHTTPClient(client))
		}
	}
}

//DefaultClient returns new Client instance with default setting
func DefaultClient() *Client {
	return New().CreateClient()
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
