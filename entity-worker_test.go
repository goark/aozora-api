package aozora

import (
	"testing"
)

var inputWorkerJSON = `{"id":845,"name":"雪森"}`
var inputWorkersJSON = "[" + inputWorkerJSON + "]"
var responseWorkerJSON = `{"id":845,"name":"雪森"}`
var responseWorkersJSON = "[" + responseWorkerJSON + "]"

func TestDecodeWorker(t *testing.T) {
	worker, err := DecodeWorker([]byte(inputWorkerJSON))
	if err != nil {
		t.Errorf("DecodeWorker() error = \"%v\", want nil.", err)
		return
	}
	str := worker.String()
	if str != responseWorkerJSON {
		t.Errorf("EncodeWorker() = \"%v\", want \"%v\".", str, responseWorkerJSON)
	}
}

func TestDecodeWorkers(t *testing.T) {
	workers, err := DecodeWorkers([]byte(inputWorkersJSON))
	if err != nil {
		t.Errorf("DecodeWorkers() error = \"%v\", want nil.", err)
		return
	}
	b, err := EncodeWorkers(workers)
	if err != nil {
		t.Errorf("EncodeWorkers() error = \"%v\", want nil.", err)
		return
	}
	str := string(b)
	if str != responseWorkersJSON {
		t.Errorf("EncodeWorkers() = \"%v\", want \"%v\".", str, responseWorkersJSON)
	}
}

/* Copyright 2019-2022 Spiegel
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
