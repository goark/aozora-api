package facade

import (
	"fmt"
	"os"
	"strconv"
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
		if strings.EqualFold(s, v) {
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
				return errs.New("--content-type", errs.WithCause(err))
			}
			content := TypeUnknown
			if len(c) > 0 {
				content = NewContent(c)
				if content == TypeUnknown {
					return errs.New(fmt.Sprintf("content-type %v", c), errs.WithCause(os.ErrInvalid))
				}
			}
			if len(args) == 0 {
				return errs.New("book id", errs.WithCause(os.ErrInvalid))
			}
			id, err := strconv.Atoi(args[0])
			if err != nil {
				return errs.New("invalid book id", errs.WithCause(err))
			}

			client := aozora.DefaultClient()
			switch content {
			case TypeText:
				resp, err := client.LookupBookContentRaw(id, aozora.Text)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			case TypeHTML:
				resp, err := client.LookupBookContentRaw(id, aozora.HTML)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			case TypeCard:
				resp, err := client.LookupBookCardRaw(id)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}
			if rawFlag {
				resp, err := aozora.DefaultClient().LookupBookRaw(id)
				if err != nil {
					return debugPrint(ui, err)
				}
				return debugPrint(ui, ui.OutputBytes(resp))
			}

			book, err := aozora.DefaultClient().LookupBook(id)
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
