package facade

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/aozora-api"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

//newRankingCmd returns cobra.Command instance for show sub-command
func newRankingCmd(ui *rwi.RWI) *cobra.Command {
	rankingCmd := &cobra.Command{
		Use:   "ranking [flags] YYYY-MM",
		Short: "Lookup ranking data in Aozora-bunko",
		Long:  "Lookup ranking data in Aozora-bunko",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errs.Wrap(os.ErrInvalid, "year-month")
			}
			tm, err := time.Parse("2006-01", args[0])
			if err != nil {
				return errs.Wrap(err, fmt.Sprintf("argument: %v", args[0]))
			}

			if rawFlag {
				resp, err := aozora.DefaultClient().RankingRaw(tm)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			ranking, err := aozora.DefaultClient().Ranking(tm)
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := aozora.EncodeRanking(ranking)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}

	return rankingCmd
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
