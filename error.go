package aozora

import "errors"

var (
	ErrNullPointer = errors.New("Null reference instance")
	ErrHTTPStatus  = errors.New("Bad HTTP status")
	ErrNoData      = errors.New("No response data")
)

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
