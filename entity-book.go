package aozora

import (
	"bytes"
	"encoding/json"

	"github.com/spiegel-im-spiegel/errs"
)

//Author is entity class of author and translator info.
type Author struct {
	PersonID  int    `json:"person_id"`
	LastName  string `json:"last_name"`
	FirstName string `json:"first_name"`
}

//Book is entity class of book info.
type Book struct {
	BookID                      int      `json:"book_id"`
	Title                       string   `json:"title"`
	TitleYomi                   string   `json:"title_yomi"`
	TitleSort                   string   `json:"title_sort"`
	Subtitle                    string   `json:"subtitle"`
	SubtitleYomi                string   `json:"subtitle_yomi"`
	OriginalTitle               string   `json:"original_title"`
	FirstAppearance             string   `json:"first_appearance"`
	NDCCode                     string   `json:"ndc_code"`
	FontKanaType                string   `json:"font_kana_type"`
	Copyright                   bool     `json:"copyright"`
	ReleaseDate                 Date     `json:"release_date"`
	LastModified                Date     `json:"last_modified"`
	CardURL                     string   `json:"card_url"`
	BaseBook1                   string   `json:"base_book_1"`
	BaseBookPublisher1          string   `json:"base_book_1_publisher"`
	BaseBookFirstEdition1       string   `json:"base_book_1_1st_edition"`
	BaseBookEditionInput1       string   `json:"base_book_1_edition_input"`
	BaseBookEditionProofing1    string   `json:"base_book_1_edition_proofing"`
	BaseBookParent1             string   `json:"base_book_1_parent"`
	BaseBookParentPublisher1    string   `json:"base_book_1_parent_publisher"`
	BaseBookParentFirstEdition1 string   `json:"base_book_1_parent_1st_edition"`
	BaseBook2                   string   `json:"base_book_2"`
	BaseBookPublisher2          string   `json:"base_book_2_publisher"`
	BaseBookFirstEdition2       string   `json:"base_book_2_1st_edition"`
	BaseBookEditionInput2       string   `json:"base_book_2_edition_input"`
	BaseBookEditionProofing2    string   `json:"base_book_2_edition_proofing"`
	BaseBookParent2             string   `json:"base_book_2_parent"`
	BaseBookParentPublisher2    string   `json:"base_book_2_parent_publisher"`
	BaseBookParentFirstEdition2 string   `json:"base_book_2_parent_1st_edition"`
	Input                       string   `json:"input"`
	Proofing                    string   `json:"proofing"`
	TextURL                     string   `json:"text_url"`
	TextLastModified            Date     `json:"text_last_modified"`
	TextEncoding                string   `json:"text_encoding"`
	TextCharset                 string   `json:"text_charset"`
	TextUpdated                 int      `json:"text_updated"`
	HTMLURL                     string   `json:"html_url"`
	HTMLLastModified            Date     `json:"html_last_modified"`
	HTMLEncoding                string   `json:"html_encoding"`
	HTMLCharset                 string   `json:"html_charset"`
	HTMLUpdated                 int      `json:"html_updated"`
	Translators                 []Author `json:"translators"`
	Authors                     []Author `json:"authors"`
}

func (book *Book) String() string {
	if b, err := EncodeBook(book); err == nil {
		return string(b)
	}
	return ""
}

//DecodeBook returns Book instance from byte buffer
func DecodeBook(b []byte) (*Book, error) {
	book := Book{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&book); err != nil {
		return &book, errs.Wrap(err, "error in DecodeBook() function")
	}
	return &book, nil
}

//DecodeBooks returns array of Book instance from byte buffer
func DecodeBooks(b []byte) ([]Book, error) {
	books := []Book{}
	if err := json.NewDecoder(bytes.NewReader(b)).Decode(&books); err != nil {
		return books, errs.Wrap(err, "error in DecodeBooks() function")
	}
	return books, nil
}

//EncodeBook returns bytes encoded from Book instance
func EncodeBook(book *Book) ([]byte, error) {
	return json.Marshal(book)
}

//EncodeBooks returns bytes encoded from list of Book
func EncodeBooks(books []Book) ([]byte, error) {
	return json.Marshal(books)
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
