package facade

import (
	"github.com/goark/aozora-api/cli/aozora-bunko/ecode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newLookupCmd returns cobra.Command instance for show sub-command
func newLookupCmd(ui *rwi.RWI) *cobra.Command {
	lookupCmd := &cobra.Command{
		Use:   "lookup",
		Short: "Lookup Aozora-bunko data",
		Long:  "Lookup Aozora-bunko data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, ecode.ErrNoCommand)
		},
	}
	lookupCmd.AddCommand(newLookupBookCmd(ui))
	lookupCmd.AddCommand(newLookupPersonCmd(ui))
	lookupCmd.AddCommand(newLookupWorkerCmd(ui))

	return lookupCmd
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
