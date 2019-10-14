package aozora

import "testing"

func TestServer(t *testing.T) {
	testCases := []struct {
		s   *Server
		str string
	}{
		{s: (*Server)(nil), str: defaultScheme + "://" + defaultHost},
		{s: New(), str: defaultScheme + "://" + defaultHost},
		{s: New(WithScheme("foo"), WithServerName("bar")), str: "foo://bar"},
	}

	for _, tc := range testCases {
		str := tc.s.URL().String()
		if str != tc.str {
			t.Errorf("Server.name is \"%v\", want \"%v\"", str, tc.str)
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
