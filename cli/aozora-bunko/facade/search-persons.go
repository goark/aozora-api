package facade

import (
	"github.com/goark/aozora-api"
	"github.com/goark/aozora-api/cli/aozora-bunko/ecode"
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newSearchPersonsCmd returns cobra.Command instance for show sub-command
func newSearchPersonsCmd(ui *rwi.RWI) *cobra.Command {
	searchPersonsCmd := &cobra.Command{
		Use:   "persons",
		Short: "Search for Aozora-bunko authors data",
		Long:  "Search for Aozora-bunko authors data",
		RunE: func(cmd *cobra.Command, args []string) error {
			//title option
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return errs.New("--name", errs.WithCause(err))
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().SearchPersonsRaw(
					aozora.WithPersonName(name),
				)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			persons, err := aozora.DefaultClient().SearchPersons(
				aozora.WithPersonName(name),
			)
			if err != nil {
				return debugPrint(ui, err)
			}
			if len(persons) == 0 {
				return debugPrint(ui, errs.New("error in search persons", errs.WithCause(ecode.ErrNoData)))
			}
			b, err := aozora.EncodePersons(persons)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}
	//options
	searchPersonsCmd.Flags().StringP("name", "n", "", "Search option: author name")

	return searchPersonsCmd
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
