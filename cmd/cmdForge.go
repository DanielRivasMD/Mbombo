/*
Copyright © 2024 danielrivasmd@gmail.com

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"errors"
	"strings"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO: evaluate argument refactoring => one struct to capture all flags
// TODO: evaluate reading toml config
var (
	options      forgeOptions
	replacePairs replaceFlags
)

type forgeOptions struct {
	inPath  string
	outPath string
	files   []string
}

type replacement struct {
	new string
	old string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// implement pflag.Value
type replaceFlags []replacement

func (r *replaceFlags) String() string {
	parts := make([]string, len(*r))
	for i, rep := range *r {
		parts[i] = rep.old + "=" + rep.new
	}
	return strings.Join(parts, ",")
}

func (r *replaceFlags) Set(val string) error {
	op := "flag.set"

	parts := strings.SplitN(val, "=", 2)
	if len(parts) != 2 {
		horus.CheckErr(
			errors.New(""),
			horus.WithOp(op),
			horus.WithMessage("invalid replace pair"),
			horus.WithExitCode(2),
			horus.WithFormatter(func(he *horus.Herror) string {
				return he.Message
			}),
		)
	}
	*r = append(*r, replacement{old: parts[0], new: parts[1]})
	return nil
}

func (r *replaceFlags) Type() string {
	return "old=new"
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(forgeCmd)

	forgeCmd.Flags().StringVarP(&options.inPath, "in", "", "", "Where are the itmes to be forged?")
	forgeCmd.Flags().StringVarP(&options.outPath, "out", "", "", "Where will the forge be delivered?")
	forgeCmd.Flags().StringArrayVarP(&options.files, "files", "", []string{}, "These items will create...")
	horus.CheckErr(forgeCmd.MarkFlagRequired("in"))
	horus.CheckErr(forgeCmd.MarkFlagRequired("out"))
	horus.CheckErr(forgeCmd.MarkFlagRequired("files"))

	forgeCmd.Flags().VarP(&replacePairs, "replace", "r", "replacement in form old=new, comma-separated")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var forgeCmd = &cobra.Command{
	Use:     "forge",
	Short:   "Forge products",
	Long:    helpForge,
	Example: exampleForge,

	Run: runForge,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runForge(cmd *cobra.Command, args []string) {
	op := "mbombo.forge"

	horus.CheckErr(
		catFiles(options),
		horus.WithOp(op),
		horus.WithMessage("Error during concatenation execution"),
		horus.WithExitCode(2),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
