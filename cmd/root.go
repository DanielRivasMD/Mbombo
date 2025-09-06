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

var rootCmd = &cobra.Command{
	Use:     "mbombo",
	Long:    helpRoot,
	Example: exampleRoot,
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	horus.CheckErr(rootCmd.Execute())
}

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	verbose bool
)

var (
	out   string
	path  string
	files []string
	old   []string
	new   []string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose diagnostics")

	rootCmd.PersistentFlags().StringVarP(&path, "path", "p", "", "Where are the itmes to be forged?")
	rootCmd.MarkFlagRequired("path")
	rootCmd.PersistentFlags().StringArrayVarP(&files, "files", "f", []string{}, "These items will create...")
	rootCmd.MarkFlagRequired("files")
	rootCmd.PersistentFlags().StringVarP(&out, "out", "", "", "Where will the forge be delivered?")
	rootCmd.MarkFlagRequired("out")
	rootCmd.PersistentFlags().StringArrayVarP(&old, "old", "o", []string{}, "Value to replace.")
	rootCmd.PersistentFlags().StringArrayVarP(&new, "new", "n", []string{}, "Value replacement.")
}

////////////////////////////////////////////////////////////////////////////////////////////////////
