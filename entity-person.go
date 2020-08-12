package aozora

import (
	"bytes"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

//Person is entity class of person info.
type Person struct {
	PersonID        int    `json:"person_id"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	LastNameYomi    string `json:"last_name_yomi"`
	FirstNameYomi   string `json:"first_name_yomi"`
	LastNameSort    string `json:"last_name_sort"`
	FirstNameSort   string `json:"first_name_sort"`
	LastNameRoman   string `json:"last_name_roman"`
	FirstNameRoman  string `json:"first_name_roman"`
	DateOfBirth     Date   `json:"date_of_birth"`
	DateOfDeath     Date   `json:"date_of_death"`
	AuthorCopyright bool   `json:"author_copyright"`
}

func (person *Person) String() string {
	if b, err := EncodePerson(person); err == nil {
		return string(b)
	}
	return ""
}

//DecodePerson returns Person instance from byte buffer
func DecodePerson(b []byte) (*Person, error) {
	person := Person{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&person); err != nil {
		return &person, errs.New("error in DecodePerson() function", errs.WithCause(err))
	}
	return &person, nil
}

//DecodePersons returns array of Person instance from byte buffer
func DecodePersons(b []byte) ([]Person, error) {
	persons := []Person{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&persons); err != nil {
		return persons, errs.New("error in DecodeBooks() function", errs.WithCause(err))
	}
	return persons, nil
}

//EncodePersons returns bytes encoded from Person instance
func EncodePerson(person *Person) ([]byte, error) {
	return json.Marshal(person)
}

//EncodePerson returns bytes encoded from list of Person
func EncodePersons(persons []Person) ([]byte, error) {
	return json.Marshal(persons)
}

/* Copyright 2019,2020 Spiegel
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
