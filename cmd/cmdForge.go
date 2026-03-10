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

func init() {
	forgeCmd := MakeCmd("forge", runForge)

	forgeCmd.Flags().StringVarP(&forgeFlags.inPath, "in", "", "", "Where are the itmes to be forged?")
	forgeCmd.Flags().StringVarP(&forgeFlags.outPath, "out", "", "", "Where will the forge be delivered?")
	forgeCmd.Flags().StringArrayVarP(&forgeFlags.inFiles, "files", "", []string{}, "These items will create...")
	forgeCmd.Flags().VarP(&replacePairs, "replace", "", "replacement in form old=new, comma-separated")

	horus.CheckErr(forgeCmd.MarkFlagRequired("out"))
	horus.CheckErr(forgeCmd.MarkFlagRequired("files"))

	rootCmd.AddCommand(forgeCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func runForge(cmd *cobra.Command, args []string) {
	op := "mbombo.forge"

	normalizeForgeOptions(&forgeFlags)

	horus.CheckErr(
		catFiles(forgeFlags),
		horus.WithOp(op),
		horus.WithMessage("Error during concatenation execution"),
		horus.WithExitCode(2),
	)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
