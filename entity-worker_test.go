package aozora

import (
	"testing"
)

var inputWorkerJSON = `{"id":845,"name":"雪森"}`
var inputWorkersJSON = "[" + inputWorkerJSON + "]"
var respWorkerJSON = `{"id":845,"name":"雪森"}`
var respWorkersJSON = "[" + respWorkerJSON + "]"

func TestDecodeWorker(t *testing.T) {
	worker, err := DecodeWorker([]byte(inputWorkerJSON))
	if err != nil {
		t.Errorf("DecodeWorker() error = \"%v\", want nil.", err)
		return
	}
	str := worker.String()
	if str != respWorkerJSON {
		t.Errorf("EncodeWorker() = \"%v\", want \"%v\".", str, respWorkerJSON)
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
	if str != respWorkersJSON {
		t.Errorf("EncodeWorkers() = \"%v\", want \"%v\".", str, respWorkersJSON)
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
