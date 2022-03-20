package facade

import (
	"github.com/goark/aozora-api/cli/aozora-bunko/ecode"
	"github.com/goark/gocli/rwi"
	"github.com/spf13/cobra"
)

//newSearchCmd returns cobra.Command instance for show sub-command
func newSearchCmd(ui *rwi.RWI) *cobra.Command {
	searchCmd := &cobra.Command{
		Use:   "search",
		Short: "Search for Aozora-bunko data",
		Long:  "Search for Aozora-bunko data",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, ecode.ErrNoCommand)
		},
	}
	searchCmd.AddCommand(newSearchBooksCmd(ui))
	searchCmd.AddCommand(newSearchPersonsCmd(ui))
	searchCmd.AddCommand(newSearchWorkersCmd(ui))

	return searchCmd
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
