package facade

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/aozora-api"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newLookupPersonCmd returns cobra.Command instance for show sub-command
func newLookupPersonCmd(ui *rwi.RWI) *cobra.Command {
	lookupPersonCmd := &cobra.Command{
		Use:   "person [flags] <person id>",
		Short: "Lookup author data in Aozora-bunko",
		Long:  "Lookup author data in Aozora-bunko",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errs.New("person id", errs.WithCause(os.ErrInvalid))
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errs.New("invalid person id", errs.WithCause(err))
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().LookupPersonRaw(id)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			person, err := aozora.DefaultClient().LookupPerson(id)
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := aozora.EncodePerson(person)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}

	return lookupPersonCmd
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
