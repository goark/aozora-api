package facade

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/aozora-api"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

type ContentType int

const (
	TypeUnknown ContentType = iota
	TypeCard
	TypeText
	TypeHTML
)

var contentMap = map[ContentType]string{
	TypeUnknown: "unknown",
	TypeCard:    "card",
	TypeText:    "text",
	TypeHTML:    "html",
}

func NewContent(s string) ContentType {
	for k, v := range contentMap {
		if strings.ToLower(s) == strings.ToLower(v) {
			return k
		}
	}
	return TypeUnknown
}

func (c ContentType) String() string {
	if s, ok := contentMap[c]; ok {
		return s
	}
	return contentMap[TypeUnknown]
}

//newLookupBookCmd returns cobra.Command instance for show sub-command
func newLookupBookCmd(ui *rwi.RWI) *cobra.Command {
	lookupBookCmd := &cobra.Command{
		Use:   "book [flags] <book id>",
		Short: "Lookup book data in Aozora-bunko",
		Long:  "Lookup book data in Aozora-bunko",
		RunE: func(cmd *cobra.Command, args []string) error {
			//content-type option
			c, err := cmd.Flags().GetString("content-type")
			if err != nil {
				return errs.Wrap(err, "--content-type")
			}
			content := TypeUnknown
			if len(c) > 0 {
				content = NewContent(c)
				if content == TypeUnknown {
					return errs.Wrapf(os.ErrInvalid, "content-type %v", c)
				}
			}
			if len(args) == 0 {
				return errs.Wrap(os.ErrInvalid, "book id")
			}

			client := aozora.DefaultClient()
			switch content {
			case TypeText:
				resp, err := client.LookupBookContentRaw(args[0], aozora.Text)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			case TypeHTML:
				resp, err := client.LookupBookContentRaw(args[0], aozora.HTML)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			case TypeCard:
				resp, err := client.LookupBookCardRaw(args[0])
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}
			if rawFlag {
				resp, err := aozora.DefaultClient().LookupBookRaw(args[0])
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			book, err := aozora.DefaultClient().LookupBook(args[0])
			if err != nil {
				return debugPrint(ui, err)
			}
			b, err := aozora.EncodeBook(book)
			if err != nil {
				return debugPrint(ui, err)
			}
			return debugPrint(ui, ui.OutputBytes(b))
		},
	}
	//options
	lookupBookCmd.Flags().StringP("content-type", "c", "", "Content type (card/text/html)")

	return lookupBookCmd
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
