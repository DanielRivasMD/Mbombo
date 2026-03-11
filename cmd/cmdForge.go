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
	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func ForgeCmd() *cobra.Command {
	d := horus.Must(domovoi.GlobalDocs())
	cmd := horus.Must(d.MakeCmd("forge", runForge))

	cmd.Flags().StringVarP(&forgeFlags.inPath, "in", "", "", "itmes to be forged")
	cmd.Flags().StringVarP(&forgeFlags.outPath, "out", "", "", "path to forge")
	cmd.Flags().StringArrayVarP(&forgeFlags.inFiles, "files", "", []string{}, "forge components")
	cmd.Flags().VarP(&replacePairs, "replace", "", "replacement in form old=new, space-separated")

	horus.CheckErr(cmd.MarkFlagRequired("out"))
	horus.CheckErr(cmd.MarkFlagRequired("files"))

	return cmd
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
