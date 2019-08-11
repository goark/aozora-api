package aozora

import (
	"bytes"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

//Worker is entity class of worker info.
type Worker struct {
	WorkerID int    `json:"id"`
	Name     string `json:"name"`
}

func (worker *Worker) String() string {
	if b, err := EncodeWorker(worker); err == nil {
		return string(b)
	}
	return ""
}

//DecodeWorker returns Worker instance from byte buffer
func DecodeWorker(b []byte) (*Worker, error) {
	worker := Worker{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&worker); err != nil {
		return &worker, errs.Wrap(err, "error in DecodeWorker() function")
	}
	return &worker, nil
}

//DecodeWorkers returns array of Worker instance from byte buffer
func DecodeWorkers(b []byte) ([]Worker, error) {
	workers := []Worker{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&workers); err != nil {
		return workers, errs.Wrap(err, "error in DecodeWorkers() function")
	}
	return workers, nil
}

//EncodeWorkers returns bytes encoded from Worker instance
func EncodeWorker(worker *Worker) ([]byte, error) {
	return json.Marshal(worker)
}

//EncodeWorker returns bytes encoded from list of Worker
func EncodeWorkers(workers []Worker) ([]byte, error) {
	return json.Marshal(workers)
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
