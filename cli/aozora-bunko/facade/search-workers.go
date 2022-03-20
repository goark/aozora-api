package facade

import (
	"github.com/goark/aozora-api"
	"github.com/goark/aozora-api/cli/aozora-bunko/ecode"
	"github.com/goark/errs"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newSearchWorkersCmd returns cobra.Command instance for show sub-command
func newSearchWorkersCmd(ui *rwi.RWI) *cobra.Command {
	searchWorkersCmd := &cobra.Command{
		Use:   "workers",
		Short: "Search for Aozora-bunko workers data",
		Long:  "Search for Aozora-bunko workers data",
		RunE: func(cmd *cobra.Command, args []string) error {
			//title option
			name, err := cmd.Flags().GetString("name")
			if err != nil {
				return errs.New("--name", errs.WithCause(err))
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().SearchWorkersRaw(
					aozora.WithWorkerName(name),
				)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			workers, err := aozora.DefaultClient().SearchWorkers(
				aozora.WithWorkerName(name),
			)
			if err != nil {
				return debugPrint(ui, err)
			}
			if len(workers) == 0 {
				return debugPrint(ui, errs.New("error in search workers", errs.WithCause(ecode.ErrNoData)))
			}
			b, err := aozora.EncodeWorkers(workers)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}
	//options
	searchWorkersCmd.Flags().StringP("name", "n", "", "Search option: worker name")

	return searchWorkersCmd
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
