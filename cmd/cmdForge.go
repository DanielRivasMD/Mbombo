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

import (
	"log"
	"os"

	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	outProduct string
)

////////////////////////////////////////////////////////////////////////////////////////////////////

var forgeCmd = &cobra.Command{
	Use:   "forge",
	Short: "Forge products.",
	Long: `Forge by defining the materials & the destination
Moreover, indicate pieces to replace.`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(cmd *cobra.Command, args []string) {
		// forge file
		params := copyCR(path, out)
		params.files = files
		params.reps = repsForge() // automatic binding cli flags

		// Call catFiles and capture any error.
		if err := catFiles(params); err != nil {
			// Log the error in JSON format for better debugging.
			log.Printf("Error during catFiles execution: %s", horus.FormatError(err, horus.JSONFormatter))
			os.Exit(1)
		}
	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {
	rootCmd.AddCommand(forgeCmd)

	// flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////
