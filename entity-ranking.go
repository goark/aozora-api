package aozora

import (
	"bytes"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

//Ranking is entity class of ranking info.
type Ranking []struct {
	BookID  int      `json:"book_id"`
	Access  int      `json:"access"`
	Title   string   `json:"title"`
	Authors []string `json:"authors"`
}

func (ranking Ranking) String() string {
	if b, err := EncodeRanking(ranking); err == nil {
		return string(b)
	}
	return ""
}

//DecodeRanking returns Ranking instance from byte buffer
func DecodeRanking(b []byte) (Ranking, error) {
	ranking := Ranking{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&ranking); err != nil {
		return ranking, errs.Wrap(err, "error in DecodeRanking() function")
	}
	return ranking, nil
}

//EncodeRanking returns bytes encoded from Worker instance
func EncodeRanking(ranking Ranking) ([]byte, error) {
	return json.Marshal(ranking)
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
