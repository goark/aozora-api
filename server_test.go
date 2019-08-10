package aozora

import "testing"

func TestServer(t *testing.T) {

	testCases := []struct {
		s    *Server
		host string
	}{
		{s: New(), host: DefaultHost},
		{s: New(WithServerName("foo")), host: "foo"},
	}

	for _, tc := range testCases {
		if tc.s.name != tc.host {
			t.Errorf("Server.name is \"%v\", want \"%v\"", tc.host, tc.host)
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
