package aozora

import (
	"testing"
)

var inputRankingJSON = `[{"book_id":773,"access":25583,"title":"こころ","authors":["夏目 漱石"]},{"book_id":51732,"access":798,"title":"古事記","authors":["太 安万侶","稗田 阿礼"]},{"book_id":2272,"access":322,"title":"水仙","authors":["太宰 治"]}]`
var respRankingJSON = inputRankingJSON

func TestDecodeRanking(t *testing.T) {
	worker, err := DecodeRanking([]byte(inputRankingJSON))
	if err != nil {
		t.Errorf("DecodeRanking() error = \"%v\", want nil.", err)
		return
	}
	str := worker.String()
	if str != respRankingJSON {
		t.Errorf("EncodeRanking() = \"%v\", want \"%v\".", str, respRankingJSON)
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
