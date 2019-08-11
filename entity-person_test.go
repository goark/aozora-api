package aozora

import (
	"testing"
)

var inputPersonJSON = `{"person_id":10,"last_name":"大久保","first_name":"ゆう","last_name_yomi":"おおくぼ","first_name_yomi":"ゆう","last_name_sort":"おおくほ","first_name_sort":"ゆう","last_name_roman":"Okubo","first_name_roman":"Yu","date_of_birth":"1982","date_of_death":"","author_copyright":true}`
var inputPersonsJSON = "[" + inputPersonJSON + "]"
var respPersonJSON = `{"person_id":10,"last_name":"大久保","first_name":"ゆう","last_name_yomi":"おおくぼ","first_name_yomi":"ゆう","last_name_sort":"おおくほ","first_name_sort":"ゆう","last_name_roman":"Okubo","first_name_roman":"Yu","date_of_birth":"1982-01-01","date_of_death":"","author_copyright":true}`
var respPersonsJSON = "[" + respPersonJSON + "]"

func TestDecodePerson(t *testing.T) {
	person, err := DecodePerson([]byte(inputPersonJSON))
	if err != nil {
		t.Errorf("DecodePerson() error = \"%v\", want nil.", err)
		return
	}
	str := person.String()
	if str != respPersonJSON {
		t.Errorf("EncodePerson() = \"%v\", want \"%v\".", str, respPersonJSON)
	}
}

func TestDecodePersons(t *testing.T) {
	persons, err := DecodePersons([]byte(inputPersonsJSON))
	if err != nil {
		t.Errorf("DecodePersons() error = \"%v\", want nil.", err)
		return
	}
	b, err := EncodePersons(persons)
	if err != nil {
		t.Errorf("EncodePersons() error = \"%v\", want nil.", err)
		return
	}
	str := string(b)
	if str != respPersonsJSON {
		t.Errorf("EncodePersons() = \"%v\", want \"%v\".", str, respPersonsJSON)
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
