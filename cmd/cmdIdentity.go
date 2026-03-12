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
	"fmt"

	"github.com/DanielRivasMD/domovoi"
	"github.com/DanielRivasMD/horus"
	"github.com/spf13/cobra"
)

////////////////////////////////////////////////////////////////////////////////////////////////////

func IdentityCmd() *cobra.Command {
	return horus.Must(horus.Must(domovoi.GlobalDocs()).MakeCmd("identity", runIdentity,
		domovoi.WithAliases([]string{"id"}),
	))
}

////////////////////////////////////////////////////////////////////////////////////////////////////

const IDENT = `Mbombo, also called Bumba, is the creator god in the religion
and mythology of the Kuba people of Central Africa in the area
that is now known as Democratic Republic of the Congo

In the Mbombo creation myth, Mbombo was a giant in form and
white in color. The myth describes the creation of the universe
from nothing

Role: Mbombo is considered a creator god in Bushongo mythology
Creation story: According to legend, in the beginning, there was only darkness and water, and Mbombo was the only being. He was a giant, pale god who eventually felt pain in his stomach and vomited up the sun, moon, stars, and then the Earth itself, including animals and people
Symbolism: His creation of the world through vomiting is unique and symbolic, often interpreted as a metaphor for creative force through suffering or sacrifice
Assistants: After creation, Mbombo delegated tasks to his sons and some of the first humans and animals to help finish shaping the world`

////////////////////////////////////////////////////////////////////////////////////////////////////

func runIdentity(cmd *cobra.Command, args []string) {
	fmt.Println()
	fmt.Println(IDENT)
	fmt.Println()
}

////////////////////////////////////////////////////////////////////////////////////////////////////
