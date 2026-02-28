/*
Copyright © 2024 Daniel Rivas <danielrivasmd@gmail.com>

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
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

const APP = "mbombo"
const VERSION = "v1.0.0"
const NAME = "Daniel Rivas"
const EMAIL = "<danielrivasmd@gmail.com>"

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:     GetUse("root"),
	Long:    formatLongHelp(GetHelp("root")),
	Example: GetExample("root"),
	Version: VERSION,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "Enable verbose diagnostics")
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	flags mbomboFlags
)

type mbomboFlags struct {
	verbose bool
}

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
	outFile string
	inFiles []string
}

type replacement struct {
	old string // anchor or token
	new string // full replacement string
	// TODO: add tab completion?
	mode string // "token" or "line"
}

////////////////////////////////////////////////////////////////////////////////////////////////////
