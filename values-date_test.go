package aozora

import (
	"encoding/json"
	"fmt"
	"testing"
)

type ForTestStruct struct {
	DateTaken Date `json:"date_taken,omitempty"`
}

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		s   string
		str string
		jsn string
	}{
		{s: `{"date_taken": "2005-03-26T00:00:00+09:00"}`, str: "2005-03-26", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03-26"}`, str: "2005-03-26", jsn: `{"date_taken":"2005-03-26"}`},
		{s: `{"date_taken": "2005-03"}`, str: "2005-03-01", jsn: `{"date_taken":"2005-03-01"}`},
		{s: `{"date_taken": "2005"}`, str: "2005-01-01", jsn: `{"date_taken":"2005-01-01"}`},
		{s: `{"date_taken": ""}`, str: "0001-01-01", jsn: `{"date_taken":""}`},
		{s: `{"date_taken": null}`, str: "0001-01-01", jsn: `{"date_taken":""}`},
		{s: `{}`, str: "0001-01-01", jsn: `{"date_taken":""}`},
	}

	for _, tc := range testCases {
		tst := &ForTestStruct{}
		if err := json.Unmarshal([]byte(tc.s), tst); err != nil {
			t.Errorf("json.Unmarshal() is \"%v\", want nil.", err)
			continue
		}
		str := tst.DateTaken.String()
		if str != tc.str {
			t.Errorf("Date = \"%v\", want \"%v\".", str, tc.str)
		}
		b, err := json.Marshal(tst)
		if err != nil {
			t.Errorf("json.Marshal() is \"%v\", want nil.", err)
			continue
		}
		str = string(b)
		if str != tc.jsn {
			t.Errorf("ForTestStruct = \"%v\", want \"%v\".", str, tc.jsn)
		}
	}
}

func TestUnmarshalErr(t *testing.T) {
	data := `{"date_taken": "2005/03/26"}`
	tst := &ForTestStruct{}
	if err := json.Unmarshal([]byte(data), tst); err == nil {
		t.Error("Unmarshal() error = nil, not want nil.")
	} else {
		fmt.Printf("Info: %+v\n", err)
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
