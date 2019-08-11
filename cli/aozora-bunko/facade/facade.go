package facade

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spiegel-im-spiegel/aozora-api/cli/aozora-bunko/ecode"
	"github.com/spiegel-im-spiegel/errs"
	"github.com/spiegel-im-spiegel/gocli/exitcode"
	"github.com/spiegel-im-spiegel/gocli/rwi"
)

var (
	//Name is applicatin name
	Name = "aozora-bunko"
)
var (
	debugFlag bool //debug flag
	rawFlag   bool //raw flag
)

//newRootCmd returns cobra.Command instance for root command
func newRootCmd(ui *rwi.RWI, args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   Name,
		Short: "Search for books data",
		Long:  "Search for books data from Aozora-bunko API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return debugPrint(ui, ecode.ErrNoCommand)
		},
	}
	rootCmd.SilenceUsage = true
	rootCmd.SetArgs(args)               //arguments of command-line
	rootCmd.SetIn(ui.Reader())          //Stdin
	rootCmd.SetOutput(ui.ErrorWriter()) //Stdout and Stderr
	rootCmd.AddCommand(newSearchCmd(ui))
	rootCmd.AddCommand(newLookupCmd(ui))
	rootCmd.AddCommand(newRankingCmd(ui))

	//global options (others)
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "", false, "for debug")
	rootCmd.PersistentFlags().BoolVarP(&rawFlag, "raw", "", false, "Output raw data from API")

	return rootCmd
}

func debugPrint(ui *rwi.RWI, err error) error {
	if debugFlag && err != nil {
		fmt.Fprintf(ui.ErrorWriter(), "Error: %+v\n", err)
		return nil
	}
	return errs.Cause(err)
}

//Execute is called from main function
func Execute(ui *rwi.RWI, args []string) (exit exitcode.ExitCode) {
	defer func() {
		//panic hundling
		if r := recover(); r != nil {
			_ = ui.OutputErrln("Panic:", r)
			for depth := 0; ; depth++ {
				pc, src, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				_ = ui.OutputErrln(" ->", depth, ":", runtime.FuncForPC(pc).Name(), ":", src, ":", line)
			}
			exit = exitcode.Abnormal
		}
	}()

	//execution
	exit = exitcode.Normal
	if err := newRootCmd(ui, args).Execute(); err != nil {
		exit = exitcode.Abnormal
	}
	return
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
