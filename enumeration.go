package aozora

//Target is enumrate class of data type
type Target int

const (
	TargetBooks Target = iota + 1
	TargetPersons
	TargetWorkers
	TargetRanking
)

var targetMap = map[Target]string{
	TargetBooks:   "books",
	TargetPersons: "persons",
	TargetWorkers: "workers",
	TargetRanking: "ranking",
}

func (t Target) String() string {
	if s, ok := targetMap[t]; ok {
		return s
	}
	return "unknown"
}

//Format is enumrate class of content type
type Format int

const (
	Text Format = iota + 1
	HTML
)

var formatMap = map[Format]string{
	Text: "txt",
	HTML: "html",
}

func (f Format) String() string {
	if s, ok := formatMap[f]; ok {
		return s
	}
	return formatMap[Text]
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
