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
	"os"

	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	// outProduct string
	options forgeOptions
)

type forgeOptions struct {
	Out   string
	Path  string
	Files []string
	Old   []string
	New   []string
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(forgeCmd)

	// options := &forgeOptions{}

	forgeCmd.Flags().StringVarP(&options.Path, "path", "p", "", "Where are the itmes to be forged?")
	// forgeCmd.MarkFlagRequired("path")
	forgeCmd.Flags().StringArrayVarP(&options.Files, "files", "f", []string{}, "These items will create...")
	// forgeCmd.MarkFlagRequired("files")
	forgeCmd.Flags().StringVarP(&options.Out, "out", "", "", "Where will the forge be delivered?")
	// forgeCmd.MarkFlagRequired("out")
	forgeCmd.Flags().StringArrayVarP(&options.Old, "old", "o", []string{}, "Value to replace.")
	forgeCmd.Flags().StringArrayVarP(&options.New, "new", "n", []string{}, "Value replacement.")
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

// TODO: pass boolean for domovoi actions
func runForge(cmd *cobra.Command, args []string) {
	// forge file
		params := copyCR(options.Path, options.Out)
	params.files = options.Files
	params.reps = repsForge() // automatic binding cli flags

	// Call catFiles and capture any error.
	if err := catFiles(params); err != nil {
		// Log the error in JSON format for better debugging.
		// log.Printf("Error during catFiles execution: %s", horus.JSONFormatter(err))
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////
