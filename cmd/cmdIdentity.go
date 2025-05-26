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

import (
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var (
	iden = chalk.Yellow.Color("Mbombo") + `, also called ` + chalk.Yellow.Color("Bumba") + `, is the creator god in the religion
and mythology of the ` + chalk.Cyan.Color("Kuba") + ` people of ` + chalk.Green.Color("Central Africa") + ` in the area
that is now known as Democratic Republic of the Congo

In the Mbombo creation myth, ` + chalk.Yellow.Color("Mbombo") + ` was a giant in form and
white in color. The myth describes the creation of the universe
from nothing`
)

////////////////////////////////////////////////////////////////////////////////////////////////////

// identityCmd
var identityCmd = &cobra.Command{
	Use:     "identity",
	Short:   `Reveal identity`,
	Long:    `Reveal identity`,
	Example: ``,

	////////////////////////////////////////////////////////////////////////////////////////////////////

	Run: func(κ *cobra.Command, args []string) {
		domovoi.LineBreaks()
		fmt.Println()

		fmt.Println(iden)

		domovoi.LineBreaks()
		fmt.Println()

	},
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// execute prior main
func init() {
	rootCmd.AddCommand(identityCmd)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
