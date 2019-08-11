package aozora

import (
	"testing"
)

var inputBookJSON = `{"book_id":55881,"title":"緋のエチュード","title_yomi":"ひのエチュード","title_sort":"ひのえちゆうと","subtitle":"","subtitle_yomi":"","original_title":"A STUDY IN SCARLET","first_appearance":"","ndc_code":"NDC K933","font_kana_type":"新字新仮名","copyright":true,"release_date":"2013-02-02T00:00:00.000Z","last_modified":"2014-09-16T00:00:00.000Z","card_url":"https://www.aozora.gr.jp/cards/000009/card55881.html","base_book_1":"","base_book_1_publisher":"","base_book_1_1st_edition":"","base_book_1_edition_input":"","base_book_1_edition_proofing":"","base_book_1_parent":"","base_book_1_parent_publisher":"","base_book_1_parent_1st_edition":"","base_book_2":"","base_book_2_publisher":"","base_book_2_1st_edition":"","base_book_2_edition_input":"","base_book_2_edition_proofing":"","base_book_2_parent":"","base_book_2_parent_publisher":"","base_book_2_parent_1st_edition":"","input":"大久保ゆう","proofing":"","text_url":"https://www.aozora.gr.jp/cards/000009/files/55881_ruby_50043.zip","text_last_modified":"2014-04-03T00:00:00.000Z","text_encoding":"ShiftJIS","text_charset":"JIS X 0208","text_updated":1,"html_url":"https://www.aozora.gr.jp/cards/000009/files/55881_50044.html","html_last_modified":"2014-04-03T00:00:00.000Z","html_encoding":"ShiftJIS","html_charset":"JIS X 0208","html_updated":1,"translators":[{"person_id":10,"last_name":"大久保","first_name":"ゆう"}],"authors":[{"person_id":9,"last_name":"ドイル","first_name":"アーサー・コナン"}]}`
var inputBooksJSON = "[" + inputBookJSON + "]"
var respBookJSON = `{"book_id":55881,"title":"緋のエチュード","title_yomi":"ひのエチュード","title_sort":"ひのえちゆうと","subtitle":"","subtitle_yomi":"","original_title":"A STUDY IN SCARLET","first_appearance":"","ndc_code":"NDC K933","font_kana_type":"新字新仮名","copyright":true,"release_date":"2013-02-02","last_modified":"2014-09-16","card_url":"https://www.aozora.gr.jp/cards/000009/card55881.html","base_book_1":"","base_book_1_publisher":"","base_book_1_1st_edition":"","base_book_1_edition_input":"","base_book_1_edition_proofing":"","base_book_1_parent":"","base_book_1_parent_publisher":"","base_book_1_parent_1st_edition":"","base_book_2":"","base_book_2_publisher":"","base_book_2_1st_edition":"","base_book_2_edition_input":"","base_book_2_edition_proofing":"","base_book_2_parent":"","base_book_2_parent_publisher":"","base_book_2_parent_1st_edition":"","input":"大久保ゆう","proofing":"","text_url":"https://www.aozora.gr.jp/cards/000009/files/55881_ruby_50043.zip","text_last_modified":"2014-04-03","text_encoding":"ShiftJIS","text_charset":"JIS X 0208","text_updated":1,"html_url":"https://www.aozora.gr.jp/cards/000009/files/55881_50044.html","html_last_modified":"2014-04-03","html_encoding":"ShiftJIS","html_charset":"JIS X 0208","html_updated":1,"translators":[{"person_id":10,"last_name":"大久保","first_name":"ゆう"}],"authors":[{"person_id":9,"last_name":"ドイル","first_name":"アーサー・コナン"}]}`
var respBooksJSON = "[" + respBookJSON + "]"

func TestDecodeBook(t *testing.T) {
	book, err := DecodeBook([]byte(inputBookJSON))
	if err != nil {
		t.Errorf("DecodeBook() error = \"%v\", want nil.", err)
		return
	}
	str := book.String()
	if str != respBookJSON {
		t.Errorf("EncodeBook() = \"%v\", want \"%v\".", str, respBookJSON)
	}
}

func TestDecodeBooks(t *testing.T) {
	books, err := DecodeBooks([]byte(inputBooksJSON))
	if err != nil {
		t.Errorf("DecodeBooks() error = \"%v\", want nil.", err)
		return
	}
	b, err := EncodeBooks(books)
	if err != nil {
		t.Errorf("EncodeBooks() error = \"%v\", want nil.", err)
		return
	}
	str := string(b)
	if str != respBooksJSON {
		t.Errorf("EncodeBooks() = \"%v\", want \"%v\".", str, respBooksJSON)
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
