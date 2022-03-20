package facade

import (
	"github.com/goark/aozora-api"
	"github.com/goark/aozora-api/cli/aozora-bunko/ecode"
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newSearchBooksCmd returns cobra.Command instance for show sub-command
func newSearchBooksCmd(ui *rwi.RWI) *cobra.Command {
	searchBooksCmd := &cobra.Command{
		Use:   "books",
		Short: "Search for Aozora-bunko books data",
		Long:  "Search for Aozora-bunko books data",
		RunE: func(cmd *cobra.Command, args []string) error {
			//title option
			title, err := cmd.Flags().GetString("title")
			if err != nil {
				return debugPrint(ui, errs.New("--title", errs.WithCause(err)))
			}
			//author option
			author, err := cmd.Flags().GetString("author")
			if err != nil {
				return debugPrint(ui, errs.New("--author", errs.WithCause(err)))
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().SearchBooksRaw(
					aozora.WithBookTitle(title),
					aozora.WithBookAuthor(author),
				)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			books, err := aozora.DefaultClient().SearchBooks(
				aozora.WithBookTitle(title),
				aozora.WithBookAuthor(author),
			)
			if err != nil {
				return debugPrint(ui, err)
			}
			if len(books) == 0 {
				return debugPrint(ui, errs.New("error in search books", errs.WithCause(ecode.ErrNoData)))
			}
			b, err := aozora.EncodeBooks(books)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}
	//options
	searchBooksCmd.Flags().StringP("title", "t", "", "Search option: book title")
	searchBooksCmd.Flags().StringP("author", "a", "", "Search option: author name")

	return searchBooksCmd
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
