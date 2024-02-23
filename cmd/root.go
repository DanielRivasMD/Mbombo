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
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

////////////////////////////////////////////////////////////////////////////////////////////////////
// Mbombo, also called Bumba, is the creator god in the religion
// and mythology of the Kuba people of Central Africa in the area
// that is now known as Democratic Republic of the Congo.
// In the Mbombo creation myth, Mbombo was a giant in form and
// white in color. The myth describes the creation of the universe
// from nothing.
////////////////////////////////////////////////////////////////////////////////////////////////////

// declarations
var ()

////////////////////////////////////////////////////////////////////////////////////////////////////

var rootCmd = &cobra.Command{
	Use:   "mbombo",
	Short: "",
	Long: chalk.Green.Color("Daniel Rivas <danielrivasmd@gmail.com>") + `

` + chalk.Green.Color("Mbombo") + chalk.Blue.Color(` will forge a unifed product.


`) + chalk.Green.Color("Mbombo") + ` creates a convenient command line interphase
with built-in and accessible documentation.
`,

	Example: `
` + chalk.Cyan.Color("mbombo") + ` help`,

	////////////////////////////////////////////////////////////////////////////////////////////////////

}

////////////////////////////////////////////////////////////////////////////////////////////////////

func Execute() {
	ε := rootCmd.Execute()
	if ε != nil {
		log.Fatal(ε)
		os.Exit(1)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func initializeConfig(κ *cobra.Command, configPath string, configName string) error {

	// initialize viper
	ʌ := viper.New()

	// collect config path & file from persistent flags
	ʌ.AddConfigPath(configPath)
	ʌ.SetConfigName(configName)

	// read the config file
	ε := ʌ.ReadInConfig()
	if ε != nil {
		// okay if there isn't a config file
		_, ϙ := ε.(viper.ConfigFileNotFoundError)
		if !ϙ {
			// return an error if we cannot parse the config file
			return ε
		}
	}

	// bind flags to viper
	bindFlags(κ, ʌ)

	return nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////

// bind each cobra flag to its associated viper configuration
func bindFlags(κ *cobra.Command, ʌ *viper.Viper) {

	κ.Flags().VisitAll(func(σ *pflag.Flag) {

		// apply the viper config value to the flag when the flag is not set and viper has a value
		if !σ.Changed && ʌ.IsSet(σ.Name) {
			ν := ʌ.Get(σ.Name)
			κ.Flags().Set(σ.Name, fmt.Sprintf("%v", ν))
		}
	})
}

////////////////////////////////////////////////////////////////////////////////////////////////////

func init() {

	// persistent flags
}

////////////////////////////////////////////////////////////////////////////////////////////////////