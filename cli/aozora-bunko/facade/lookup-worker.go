package facade

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/aozora-api"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newLookupPersonCmd returns cobra.Command instance for show sub-command
func newLookupWorkerCmd(ui *rwi.RWI) *cobra.Command {
	lookupWorkerCmd := &cobra.Command{
		Use:   "worker [flags] <worker id>",
		Short: "Lookup worker data in Aozora-bunko",
		Long:  "Lookup worker data in Aozora-bunko",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errs.Wrap(os.ErrInvalid, "worker id")
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().LookupWorkerRaw(args[0])
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			worker, err := aozora.DefaultClient().LookupWorker(args[0])
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := aozora.EncodeWorker(worker)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}

	return lookupWorkerCmd
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
